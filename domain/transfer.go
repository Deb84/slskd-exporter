package domain

import "time"

type Transfer struct {
	Username         string
	TransferID       string
	File             *File
	Direction        string
	Size             int64
	State            string
	RequestedAt      time.Time
	EnqueuedAt       time.Time
	StartedAt        time.Time
	EndedAt          time.Time
	BytesTransferred int64
	AverageSpeed     float64
	Exception        string
	Attempts         int
	BytesRemaining   int64
	ElapsedTime      time.Duration
	PercentComplete  float64
	RemainingTime    time.Duration
}
