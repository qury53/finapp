package income_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/finapp/internal/core/domain"
)

func (r *IncomeRepository) PatchIncome(
	ctx context.Context,
	incomeID int,
	income domain.Income,
) (domain.Income, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	UPDATE finapp.income
	SET
		value=$1,
		title=$2,
		description=$3,
		status=$4,
		version=version+1
	WHERE id=$5 AND version=$6
	RETURNING
		id,
		version,
		value,
		title,
		description,
		status,
		created_at,
		user_id
	`

	row := r.Pool.QueryRow(
		ctx,
		query,
		income.Value,
		income.Title,
		income.Description,
		income.Status,
		incomeID,
		income.Version,
	)

	var incomeModel IncomeModel
	if err := row.Scan(
		&incomeModel.ID,
		&incomeModel.Version,
		&incomeModel.Value,
		&incomeModel.Title,
		&incomeModel.Description,
		&incomeModel.Status,
		&incomeModel.CreatedAt,
		&incomeModel.UserID,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Income{}, fmt.Errorf(
				"repository: income concurrently accessed: %w",
				err,
			)
		}

		return domain.Income{}, fmt.Errorf(
			"repository: scan error: %w",
			err,
		)
	}

	return domainFromModel(incomeModel), nil
}
