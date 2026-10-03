package auth_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
	core_encrypt "github.com/qury53/finapp/internal/core/encrypt"
)

func (s *AuthService) Register(
	ctx context.Context,
	user domain.User,
) (domain.User, string, string, error) {
	hashedPassword, err := core_encrypt.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, "", "", fmt.Errorf(
			"service: failed to hash password: %w",
			err,
		)
	}

	user.Password = hashedPassword
	user, err = s.UserWriter.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, "", "", fmt.Errorf(
			"service: failed to register user: %w",
			err,
		)
	}

	accessToken, err := s.GenerateAccessToken(user.ID)
	if err != nil {
		return domain.User{}, "", "", err
	}

	refreshToken, err := s.GenerateRefreshToken(user.ID)
	if err != nil {
		return domain.User{}, "", "", err
	}

	return user, accessToken, refreshToken, nil
}
