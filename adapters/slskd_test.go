package adapters

import (
	"testing"
	"time"
)

func TestConvertSlskdDuration(t *testing.T) {
	tests := []struct {
		duration, transferId string
		expected             time.Duration
	}{
		{"00:00:00", "t1", 0},
		{"00:25:00", "t2", 25 * time.Minute},
		{"14:24:59", "t3", 14*time.Hour + 24*time.Minute + 59*time.Second},
		{"02:59:45.590042", "t4", 2*time.Hour + 59*time.Minute + 45*time.Second + 590042*time.Microsecond},
		{"xx", "t5", 0},
	}

	for _, test := range tests {
		t.Run(test.transferId, func(t *testing.T) {
			result := convertSlskdDuration(test.duration, test.transferId)

			if result != test.expected {
				t.Errorf("%s failed, expected: %d, got: %d", test.transferId, test.expected, result)
			}
		})
	}
}
