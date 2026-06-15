package retry

import (
	"context"
	"time"
)

func RetryWithTimeout[T any](
	timeout time.Duration,
	interval time.Duration,
	multiplier int,
	CB func(int) (T, error),
	doneCB func(error, int) error,
) (T, error) {

	var nul T // for undefined T value
	var lastErr error
	var count int

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	if multiplier == 0 {
		multiplier = 1 // avoid *0 multiplication
	}

	stop := func() {
		ticker.Stop()
		cancel()
	}

	// 0 is the initial try
	result, err := CB(0)

	if err == nil {
		stop()
		return result, nil // return result there's no error
	}
	lastErr = err

	for {
		select {
		case <-ctx.Done():
			return nul, doneCB(lastErr, count)

		case <-ticker.C:
			count++

			result, err := CB(count)

			if err == nil {
				stop()
				return result, nil // return result there's no error
			}

			lastErr = err
			interval *= time.Duration(multiplier)
			ticker = time.NewTicker(interval)
		}
	}
}

func RetryWithNTry[T any](nTry int, interval time.Duration, multiplier int, CB func(int) (T, error)) (T, error) {
	var nul T
	var lastErr error

	if multiplier == 0 {
		multiplier = 1 // avoid *0 multiplication
	}

	// 0 is the initial try
	result, err := CB(0)

	if err == nil {
		return result, nil // return result there's no error
	}
	lastErr = err

	for count := 1; count <= nTry; count++ {
		interval *= time.Duration(multiplier)

		result, err := CB(count)

		if err == nil {
			return result, nil // if no error, return result
		}

		lastErr = err

		time.Sleep(interval)
	}

	return nul, lastErr
}
