package models

import "github.com/google/uuid"

type RelationContext struct {
	TransferID string
	FilePath   string
	BatchRef   uuid.UUID
}
