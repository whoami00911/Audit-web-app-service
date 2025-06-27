package service

import (
	"context"

	"github.com/whoami00911/Audit-web-app-service/pkg/grpcPb"
	"github.com/whoami00911/Audit-web-app-service/pkg/logEntities"
	"github.com/whoami00911/Audit-web-app-service/pkg/logger"
)

type Logging interface {
	Log(ctx context.Context, log logEntities.Log) (logEntities.Status, error)
	GinLog(ctx context.Context, ginLog logEntities.GinLog) (logEntities.Status, error)
}

type Service struct {
	Logging
	logger *logger.Logger
}

func InitService(logging Logging, logger *logger.Logger) *Service {
	return &Service{
		Logging: logging,
		logger:  logger,
	}
}

func (s *Service) Log(ctx context.Context, req *grpcPb.LogRequest) (logEntities.Status, error) {
	log := logEntities.Log{
		Action: req.GetAction().String(),
		Method: req.GetMethod().String(),
		UserId: int(req.GetUserId()),
		ObjectId: func() []string {
			if req.ObjectId != nil {
				return req.ObjectId.ObjectId
			}
			return nil
		}(),
		Url:       req.Url,
		Timestamp: req.GetTimestamp().AsTime(),
	}

	status, err := s.Logging.Log(ctx, log)
	if err != nil {
		s.logger.Error(err)
		return logEntities.Status{
			Status: status.Status,
		}, err
	}

	return logEntities.Status{
		Status: status.Status,
	}, nil
}

func (s *Service) GinLog(ctx context.Context, req *grpcPb.GinLogRequest) (logEntities.Status, error) {
	ginLog := logEntities.GinLog{
		Timestamp:  req.GetTimestamp().AsTime(),
		StatusCode: int(req.StatusCode),
		Latency:    req.Latency,
		ClientIp:   req.ClientIp,
		Method:     req.Method,
		Path:       req.Path,
		UserAgent:  req.UserAgent,
	}

	status, err := s.Logging.GinLog(ctx, ginLog)
	if err != nil {
		s.logger.Error(err)
		return logEntities.Status{
			Status: status.Status,
		}, err
	}

	return logEntities.Status{
		Status: status.Status,
	}, nil
}
