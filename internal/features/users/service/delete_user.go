package users_service

import (
	"context"
	"fmt"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *UsersService) DeleteUser(
	ctx context.Context,
	userID int,
) error {
	if userID < 1 {
		return fmt.Errorf("service: userID has to be positive: %w", core_errors.ErrInvalidArgument)
	}

	if err := s.UsersRepository.DeleteUser(ctx, userID); err != nil {
		return fmt.Errorf("service: failed to delete user: %w", err)
	}

	return nil
}
