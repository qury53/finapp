package users_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
	core_encrypt "github.com/qury53/finapp/internal/core/encrypt"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("service validation: %w", err)
	}

	hashedPassword, err := core_encrypt.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = hashedPassword
	user, err = s.UsersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("service: %w", err)
	}

	return user, nil
}
