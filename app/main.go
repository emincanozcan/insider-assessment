package main

import (
	"fmt"
	"time"

	"github.com/emincanozcan/insider-assessment/internal/api"
	"github.com/emincanozcan/insider-assessment/internal/config"
	"github.com/emincanozcan/insider-assessment/internal/database"
	"github.com/emincanozcan/insider-assessment/internal/database/sqlc"
	"github.com/emincanozcan/insider-assessment/internal/redis"
	"github.com/emincanozcan/insider-assessment/internal/service"
	"github.com/emincanozcan/insider-assessment/internal/worker"
	"github.com/emincanozcan/insider-assessment/pkg/webhook_client"
	"github.com/emincanozcan/insider-assessment/pkg/webhook_server"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic("Missing environment variables!" + err.Error())
	}

	switch config.Mode {
	case "appserver":
		fmt.Println("Starting application server...")

		queries := sqlc.New(database.Initialize(config.DatabaseURL))

		redis, err := redis.Connect(config.RedisURL)
		if err != nil {
			panic("Can't connect to redis")
		}

		messageService := service.NewMessageService(queries, redis, webhook_client.NewClient(config.WebhookURL, config.WebhookAuthKey))

		messageSender := worker.MakeOrGet(messageService, time.Duration(config.MessageSendInterval)*time.Second, config.MessageSendBatchSize)
		go messageSender.StartBackgroundJob()

		api.InitializeApi(messageSender, messageService, config.ServerPort)

	case "webhookserver":
		fmt.Println("Starting webhook server...")
		webhook_server.InitializeServer(config.LocalWebhookServerPort)

	default:
		fmt.Println("Unknown mode. Please provide a valid mode: appserver or webhookserver")
	}
}
