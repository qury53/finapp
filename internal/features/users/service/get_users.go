package users_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *UsersService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit cannot be negative: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset cannot be negative: %w", core_errors.ErrInvalidArgument)
	}

	users, err := s.UsersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("service: :%w", err)
	}

	return users, nil
}
