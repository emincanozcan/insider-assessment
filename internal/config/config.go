package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL             string
	RedisURL                string
	ServerPort              string
	WebhookURL              string
	WebhookAuthKey          string
	StartLocalWebhookServer bool
	LocalWebhookServerPort  string
	MessageSendInterval     int
	MessageSendBatchSize    int32
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		DatabaseURL:             viper.GetString("DATABASE_URL"),
		RedisURL:                viper.GetString("REDIS_URL"),
		ServerPort:              viper.GetString("SERVER_PORT"),
		WebhookURL:              viper.GetString("WEBHOOK_URL"),
		WebhookAuthKey:          viper.GetString("WEBHOOK_AUTH_KEY"),
		StartLocalWebhookServer: viper.GetBool("START_LOCAL_WEBHOOK_SERVER"),
		LocalWebhookServerPort:  viper.GetString("LOCAL_WEBHOOK_SERVER_PORT"),
		MessageSendInterval:     viper.GetInt("MESSAGE_SEND_INTERVAL"),
		MessageSendBatchSize:    viper.GetInt32("MESSAGE_SEND_BATCH_SIZE"),
	}, nil
}
