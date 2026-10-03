package spending_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (r *SpendingRepository) CreateSpending(
	ctx context.Context,
	spending domain.Spending,
) (domain.Spending, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	INSERT INTO finapp.spending (value, title, description, status, category, user_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, version, value, title, description, status, category, created_at, user_id;
	`

	row := r.Pool.QueryRow(
		ctx,
		query,
		spending.Value,
		spending.Title,
		spending.Description,
		spending.Status,
		spending.Category,
		spending.UserID,
	)

	var spendingModel SpendingModel
	if err := row.Scan(
		&spendingModel.ID,
		&spendingModel.Version,
		&spendingModel.Value,
		&spendingModel.Title,
		&spendingModel.Description,
		&spendingModel.Status,
		&spendingModel.Category,
		&spendingModel.CreatedAt,
		&spendingModel.UserID,
	); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return domain.Spending{}, fmt.Errorf(
				"repository: user with id: %d not found: %w",
				spending.UserID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Spending{}, fmt.Errorf("repository: scan error: %w", err)
	}

	return domainFromModel(spendingModel), nil
}
