package spending_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (r *SpendingRepository) GetSpending(
	ctx context.Context,
	spendingID int,
) (domain.Spending, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	SELECT id, version, value, title, description, status, category, created_at, user_id
	FROM finapp.spending
	WHERE id=$1;
	`

	row := r.Pool.QueryRow(ctx, query, spendingID)

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
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Spending{}, fmt.Errorf(
				"repository: spending not found: %w",
				core_errors.ErrNotFound,
			)
		}

		return domain.Spending{}, fmt.Errorf(
			"repository: scan error: %w",
			err,
		)
	}

	return domainFromModel(spendingModel), nil
}
