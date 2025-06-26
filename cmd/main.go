package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/whoami00911/Audit-web-app-service/internal/repository"
	"github.com/whoami00911/Audit-web-app-service/internal/server"
	"github.com/whoami00911/Audit-web-app-service/internal/service"
	"github.com/whoami00911/Audit-web-app-service/pkg/logger"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка при чтении конфигурации: %v", err)
	}
}

func main() {
	logger := logger.GetLogger()

	tryMongo := repository.TryMongoConnect(repository.ConnectMongo)

	db, err := tryMongo()
	if err != nil {
		logger.Error(err)
		log.Fatalf("Can't connect mongodb")
	}

	repo := repository.InitRepo(repository.InitRepoLogMethods(db, logger))
	service := service.InitService(repo, logger)
	serverHandlers := server.InitGrpcServerHandlers(service, logger)
	server := server.InitGrpcServer(serverHandlers, logger)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server can't start")
		}
	}()

	fmt.Println("gRPC server successfully started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("Shutdown error: %s", err)
	}
}
