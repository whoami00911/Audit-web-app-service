package repository

import (
	"context"

	"github.com/whoami00911/Audit-web-app-service/pkg/logEntities"
	"github.com/whoami00911/Audit-web-app-service/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type AddLog struct {
	db     *mongo.Database
	logger *logger.Logger
}

func InitRepoLogMethods(db *mongo.Database, logger *logger.Logger) *AddLog {
	return &AddLog{
		db:     db,
		logger: logger,
	}
}

func (a *AddLog) Log(ctx context.Context, log logEntities.Log) (logEntities.Status, error) {
	_, err := a.db.Collection("Logs").InsertOne(ctx, log)
	if err != nil {
		a.logger.Error(err)
		return logEntities.Status{
			Status: false,
		}, err
	}

	return logEntities.Status{
		Status: true,
	}, nil
}

func (a *AddLog) GinLog(ctx context.Context, ginLog logEntities.GinLog) (logEntities.Status, error) {
	_, err := a.db.Collection("GinLogs").InsertOne(ctx, ginLog)
	if err != nil {
		a.logger.Error(err)
		return logEntities.Status{
			Status: false,
		}, err
	}

	return logEntities.Status{
		Status: true,
	}, nil
}
