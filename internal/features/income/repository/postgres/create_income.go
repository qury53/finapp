package income_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (r *IncomeRepository) CreateIncome(
	ctx context.Context,
	income domain.Income,
) (domain.Income, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.Timeout())
	defer cancel()

	query := `
	INSERT INTO finapp.income (value, title, description, status, user_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, version, value, title, description, status, created_at, user_id;
	`

	row := r.Pool.QueryRow(
		ctx,
		query,
		income.Value,
		income.Title,
		income.Description,
		income.Status,
		income.UserID,
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return domain.Income{}, fmt.Errorf(
				"repository: user with id: %d not found: %w",
				income.UserID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Income{}, fmt.Errorf("repository: scan error: %w", err)
	}

	fmt.Println(incomeModel.Status)
	return domainFromModel(incomeModel), nil
}
