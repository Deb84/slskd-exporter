package main

import (
	"fmt"
	"log/slog"
	"slskd-exporter/adapters"
	"slskd-exporter/database"
	"slskd-exporter/models"
	"slskd-exporter/models/postgres"
	slskdModels "slskd-exporter/models/slskd"
	"slskd-exporter/slskd"
	"sync"
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
	err = db.CreateTransfers(convertedUpload)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	env := GetEnv()
	db, err := database.NewDB(env.DbEnv, &postgres.Transfer{}, &postgres.File{}, &postgres.Batch{})
	if err != nil {
		sorry(err)
	}

	session := slskd.NewExtraction(env.Slskd)

	ticker := time.NewTicker(env.ScrapeInterval)
	defer ticker.Stop()

	var mu sync.Mutex

	for range ticker.C {
		if !mu.TryLock() {
			continue
		}

		err := scrape(db, session)
		if err != nil {
			slog.Error(fmt.Sprintf("Unable to complete this scrape, retrying in %fs", env.ScrapeInterval.Seconds()))
		}
		mu.Unlock()
	}
}
