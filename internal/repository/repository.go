package repository

import (
	"context"

	"github.com/whoami00911/Audit-web-app-service/pkg/logEntities"
)

type Logging interface {
	Log(ctx context.Context, log logEntities.Log) (logEntities.Status, error)
	GinLog(ctx context.Context, ginLog logEntities.GinLog) (logEntities.Status, error)
}

type Repository struct {
	Logging
}

func InitRepo(logging Logging) *Repository {
	return &Repository{
		Logging: logging,
	}
}
