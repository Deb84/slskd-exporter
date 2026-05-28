package database

import (
	"errors"
	"fmt"
	"slskd-exporter/models"
	"slskd-exporter/models/postgres"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BatchID struct {
	Value int64
	Set   bool
}

func (batchID *BatchID) SetValue(v int64) {
	batchID.Value = v
	batchID.Set = true
}

// err = ErrRecordNotFound -> false, err = nil -> true, err != nil & !ErrRecordNotFound -> false, err
func tryFind(tx *gorm.DB, query interface{}, dest interface{}) (bool, error) {
	err := tx.Where(query).Take(&dest).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

func (db *DB) CreateTransfer(tx *gorm.DB, transfer *postgres.Transfer, file *postgres.File, batchID *BatchID) error {

	//file
	handleFile := func(file *postgres.File) (int64, error) {
		var existingFile postgres.File

		fileExists, err := tryFind(tx, &postgres.File{Path: file.Path}, &existingFile)
		if err != nil {
			return 0, err
		}

		// file exists in db
		if fileExists {
			db.Logger.Verbose("File %s already exists", file.Path)
			file.ID = existingFile.ID
			return file.ID, nil
		}

		// file doesn't exists in db
		if err := tx.Create(file).Error; err != nil {
			return 0, fmt.Errorf(`Create file "%s": %w`, file.Path, err)
		}
		db.Logger.Info("File %s created", file.Path)
		return file.ID, nil
	}

	// batch
	newBatch := func(batchID *BatchID) *postgres.Batch {
		var batch postgres.Batch
		if batchID.Set { // batchID already set
			batch.ID = batchID.Value
			return &batch
		}
		return &batch
	}

	createBatch := func(batch *postgres.Batch, batchID *BatchID) error {
		if !batchID.Set {
			if err := tx.Create(batch).Error; err != nil {
				return fmt.Errorf("Create batch: %w", err)
			}

			batchID.SetValue(batch.ID)
			db.Logger.Info("Batch %s created", strconv.FormatInt(batch.ID, 10))
		}
		return nil
	}

	// transfer
	handleTransfer := func(transfer *postgres.Transfer, batchID *BatchID) error {
		var existingTransfer postgres.Transfer

		transferExists, err := tryFind(tx, &postgres.Transfer{TransferID: transfer.TransferID}, &existingTransfer)
		if err != nil {
			return err
		}

		batch := newBatch(batchID)

		// transfer exists and it's not different in db
		if transferExists && !IsDifferent(&existingTransfer, transfer) {
			db.Logger.Verbose("Transfer %s already exists", transfer.TransferID)
			batch.ID = existingTransfer.BatchID
			return nil
		}

		// transfer exists and it's different in db
		if transferExists {
			db.Logger.Verbose("Transfer %s already exists but it's different", transfer.TransferID)
			updatedTransfer := UpdateTransfer(&existingTransfer, transfer)
			tx.Save(updatedTransfer)
			db.Logger.Info("Transfer %s updated", transfer.TransferID)
			return nil
		}

		// transfer doesn't exists in db
		if err := createBatch(batch, batchID); err != nil {
			return err
		}

		transfer.BatchID = batch.ID

		if err := tx.Create(transfer).Error; err != nil {
			return fmt.Errorf(`Create transfer "%s": %w`, transfer.TransferID, err)
		}
		db.Logger.Info("Transfer %s created", transfer.TransferID)
		return nil

	}

	fileID, err := handleFile(file)
	if err != nil {
		return err
	}

	transfer.FileID = fileID

	if err := handleTransfer(transfer, batchID); err != nil {
		return err
	}
	return nil
}

func (db *DB) CreateTransfers(transfersContext *models.TransfersContext) error {
	files := transfersContext.Files                       // map[string]*postgres.File
	transfers := transfersContext.Transfers               // map[string]*postgres.Transfer
	relationContexts := transfersContext.RelationContexts // []*models.RelationContext

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		batchIDs := make(map[uuid.UUID]*BatchID)

		var errCount int
		var lastErr error

		for _, relation := range relationContexts {

			transfer := transfers[relation.TransferID] // *postgres.Transfer
			file := files[relation.FilePath]           // *postgres.File
			batchRef := relation.BatchRef              // uuid.UUID, runtime batch ref

			savePoint := fmt.Sprintf(`"%s"`, transfer.TransferID) // fix postgres sql cant parse "-"
			tx.SavePoint(savePoint)

			batchID, existing := batchIDs[batchRef]
			if !existing {
				batchIDs[batchRef] = &BatchID{Set: false}
				batchID = batchIDs[batchRef]
			}

			if err := db.CreateTransfer(tx, transfer, file, batchID); err != nil {
				errCount++
				lastErr = err
				tx.RollbackTo(savePoint)
				db.Logger.Warn("Unable to create transfer transferID=%s error=%s", transfer.TransferID, err.Error())
			}
		}

		if errCount >= len(relationContexts) {
			return lastErr
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
