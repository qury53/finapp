package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel  int    `envconfig:"LEVEL" default:"-4"`
	LogFolder string `envconfig:"FOLDER" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("LOG", &config); err != nil {
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
