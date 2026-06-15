package database

import (
	"fmt"
	"log/slog"
	"slskd-exporter/retry"

	"gorm.io/gorm"
)

func connectDB(dbContext *DBContext) (*gorm.DB, error) {
	cb := func(count int) (*gorm.DB, error) {
		db, err := gorm.Open(dbContext.Dialector, dbContext.GormOptions...)
		if err == nil {
			return db, nil
		}
		slog.Warn("Unable to connect to database, retrying...", "Try", count)
		return nil, err
	}

	doneCB := func(err error, count int) error {
		return fmt.Errorf("Unable to connect to database, please check the connection information. Error: %w", err)
	}

	db, err := retry.RetryWithTimeout(
		dbContext.Timeout,       // timeout
		dbContext.RetryInterval, // interval between each try
		1,                       // backoff multiplier
		cb,                      // callback to call at each try
		doneCB,                  // callback called at the end of timeout
	)
	if err != nil {
		return nil, err
	}

	slog.Info("Database connected")
	return db, nil
}
