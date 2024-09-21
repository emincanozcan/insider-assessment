package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/emincanozcan/insider-assessment/internal/config"
	"github.com/emincanozcan/insider-assessment/internal/database"
	"github.com/emincanozcan/insider-assessment/internal/database/sqlc"
	"github.com/emincanozcan/insider-assessment/internal/redis"
	"github.com/emincanozcan/insider-assessment/internal/service"
	"github.com/emincanozcan/insider-assessment/pkg/webhook_client"
	"github.com/emincanozcan/insider-assessment/pkg/webhook_server"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic("Missing environment variables!" + err.Error())
	}

	database.RunMigrations(config.DatabaseURL)
	queries := sqlc.New(initDB(config.DatabaseURL))

	switch getMode() {
	case "appserver":
		fmt.Println("Starting application server...")
		redis, err := redis.Connect(config.RedisURL)
		if err != nil {
			panic("Can't connect to redis")
		}
		messageService := service.NewMessageService(queries, redis, webhook_client.NewClient(config.WebhookURL, config.WebhookAuthKey))
		sendMessagesLoop(messageService, config.MessageSendInterval, config.MessageSendBatchSize)

	case "webhookserver":
		fmt.Println("Starting webhook server...")
		webhook_server.InitializeServer(config.LocalWebhookServerPort)

	case "randomMessager":
		addSomeRandomMessages(queries)
		fmt.Println("100 messages added.")
		os.Exit(1)

	default:
		fmt.Println("Unknown mode. Please provide a valid mode: appserver, webhookserver, or randomMessager.")
	}
}

func getMode() string {
	var mode string
	flag.StringVar(&mode, "mode", "appserver", ` The mode flag can have 3 different values.
appserver (default) -> Starts the main application server, the one wanted in the case study.
webhookserver       -> Starts the webhook server, this replicates the webhook.site, so the appserver can be easily tested.
randomMessager      -> This adds random messages to the database in the background.
	`)
	flag.Parse()
	return mode
}

func addSomeRandomMessages(queries *sqlc.Queries) {
	for i := 0; i < 100; i++ {
		queries.CreateMessage(context.Background(), sqlc.CreateMessageParams{
			Content:   "Friendly reminder for your interview! " + string(i),
			Recipient: "emincan@emincanozcan.com",
		})
	}
}

func initDB(databaseURL string) *sql.DB {
	var err error
	err = database.RunMigrations(databaseURL)
	if err != nil {
		panic("Cant run migrations" + err.Error())
	}

	db, err := database.NewDB(databaseURL)
	return db
}

var isSending bool = true

func sendMessagesLoop(svc *service.MessageService, timeInSeconds int, batchSize int32) {
	ticker := time.NewTicker(time.Duration(timeInSeconds) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if !isSending {
			continue
		}

		svc.SendPendingMessages(context.Background(), batchSize)
	}
}
