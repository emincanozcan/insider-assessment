package main

import (
	"time"

	"github.com/emincanozcan/insider-assessment/internal/api"
	"github.com/emincanozcan/insider-assessment/internal/config"
	"github.com/emincanozcan/insider-assessment/internal/database"
	"github.com/emincanozcan/insider-assessment/internal/database/sqlc"
	"github.com/emincanozcan/insider-assessment/internal/redis"
	"github.com/emincanozcan/insider-assessment/internal/service"
	"github.com/emincanozcan/insider-assessment/internal/worker"
	"github.com/emincanozcan/insider-assessment/pkg/webhook_client"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic("Missing environment variables!" + err.Error())
	}

	queries := sqlc.New(database.Initialize(config.DatabaseURL))

	redis, err := redis.Connect(config.RedisURL)
	if err != nil {
		panic("Can't connect to redis")
	}

	messageService := service.NewMessageService(queries, redis, webhook_client.NewClient(config.WebhookURL, config.WebhookAuthKey))

	messageSender := worker.MakeOrGet(messageService, time.Duration(config.MessageSendInterval)*time.Second, config.MessageSendBatchSize)
	go messageSender.StartBackgroundJob()

	api.InitializeApi(messageSender, messageService, config.Port)
}