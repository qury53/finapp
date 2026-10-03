package income_repository_postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *IncomeRepository) GetIncomes(
	ctx context.Context,
	userID *int,
	since *time.Time,
) ([]domain.Income, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	SELECT id, version, value, title, description, status, created_at, user_id
	FROM finapp.income
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

	var incomeModels []IncomeModel
	for rows.Next() {
		var incomeModel IncomeModel
		if err := rows.Scan(
			&incomeModel.ID,
			&incomeModel.Version,
			&incomeModel.Value,
			&incomeModel.Title,
			&incomeModel.Description,
			&incomeModel.Status,
			&incomeModel.CreatedAt,
			&incomeModel.UserID,
		); err != nil {
			return nil, fmt.Errorf(
				"repository: scan error: %w",
				err,
			)
		}

		incomeModels = append(incomeModels, incomeModel)
	}

	return domainsFromModels(incomeModels), nil
}
