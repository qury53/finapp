package spending_repository_postgres

import core_repository_pgx "github.com/qury53/finapp/internal/core/repository/postgres"

type SpendingRepository struct {
	Pool *core_repository_pgx.Pool
}

func NewSpendingRepository(pool *core_repository_pgx.Pool) *SpendingRepository {
	return &SpendingRepository{
		Pool: pool,
	}
}
