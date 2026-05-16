package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func connectionWithRetry(ctx context.Context, dbContext *DBContext) (*gorm.DB, error) {
	ticker := time.NewTicker(dbContext.RetryInterval)
	defer ticker.Stop()

	var lastErr error
	var tryCount int

	for {
		select {
		case <-ctx.Done():
			if lastErr != nil {
				err := fmt.Errorf("Unable to connect to database, please check the connection information. Error: %w", lastErr)
				return nil, err
			}
			return nil, ctx.Err()

		case <-ticker.C:
			tryCount++
			db, err := gorm.Open(dbContext.Dialector, dbContext.GormOptions...)
			if err == nil {
				dbContext.Logger.Info("Database connected")
				return db, nil
			}
			dbContext.Logger.Warn("Unable to connect to database, retrying... Try=%s", tryCount)
			lastErr = err
		}
	}

}

func connectDB(dbContext *DBContext) (*gorm.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbContext.Timeout)
	defer cancel()

	db, err := connectionWithRetry(ctx, dbContext)
	if err != nil {
		return nil, err
	}

	return db, nil
}
