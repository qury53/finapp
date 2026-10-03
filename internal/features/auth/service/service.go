package auth_service

import (
	"context"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

type AuthService struct {
	UserWriter      UserWriter
	UserReader      UserReader
	JWTSecret       []byte
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type UserWriter interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

type UserReader interface {
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
}

func NewAuthService(
	config *Config,
	userWriter UserWriter,
	userReader UserReader,
) *AuthService {
	return &AuthService{
		UserWriter:      userWriter,
		UserReader:      userReader,
		JWTSecret:       []byte(config.JWTSecret),
		AccessTokenTTL:  config.AccessTokenTTL,
		RefreshTokenTTL: config.RefreshTokenTTL,
	}
}
