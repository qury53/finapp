package users_repository_redis

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *UsersRedisRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	cacheKey := r.сacheClient.GenerateCacheKey(
		"users",
		"limit",
		strconv.Itoa(*limit),
		"offset",
		strconv.Itoa(*offset),
	)

	var users []domain.User
	ok, err := r.сacheClient.Get(ctx, cacheKey, &users)
	if ok {
		if err == nil {
			return users, nil
		} else {
			slog.Warn("failed to unmarshal cache", slog.String("error", err.Error()))
		}
	}

	users, err = r.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	if err := r.сacheClient.Set(ctx, cacheKey, users); err != nil {
		slog.Warn("failed to cache users", slog.String("error", err.Error()))
	}

	return users, nil
}
