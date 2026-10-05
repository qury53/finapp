package user_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	users_repository "github.com/qury53/finapp/internal/features/users/repository"
)

func (r *UsersRepository) GetUser(ctx context.Context, userID int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	SELECT id, version, username, password, phone_number, email, created_at, updated_at
	FROM finapp.users
	WHERE id = $1;
	`

	row := r.Pool.QueryRow(ctx, query, userID)
	var userModel users_repository.UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Username,
		&userModel.Password,
		&userModel.PhoneNumber,
		&userModel.Email,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("repository: user not found: %w", core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("repository: scan error: %w", err)
	}

	return users_repository.DomainFromModel(userModel), nil
}
