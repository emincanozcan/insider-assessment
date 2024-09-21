package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL          string
	RedisURL             string
	Port                 string
	WebhookURL           string
	WebhookAuthKey       string
	MessageSendInterval  int
	MessageSendBatchSize int32
}

func Load() (*Config, error) {
	// TODO: handle missing parameters
	viper.AutomaticEnv()

	return &Config{
		DatabaseURL:          viper.GetString("DATABASE_URL"),
		RedisURL:             viper.GetString("REDIS_URL"),
		Port:                 viper.GetString("PORT"),
		WebhookURL:           viper.GetString("WEBHOOK_URL"),
		WebhookAuthKey:       viper.GetString("WEBHOOK_AUTH_KEY"),
		MessageSendInterval:  viper.GetInt("MESSAGE_SEND_INTERVAL"),
		MessageSendBatchSize: viper.GetInt32("MESSAGE_SEND_BATCH_SIZE"),
	}, nil
}
