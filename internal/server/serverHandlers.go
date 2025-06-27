package server

import (
	"context"

	"github.com/whoami00911/Audit-web-app-service/pkg/grpcPb"
	"github.com/whoami00911/Audit-web-app-service/pkg/logEntities"
	"github.com/whoami00911/Audit-web-app-service/pkg/logger"
)

type Logging interface {
	Log(ctx context.Context, req *grpcPb.LogRequest) (logEntities.Status, error)
	GinLog(ctx context.Context, req *grpcPb.GinLogRequest) (logEntities.Status, error)
}

type grpcServerHandlers struct {
	grpcPb.UnimplementedLogServiceServer
	logger *logger.Logger
	Logging
}

func InitGrpcServerHandlers(logging Logging, logger *logger.Logger) *grpcServerHandlers {
	return &grpcServerHandlers{
		logger:  logger,
		Logging: logging,
	}
}

func (g *grpcServerHandlers) Log(ctx context.Context, req *grpcPb.LogRequest) (*grpcPb.LogResponce, error) {
	status, err := g.Logging.Log(ctx, req)
	if err != nil {
		g.logger.Error(err)
		return &grpcPb.LogResponce{
			Status: status.Status,
		}, err
	}

	return &grpcPb.LogResponce{
		Status: status.Status,
	}, nil
}

func (g *grpcServerHandlers) GinLog(ctx context.Context, req *grpcPb.GinLogRequest) (*grpcPb.GinLogResponce, error) {
	status, err := g.Logging.GinLog(ctx, req)
	if err != nil {
		g.logger.Error(err)
		return &grpcPb.GinLogResponce{
			Status: status.Status,
		}, err
	}

	return &grpcPb.GinLogResponce{
		Status: status.Status,
	}, nil
}
