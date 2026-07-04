package adapters

import (
	"fmt"
	"log/slog"
	"slskd-exporter/domain"
	"slskd-exporter/models/slskd"
	"strconv"
	"strings"
	"time"
)

func convertSlskdDuration(duration string, transferId string) time.Duration {
	errMsg := fmt.Sprintf("Unable to convert %s to a duration for transfer %s", duration, transferId)

	parts := strings.Split(duration, ":")

	if len(parts) != 3 {
		return 0
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		slog.Warn(errMsg)
		return 0
	}

	min, err := strconv.Atoi(parts[1])
	if err != nil {
		slog.Warn(errMsg)
		return 0
	}

	sec, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		slog.Warn(errMsg)
		return 0
	}

	return time.Duration(hours)*time.Hour +
		time.Duration(min)*time.Minute +
		time.Duration(sec*float64(time.Second))

}

func isoToTimestamp(iso string) time.Time {
	t, _ := time.Parse(time.RFC3339, iso)
	return t.UTC().Truncate(time.Microsecond)
}

func ConvertTransferSlskd(data *slskd.File) *domain.Transfer {

	file := &domain.File{
		Path: data.FileName,
	}

	return &domain.Transfer{
		Username:         data.Username,
		TransferID:       data.Id,
		File:             file,
		Direction:        data.Direction,
		Size:             data.Size,
		State:            data.State,
		RequestedAt:      isoToTimestamp(data.RequestedAt + "Z"), // slskd requestedAt & enqueuedAt field dont match RFC3339 because the time zone is missing
		EnqueuedAt:       isoToTimestamp(data.EnqueuedAt + "Z"),  // slskd requestedAt & enqueuedAt field dont match RFC3339 because the time zone is missing
		StartedAt:        isoToTimestamp(data.StartedAt),
		EndedAt:          isoToTimestamp(data.EndedAt),
		BytesTransferred: data.BytesTransferred,
		Exception:        data.Exception,
		AverageSpeed:     data.AverageSpeed,
		Attempts:         data.Attempts,
		BytesRemaining:   data.BytesRemaining,
		ElapsedTime:      convertSlskdDuration(data.ElapsedTime, data.Id),
		PercentComplete:  data.PercentComplete,
		RemainingTime:    convertSlskdDuration(data.RemainingTime, data.Id),
	}
}

func ConvertBatchSlskd(data *slskd.Transfer) *domain.Batch {
	var transfers []*domain.Transfer

	for _, directory := range data.Directories {
		for _, file := range directory.Files {
			transfers = append(transfers, ConvertTransferSlskd(file))
		}
	}

	return &domain.Batch{
		Transfers: transfers,
	}
}

func ConvertFromSlskd(data *slskd.Transfers) []*domain.Batch {
	var batchs []*domain.Batch

	for _, transfer := range *data {
		batchs = append(batchs, ConvertBatchSlskd(transfer))
	}

	return batchs
}
