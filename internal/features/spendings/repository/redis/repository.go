package spending_repository_redis

import (
	core_repository_redis "github.com/qury53/finapp/internal/core/repository/redis"
	spending_repository_postgres "github.com/qury53/finapp/internal/features/spendings/repository/postgres"
)

type SpendingRedisRepository struct {
	cacheClient        *core_repository_redis.Client
	spendingRepository *spending_repository_postgres.SpendingRepository
}

func NewSpendingRedisRepository(
	cacheClient *core_repository_redis.Client,
	spendingRedisRepository *spending_repository_postgres.SpendingRepository,
) *SpendingRedisRepository {
	return &SpendingRedisRepository{
		cacheClient:        cacheClient,
		spendingRepository: spendingRedisRepository,
	}
}
