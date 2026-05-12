package database

import (
	"fmt"
	"log/slog"
	"slskd-exporter/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	DB     *gorm.DB
	Tables []any
}

type DBContext struct {
	DSN           string
	Dialector     gorm.Dialector
	GormOptions   []gorm.Option
	Timeout       time.Duration
	RetryInterval time.Duration
}

func NewDBConnection(dsn string) (*gorm.DB, error) {
	dbContext := DBContext{
		DSN:           dsn,
		Timeout:       30 * time.Second,
		RetryInterval: 5 * time.Second,
		Dialector:     postgres.Open(dsn),
		GormOptions: []gorm.Option{
			&gorm.Config{
				Logger: logger.Default.LogMode(logger.Silent),
			},
		},
	}

	db, err := connectDB(&dbContext)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDB(dbEnv models.DbEnv, tables ...any) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbEnv.HOST,
		dbEnv.USER,
		dbEnv.PASSWORD,
		dbEnv.DB,
		dbEnv.PORT,
	)

	dbConn, err := NewDBConnection(dsn)
	if err != nil {
		return nil, err
	}

	db := &DB{
		DB: dbConn,
	}

	db.Tables = tables

	db.ApplyAutoMigrate()

	return db, nil
}

func (db *DB) ApplyAutoMigrate() {
	err := db.DB.AutoMigrate(db.Tables...)
	if err != nil {
		slog.Warn("Unable to AutoMigrate the struct in the DB")
	}
}
