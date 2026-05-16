package database

import (
	"fmt"
	"slskd-exporter/logger"
	"slskd-exporter/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DB struct {
	Logger *logger.Logger
	DB     *gorm.DB
	Tables []any
}

type DBContext struct {
	Logger        *logger.Logger
	DSN           string
	Dialector     gorm.Dialector
	GormOptions   []gorm.Option
	Timeout       time.Duration
	RetryInterval time.Duration
}

func NewDBConnection(logger *logger.Logger, dsn string) (*gorm.DB, error) {
	dbContext := DBContext{
		Logger:        logger,
		DSN:           dsn,
		Timeout:       30 * time.Second,
		RetryInterval: 5 * time.Second,
		Dialector:     postgres.Open(dsn),
		GormOptions: []gorm.Option{
			&gorm.Config{
				Logger: gormLogger.Default.LogMode(gormLogger.Silent),
			},
		},
	}

	db, err := connectDB(&dbContext)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDB(logger *logger.Logger, dbEnv models.DbEnv, tables ...any) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbEnv.HOST,
		dbEnv.USER,
		dbEnv.PASSWORD,
		dbEnv.DB,
		dbEnv.PORT,
	)

	dbConn, err := NewDBConnection(logger, dsn)
	if err != nil {
		return nil, err
	}

	db := &DB{
		Logger: logger,
		DB:     dbConn,
	}

	db.Tables = tables

	db.ApplyAutoMigrate()

	return db, nil
}

func (db *DB) ApplyAutoMigrate() {
	err := db.DB.AutoMigrate(db.Tables...)
	if err != nil {
		db.Logger.Warn("Unable to AutoMigrate the struct in the DB")
	}
}
