package income_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (r *IncomeRepository) GetIncome(
	ctx context.Context,
	incomeID int,
) (domain.Income, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	SELECT id, version, value, title, description, status, created_at, user_id
	FROM finapp.income
	WHERE id=$1;
	`

	row := r.Pool.QueryRow(ctx, query, incomeID)

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
				"repository: income not found: %w",
				core_errors.ErrNotFound,
			)
		}

		return domain.Income{}, fmt.Errorf(
			"repository: scan error: %w",
			err,
		)
	}

	return domainFromModel(incomeModel), nil
}
