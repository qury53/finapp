package user_repository_postgres

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *UsersRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	SELECT id, version, username, password, phone_number, email, created_at, updated_at
	FROM finapp.users
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	rows, err := r.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("repository: query error: %w", err)
	}

	var userModels []UserModel
	for rows.Next() {
		var userModel UserModel
		if err := rows.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.Username,
			&userModel.Password,
			&userModel.PhoneNumber,
			&userModel.Email,
			&userModel.CreatedAt,
			&userModel.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("repository: scan error: %w", err)
		}

		userModels = append(userModels, userModel)
	}

	return domainsFromModels(userModels), nil
}
