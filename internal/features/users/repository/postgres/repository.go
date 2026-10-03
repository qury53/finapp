package user_repository_postgres

import core_repository_pgx "github.com/qury53/finapp/internal/core/repository/postgres"

type UsersRepository struct {
	Pool *core_repository_pgx.Pool
}

func NewUsersRepository(pool *core_repository_pgx.Pool) *UsersRepository {
	return &UsersRepository{
		Pool: pool,
	}
}
