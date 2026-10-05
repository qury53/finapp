package user_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	users_repository "github.com/qury53/finapp/internal/features/users/repository"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	INSERT INTO finapp.users (username, password, phone_number, email, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, version, username, password, phone_number, email, created_at, updated_at;
	`

	row := r.Pool.QueryRow(
		ctx,
		query,
		user.Username,
		user.Password,
		user.PhoneNumber,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)

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
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.User{}, fmt.Errorf("repository: user already exists: %w", core_errors.ErrConflict)
		}

		return domain.User{}, fmt.Errorf("repository: scan error: %w", err)
	}

	return users_repository.DomainFromModel(userModel), nil
}
