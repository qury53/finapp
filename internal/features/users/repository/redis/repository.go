package users_repository_redis

import (
	core_repository_redis "github.com/qury53/finapp/internal/core/repository/redis"
	user_repository_postgres "github.com/qury53/finapp/internal/features/users/repository/postgres"
)

type UsersRedisRepository struct {
	сacheClient     *core_repository_redis.Client
	usersRepository *user_repository_postgres.UsersRepository
}
