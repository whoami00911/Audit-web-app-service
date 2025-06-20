package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/whoami00911/Audit-web-app-service/pkg/logger"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RetryBehavior func() (*mongo.Database, error)

type Config struct {
	maxRetries    int
	retryDelation time.Duration
	db            Mongodb
}

type Mongodb struct {
	URI      string
	User     string
	Password string
	Database string
}

func ConfigInicialize() *Config {
	return &Config{
		maxRetries:    3,
		retryDelation: 1 * time.Second,
		db:            Mongodb{},
	}
}

func ConnectMongo() (*mongo.Database, error) {
	logger := logger.GetLogger()
	cfg := ConfigInicialize()

	if err := godotenv.Load(".env"); err != nil {
		logger.Errorf("godotenv can't load env: %s", err)
		log.Panic("godotenv can't load env")

		return nil, err
	}

	if err := envconfig.Process("db", &cfg.db); err != nil {
		logger.Errorf("envconfig cant parse to struct: %s", err)
		log.Panic("envconfig cant parse to struct")

		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	opts := options.Client()

	opts.ApplyURI(cfg.db.URI)

	opts.SetAuth(options.Credential{
		Username: cfg.db.User,
		Password: cfg.db.Password,
	})

	server, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Errorf("Open connection to db server failed: %s", err)

		return nil, err
	}

	if err := server.Ping(context.Background(), nil); err != nil {
		logger.Errorf("ping db server failed: %s", err)
	}

	db := server.Database(cfg.db.Database)

	return db, nil
}

func TryMongoConnect(retryBehavior RetryBehavior) RetryBehavior {
	var err error
	var db *mongo.Database

	cfg := ConfigInicialize()

	return func() (*mongo.Database, error) {
		for i := 0; i < cfg.maxRetries; i++ {
			db, err = ConnectMongo()
			if err != nil {
				fmt.Printf("Не удалось подключиться, ожидаем %v перед повторной попыткой...\n", cfg.retryDelation)
			} else {
				return db, nil
			}

			<-time.After(cfg.retryDelation)
		}

		fmt.Printf("не удалось подключиться к базе данных после %d попыток", cfg.maxRetries)

		return nil, err
	}
}
