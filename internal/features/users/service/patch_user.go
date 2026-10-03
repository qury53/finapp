package users_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
)

func (s *UsersService) PatchUser(
	ctx context.Context,
	userID int,
	patch domain.UserPatch,
) (domain.User, error) {
	if err := patch.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("service: failed to validate patch: %w", err)
	}

	user, err := s.UsersRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("service: failed to get user: %w", err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("service: failed to apply patch: %w", err)
	}

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("service: failed to validate patched user: %w", err)
	}

	fmt.Println("2")
	user, err = s.UsersRepository.PatchUser(ctx, userID, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("service: failed to patch user: %w", err)
	}

	return user, nil
}
