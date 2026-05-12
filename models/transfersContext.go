package models

import "slskd-exporter/models/postgres"

type TransfersContext struct {
	Files            map[string]*postgres.File
	Transfers        map[string]*postgres.Transfer
	RelationContexts []*RelationContext
}
