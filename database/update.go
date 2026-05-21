package database

import (
	"slskd-exporter/models/postgres"
	"time"
)

type Pair[E any, N any] struct {
	Exising E
	New     N
}

func IsDifferent(a *postgres.Transfer, b *postgres.Transfer) bool {
	return a.Size != b.Size ||
		a.State != b.State ||
		!a.EnqueuedAt.UTC().Truncate(time.Microsecond).Equal(b.EnqueuedAt) ||
		!a.StartedAt.UTC().Truncate(time.Microsecond).Equal(b.StartedAt) ||
		!a.EndedAt.UTC().Truncate(time.Microsecond).Equal(b.EndedAt) ||
		a.BytesTransferred != b.BytesTransferred ||
		a.AverageSpeed != b.AverageSpeed ||
		a.Exception != b.Exception ||
		a.Attempts != b.Attempts ||
		a.BytesRemaining != b.BytesRemaining ||
		a.ElapsedTime != b.ElapsedTime ||
		a.PercentComplete != b.PercentComplete ||
		a.RemainingTime != b.RemainingTime
}

func UpdateTransfer(updatedTr *postgres.Transfer, newTr *postgres.Transfer) *postgres.Transfer {
	updatedTr.Size = newTr.Size                         // Size
	updatedTr.State = newTr.State                       // State
	updatedTr.EnqueuedAt = newTr.EnqueuedAt             // EnqueuedAt
	updatedTr.StartedAt = newTr.StartedAt               // StartedAt
	updatedTr.EndedAt = newTr.EndedAt                   // EndedAt
	updatedTr.BytesTransferred = newTr.BytesTransferred // BytesTransferred
	updatedTr.AverageSpeed = newTr.AverageSpeed         // AverageSpeed
	updatedTr.Exception = newTr.Exception               // Exception
	updatedTr.Attempts = newTr.Attempts                 // Attempts
	updatedTr.BytesRemaining = newTr.BytesRemaining     // BytesRemaining
	updatedTr.ElapsedTime = newTr.ElapsedTime           // ElapsedTime
	updatedTr.PercentComplete = newTr.PercentComplete   // PercentComplete
	updatedTr.RemainingTime = newTr.RemainingTime       // RemainingTime

	return updatedTr
}
