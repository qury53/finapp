package income_repository_postgres

import core_repository_pgx "github.com/qury53/finapp/internal/core/repository/postgres"

type IncomeRepository struct {
	Pool *core_repository_pgx.Pool
}

func NewIncomeRepository(pool *core_repository_pgx.Pool) *IncomeRepository {
	return &IncomeRepository{
		Pool: pool,
	}
}
