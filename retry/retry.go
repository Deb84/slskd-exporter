package retry

import (
	"context"
	"time"
)

func doRetryWithTimeout[T any](
	timeout time.Duration,
	interval time.Duration,
	multiplier int,
	CB func() (T, error),
	doneCB func(error) error,
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

	if result, lastErr := CB(); lastErr == nil {
		stop()
		return result, nil // return result there's no error
	}

	for {
		select {
		case <-ctx.Done():
			return nul, doneCB(lastErr)

		case <-ticker.C:
			count++
			interval *= time.Duration(multiplier)

			if result, lastErr := CB(); lastErr == nil {
				stop()
				return result, nil // return result there's no error
			}

		}
	}
}

func retryWithNTry[T any](nTry int, interval time.Duration, multiplier int, CB func() (T, error)) (T, error) {
	var nul T
	var lastErr error
	var count int

	if multiplier == 0 {
		multiplier = 1 // avoid *0 multiplication
	}

	if result, lastErr := CB(); lastErr == nil {
		return result, nil // return result there's no error
	}

	for count <= nTry {
		count++
		interval *= time.Duration(multiplier)

		result, lastErr := CB()

		if lastErr == nil {
			return result, nil // if no error, return result
		}

		time.Sleep(interval)
	}

	return nul, lastErr
}
