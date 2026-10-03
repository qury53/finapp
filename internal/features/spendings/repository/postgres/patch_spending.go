package spending_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/finapp/internal/core/domain"
)

func (r *SpendingRepository) PatchSpending(
	ctx context.Context,
	spendingID int,
	spending domain.Spending,
) (domain.Spending, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	UPDATE finapp.spending
	SET
		value=$1,
		title=$2,
		description=$3,
		status=$4,
		category=$5,
		version=version+1
	WHERE id=$6 AND version=$7
	RETURNING
		id,
		version,
		value,
		title,
		description,
		status,
		category,
		created_at,
		user_id;
	`

	row := r.Pool.QueryRow(
		ctx,
		query,
		spending.Value,
		spending.Title,
		spending.Description,
		spending.Status,
		spending.Category,
		spendingID,
		spending.Version,
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
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Spending{}, fmt.Errorf(
				"repository: spending concurrently accessed: %w",
				err,
			)
		}

		return domain.Spending{}, fmt.Errorf(
			"repository: scan error: %w",
			err,
		)
	}

	return domainFromModel(spendingModel), nil
}
