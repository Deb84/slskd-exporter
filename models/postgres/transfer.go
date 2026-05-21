package postgres

import "time"

type Transfer struct {
	ID               int64         `gorm:"primaryKey;autoIncrement"`
	BatchID          int64         `gorm:"not null;index"`
	TransferID       string        `gorm:"not null;uniqueIndex;column:transfer_uuid"`
	Username         string        `gorm:"not null"`
	FileID           int64         `gorm:"not null;index"`
	Direction        string        `gorm:"not null"`
	Size             int64         `gorm:"not null"`
	State            string        `gorm:"not null"`
	RequestedAt      time.Time     `gorm:"not null"`
	EnqueuedAt       time.Time     `gorm:"not null"`
	StartedAt        time.Time     `gorm:"default:null"`
	EndedAt          time.Time     `gorm:"default:null"`
	BytesTransferred int64         `gorm:"default:null"`
	AverageSpeed     float64       `gorm:"default:null"`
	Exception        string        `gorm:"default:null"`
	Attempts         int           `gorm:"not null"`
	BytesRemaining   int64         `gorm:"not null"`
	ElapsedTime      time.Duration `gorm:"not null"`
	PercentComplete  float64       `gorm:"not null"`
	RemainingTime    time.Duration `gorm:"not null"`
}
