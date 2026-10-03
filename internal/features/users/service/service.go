package users_service

import (
	"context"

	"github.com/qury53/finapp/internal/core/domain"
)

type UsersService struct {
	UsersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, userID int) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	DeleteUser(ctx context.Context, userID int) error
	PatchUser(ctx context.Context, userID int, patchedUser domain.User) (domain.User, error)
}

func NewUsersService(repository UsersRepository) *UsersService {
	return &UsersService{
		UsersRepository: repository,
	}
}
