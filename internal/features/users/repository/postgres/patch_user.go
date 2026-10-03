package user_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	userID int,
	patchedUser domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	UPDATE finapp.users
	SET
		password=$1,
		phone_number=$2,
		email=$3,
		version=version+1
	WHERE id=$4 AND version=$5
	RETURNING
		id,
		version,
		username,
		password,
		phone_number,
		email,
		created_at,
		updated_at;
	`

	var userModel UserModel
	row := r.Pool.QueryRow(
		ctx,
		query,
		patchedUser.Password,
		patchedUser.PhoneNumber,
		patchedUser.Email,
		userID,
		patchedUser.Version,
	)

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
			return domain.User{}, fmt.Errorf(
				"repository: user is concurrently accessed: %w",
				core_errors.ErrConflict,
			)
		}

		return domain.User{}, fmt.Errorf("repository: scan error: %w", err)
	}

	return domainFromModel(userModel), nil
}
