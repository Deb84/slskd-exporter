package adapters

import (
	"slskd-exporter/domain"
	"slskd-exporter/models"
	"slskd-exporter/models/postgres"

	"github.com/google/uuid"
)

// convert domain.Transfer into postgres.Transfer
func ConvertTransferPostgres(data *domain.Transfer) *postgres.Transfer {

	return &postgres.Transfer{
		Username:         data.Username,
		TransferID:       data.TransferID,
		Direction:        data.Direction,
		Size:             data.Size,
		State:            data.State,
		RequestedAt:      data.RequestedAt,
		EnqueuedAt:       data.EnqueuedAt,
		StartedAt:        data.StartedAt,
		EndedAt:          data.EndedAt,
		BytesTransferred: data.BytesTransferred,
		Exception:        data.Exception,
		AverageSpeed:     data.AverageSpeed,
		Attempts:         data.Attempts,
		BytesRemaining:   data.BytesRemaining,
		ElapsedTime:      data.ElapsedTime,
		PercentComplete:  data.PercentComplete,
		RemainingTime:    data.RemainingTime,
	}
}

// convert domain.File into postgres.File
func ConvertFilePostgres(data *domain.File) *postgres.File {
	return &postgres.File{
		Path:       data.Path,
		FileName:   data.FileName,
		ArtistName: data.ArtistName,
		AlbumName:  data.AlbumName,
		Year:       data.Year,
	}
}

func ConvertToPostgres(data []*domain.Batch) *models.TransfersContext {
	postgresFiles := make(map[string]*postgres.File)
	postgresTransfers := make(map[string]*postgres.Transfer)
	var relationContexts []*models.RelationContext

	for _, batch := range data {
		BatchRef := uuid.New()

		for _, transfer := range batch.Transfers {
			relationContexts = append(relationContexts, &models.RelationContext{
				TransferID: transfer.TransferID,
				FilePath:   transfer.File.Path,
				BatchRef:   BatchRef,
			})

			postgresFiles[transfer.File.Path] = ConvertFilePostgres(transfer.File)
			postgresTransfers[transfer.TransferID] = ConvertTransferPostgres(transfer)
		}
	}

	return &models.TransfersContext{
		Files:            postgresFiles,
		Transfers:        postgresTransfers,
		RelationContexts: relationContexts,
	}
}
