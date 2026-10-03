package users_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *UsersService) GetUser(ctx context.Context, userID int) (domain.User, error) {
	if userID < 1 {
		return domain.User{}, fmt.Errorf(
			"service: userID has to be positive interger: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	user, err := s.UsersRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("service: failed to get user: %w", err)
	}

	return user, nil
}
