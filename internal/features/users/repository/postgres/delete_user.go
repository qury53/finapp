package user_repository_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (r *UsersRepository) DeleteUser(
	ctx context.Context,
	userID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	DELETE FROM finapp.users
	WHERE id = $1;
	`

	cmdTag, err := r.Pool.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("repository: exec error: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("repository: user not found: %w", core_errors.ErrNotFound)
	}

	return nil
}
