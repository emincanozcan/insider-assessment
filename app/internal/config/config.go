package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mode string

	DatabaseURL          string
	RedisURL             string
	ServerPort           string
	WebhookURL           string
	WebhookAuthKey       string
	MessageSendInterval  int
	MessageSendBatchSize int32

	LocalWebhookServerPort string
}

func Load() (*Config, error) {
	// TODO: handle missing parameters
	viper.AutomaticEnv()

	return &Config{
		Mode:                 viper.GetString("MODE"),

		DatabaseURL:          viper.GetString("DATABASE_URL"),
		RedisURL:             viper.GetString("REDIS_URL"),
		ServerPort:           viper.GetString("SERVER_PORT"),
		WebhookURL:           viper.GetString("WEBHOOK_URL"),
		WebhookAuthKey:       viper.GetString("WEBHOOK_AUTH_KEY"),
		MessageSendInterval:  viper.GetInt("MESSAGE_SEND_INTERVAL"),
		MessageSendBatchSize: viper.GetInt32("MESSAGE_SEND_BATCH_SIZE"),

		LocalWebhookServerPort: viper.GetString("LOCAL_WEBHOOK_SERVER_PORT"),
	}, nil
}
