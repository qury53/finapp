package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel  int    `envconfig:"LOG_LEVEL" default:"-4"`
	LogFolder string `envconfig:"LOG_FOLDER" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("LOGGER", &config); err != nil {
		return nil, fmt.Errorf("failed to proccess .env file: %w", err)
	}

	return &config, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	return config
}
