package auth_service

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	JWTSecret       string        `envconfig:"JWT_SECRET"`
	AccessTokenTTL  time.Duration `envconfig:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL time.Duration `envconfig:"REFRESH_TOKEN_TTL"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("AUTH", &config); err != nil {
		return nil, fmt.Errorf(
			"failed to process auth env config: %w",
			err,
		)
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
