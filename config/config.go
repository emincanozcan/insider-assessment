package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
	RedisURL    string
	ServerPort  string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
		RedisURL:    viper.GetString("REDIS_URL"),
		ServerPort:  viper.GetString("SERVER_PORT"),
	}, nil
}
