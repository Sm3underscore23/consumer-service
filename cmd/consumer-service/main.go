package main

import (
	"consumer-service/iternal/config"
	orderRepo "consumer-service/iternal/repository/order"
	orderServ "consumer-service/iternal/service/order"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	configPath = "config/config.yaml"
)

func main() {
	ctx := context.Background()

	mainConfig, err := config.NewMainConfig(configPath)
	if err != nil {
		log.Fatalf("loading config error: %s", err)
	}

	dnsDb := mainConfig.DbConfigLoad()

	pool, err := pgxpool.New(ctx, dnsDb)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping to database: %s", err)
	}

	kafkaConfig := mainConfig.KafkaConfigLoad()

	repo := orderRepo.NewRepository(pool)
	orderServ.NewMessageService(
		ctx,
		repo,
		kafkaConfig.Assignor,
		kafkaConfig.ConsumerGroup,
		kafkaConfig.BrokerList,
		kafkaConfig.TopicName,
	)
}
