package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/emincanozcan/insider-assessment/internal/config"
	"github.com/emincanozcan/insider-assessment/internal/database"
	"github.com/emincanozcan/insider-assessment/internal/database/sqlc"
	"github.com/emincanozcan/insider-assessment/internal/redis"
	"github.com/emincanozcan/insider-assessment/pkg/webhook_client"
	"github.com/emincanozcan/insider-assessment/pkg/webhook_server"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic("Missing environment variables!" + err.Error())
	}

	if config.StartLocalWebhookServer {
		go webhook_server.InitializeServer(config.LocalWebhookServerPort)
		time.Sleep(100 * time.Millisecond)
	}

	wc := webhook_client.NewClient(config.WebhookURL, config.WebhookAuthKey)
	res, err := wc.Send("emincan@emincanozcan.com", "Reminder: interview")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	select {}

	return
	os.Exit(1)

	// MIGRATE
	err = database.RunMigrations(config.DatabaseURL)
	if err != nil {
		panic("Cant run migrations" + err.Error())
	}

	// SQLC Test
	db, err := database.NewDB(config.DatabaseURL)
	if err != nil {
		panic("Can't connect to new db")
	}
	queries := sqlc.New(db)
	queries.CreateMessage(context.Background(), sqlc.CreateMessageParams{
		Content:   "Reminder: interview",
		Recipient: "emincan@emincanozcan.com",
		Status:    0,
	})

	data, err := queries.ListPendingMessages(context.Background(), 10)

	if err != nil {
		panic("error")
	}

	fmt.Println(data)

	// TEST REDIS CONNECTION
	redis, err := redis.Connect(config.RedisURL)
	if err != nil {
		panic("Can't connect to redis")
	}

	err = redis.Set(context.Background(), "redis-demo-key", "redis-demo-value", 0).Err()
	if err != nil {
		panic("Can't set value in Redis")
	}

	val, err := redis.Get(context.Background(), "redis-demo-key").Result()
	if err != nil {
		panic("Can't get from redis" + err.Error())
	}

	fmt.Println("Redis:", val)
}

func runLocalWebhookServer() {
}
