package auth_service

import (
	"context"
	"fmt"
	"log/slog"

	core_encrypt "github.com/qury53/finapp/internal/core/encrypt"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *AuthService) Login(
	ctx context.Context,
	username string,
	password string,
) (string, string, error) {
	user, err := s.UserReader.GetUserByUsername(ctx, username)
	if err != nil {
		slog.Error("get user by username", "error", err, "username", username)
		return "", "", fmt.Errorf(
			"service: failed to get user by name:%w",
			core_errors.ErrInvalidCredentials,
		)
	}

	if err := core_encrypt.CompareHashAndPassword(
		user.Password,
		password,
	); err != nil {
		return "", "", fmt.Errorf(
			"service: %w %s %s",
			core_errors.ErrInvalidCredentials,
			string(password),
			user.Password,
		)
	}

	accessToken, err := s.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", fmt.Errorf(
			"service: %w",
			err,
		)
	}

	refreshToken, err := s.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", fmt.Errorf(
			"service: %w",
			err,
		)
	}

	return accessToken, refreshToken, nil
}
