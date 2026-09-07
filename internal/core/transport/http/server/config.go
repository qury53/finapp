package transport_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// HTTP_ADDR required
// SHUT_DOWN_TIMEOUT default 30s
type Config struct {
	HTTPAddr     string        `envconfig:"ADDR" required:"true"`
	ShutdownTime time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("HTTP", &config); err != nil {
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
