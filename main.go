package main

import (
	"log/slog"
	"slskd-exporter/adapters"
	"slskd-exporter/database"
	"slskd-exporter/logger"
	"slskd-exporter/models"
	"slskd-exporter/models/postgres"
	slskdModels "slskd-exporter/models/slskd"
	"slskd-exporter/slskd"
	"time"
)

func sorry(err error) {
	slog.Error(":(")
	panic(err)
}

func convertData(slskdTransfers *slskdModels.Transfers) *models.TransfersContext {
	transfers := adapters.ConvertFromSlskd(slskdTransfers)
	postgresTransfers := adapters.ConvertToPostgres(transfers)
	return postgresTransfers
}

func scrape(db *database.DB, session *slskd.ExtractionSession) error {
	scrapedData, err := session.DoExtraction()
	if err != nil {
		return err
	}

	convertedDownload := convertData(scrapedData.Downloads)
	convertedUpload := convertData(scrapedData.Uploads)

	err = db.CreateTransfers(convertedDownload)
	if err != nil {
		return err
	}

	err = db.CreateTransfers(convertedUpload)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	env := GetEnv()

	logger := logger.NewLogger()
	logger.Config.SetLevel(env.LogLevel)

	db, err := database.NewDB(logger, env.DbEnv, &postgres.Transfer{}, &postgres.File{}, &postgres.Batch{})
	if err != nil {
		sorry(err)
	}

	session, err := slskd.NewExtraction(logger, env.Slskd)
	if err != nil {
		sorry(err)
	}

	ticker := time.NewTicker(env.ScrapeInterval)
	defer ticker.Stop()

	for range ticker.C {

		err := scrape(db, session)
		if err != nil {
			logger.Error("Unable to complete this scrape, retrying in %fs", env.ScrapeInterval.Seconds())
		}
	}
}
