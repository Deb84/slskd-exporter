package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"slskd-exporter/models"

	"github.com/joho/godotenv"
)

// env
const (
	SLSKD_HOST     = "SLSKD_HOST"
	SLSKD_PORT     = "SLSKD_PORT"
	SLSKD_USER     = "SLSKD_USER"
	SLSKD_PASSWORD = "SLSKD_PASSWORD"
	SLSKD_CERT     = "SLSKD_TLS_CERT"

	POSTGRES_USER     = "POSTGRES_USER"
	POSTGRES_PASSWORD = "POSTGRES_PASSWORD"
	POSTGRES_DB       = "POSTGRES_DB"
	POSTGRES_HOST     = "POSTGRES_HOST"
	POSTGRES_PORT     = "POSTGRES_PORT"

	SCRAPE_INTERVAL = "SCRAPE_INTERVAL"
)

const envFile = ".env"

func parsePath(path string) string {
	if path == "" {
		return ""
	}

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		slog.Warn(fmt.Sprintf("Unable to find the certificate: File %s doesn't exist", path))
		path = ""
	}

	return ""
}

func parseScrapeInterval(intervalStr string) time.Duration {

	isValid := func(duration time.Duration) bool {
		if duration < 1*time.Second || duration == 0 {
			return false
		}
		return true
	}

	duration, err := time.ParseDuration(intervalStr)
	if err == nil && isValid(duration) {
		return duration
	}

	if interval, err := strconv.Atoi(intervalStr); err == nil {
		if duration := time.Duration(interval) * time.Second; isValid(duration) {
			return duration
		}
	}

	slog.Warn("Unable to parse Scrape Interval from env, default value (15s) will be used")
	slog.Warn(fmt.Sprintf("Env var %s must be in the time.Duration go format (e.g. 10s, 1h, etc) or a correct number (integer >= 1)", SCRAPE_INTERVAL))
	return 15 * time.Second
}

func GetEnv() models.Env {
	err := godotenv.Load(envFile)
	if err != nil {
		slog.Info("Unable to load .env, assuming Docker is used")
	}

	slskdEnv := models.SlskdEnv{
		HOST:     os.Getenv(SLSKD_HOST),
		PORT:     os.Getenv(SLSKD_PORT),
		USER:     os.Getenv(SLSKD_USER),
		PASSWORD: os.Getenv(SLSKD_PASSWORD),
		CERT:     os.Getenv(SLSKD_CERT),
	}

	dbEnv := models.DbEnv{
		USER:     os.Getenv(POSTGRES_USER),
		PASSWORD: os.Getenv(POSTGRES_PASSWORD),
		DB:       os.Getenv(POSTGRES_DB),
		HOST:     os.Getenv(POSTGRES_HOST),
		PORT:     os.Getenv(POSTGRES_PORT),
	}

	return models.Env{
		Slskd:          slskdEnv,
		DbEnv:          dbEnv,
		ScrapeInterval: parseScrapeInterval(os.Getenv(SCRAPE_INTERVAL)),
	}
}
