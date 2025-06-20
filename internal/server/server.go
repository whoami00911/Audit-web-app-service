package server

import (
	"context"
	"fmt"
	"net"

	"github.com/whoami00911/Audit-web-app-service/pkg/grpcPb"
	"github.com/whoami00911/Audit-web-app-service/pkg/logger"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type grpcServer struct {
	grpcServer       *grpc.Server
	logger           *logger.Logger
	logServiceServer grpcPb.LogServiceServer
	addr             string
}

func InitGrpcServer(logServiceServer grpcPb.LogServiceServer, logger *logger.Logger) *grpcServer {
	return &grpcServer{
		grpcServer:       grpc.NewServer(),
		logger:           logger,
		logServiceServer: logServiceServer,
		addr:             viper.GetString("server.addr"),
	}
}

func (g *grpcServer) ListenAndServe() error {
	listener, err := net.Listen("tcp", g.addr)
	if err != nil {
		g.logger.Error(err)
		return err
	}

	grpcPb.RegisterLogServiceServer(g.grpcServer, g.logServiceServer)

	if err := g.grpcServer.Serve(listener); err != nil {
		g.logger.Error(err)
		return err
	}

	return nil
}

func (g *grpcServer) Shutdown(ctx context.Context) error {
	done := make(chan struct{}, 1)

	go func() {
		g.grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-ctx.Done():
		g.grpcServer.Stop()
		g.logger.Errorf("Graceful shutdown timed out, forcing immediate stop: %s", ctx.Err())
		fmt.Println("Graceful shutdown timed out, forcing immediate stop")

		return ctx.Err()

	case <-done:
		fmt.Println("Server shutdown gracefully")

		return nil
	}
}
