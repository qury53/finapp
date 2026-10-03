package spending_repository_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (r *SpendingRepository) DeleteSpending(
	ctx context.Context,
	spendingID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	DELETE FROM finapp.spending
	WHERE id=$1;
	`

	cmdTag, err := r.Pool.Exec(ctx, query, spendingID)
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"repository: spending not found: %w",
			core_errors.ErrNotFound,
		)
	}

	if err != nil {
		return fmt.Errorf(
			"repository: failed to delete spending: %w",
			err,
		)
	}

	return nil
}
