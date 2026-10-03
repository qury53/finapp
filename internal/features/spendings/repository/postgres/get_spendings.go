package spending_repository_postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *SpendingRepository) GetSpendings(
	ctx context.Context,
	userID *int,
	since *time.Time,
) ([]domain.Spending, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	SELECT id, version, value, title, description, status, category, created_at, user_id
	FROM finapp.spending
	`

	var builder strings.Builder
	var args []any
	builder.WriteString(query)
	whereState := false

	if userID != nil {
		builder.WriteString("WHERE user_id=$1")
		whereState = true
		args = append(args, *userID)
	}

	if since != nil {
		args = append(args, *since)
		if !whereState {
			builder.WriteString("WHERE ")
		} else {
			builder.WriteString(" AND ")
		}

		s := fmt.Sprintf("created_at >= $%d", len(args))
		builder.WriteString(s)
	}

	rows, err := r.Pool.Query(ctx, builder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf(
			"repository: query error: %w",
			err,
		)
	}

	var spendingModels []SpendingModel
	for rows.Next() {
		var spendingModel SpendingModel
		if err := rows.Scan(
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
			return nil, fmt.Errorf(
				"repository: scan error: %w",
				err,
			)
		}

		spendingModels = append(spendingModels, spendingModel)
	}

	return domainsFromModels(spendingModels), nil
}
