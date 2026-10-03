package income_repository_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (r *IncomeRepository) DeleteIncome(
	ctx context.Context,
	incomeID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	DELETE FROM finapp.income
	WHERE id=$1;
	`

	cmdTag, err := r.Pool.Exec(ctx, query, incomeID)
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"repository: income not found: %w",
			core_errors.ErrNotFound,
		)
	}

	if err != nil {
		return fmt.Errorf(
			"repository: failed to delete income: %w",
			err,
		)
	}

	return nil
}
