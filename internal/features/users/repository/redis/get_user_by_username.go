package users_repository_redis

import (
	"context"
	"log/slog"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *UsersRedisRepository) GetUserByUsername(
	ctx context.Context,
	username string,
) (domain.User, error) {
	cacheKey := r.сacheClient.GenerateCacheKey(
		"users",
		"username",
		username,
	)

	var user domain.User
	ok, err := r.сacheClient.Get(ctx, cacheKey, &user)
	if ok {
		if err == nil {
			return user, nil
		}

		slog.Warn("repository: failed to unmarshal user cache", slog.String("error", err.Error()))
	}

	user, err = r.usersRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return domain.User{}, err
	}

	if err := r.сacheClient.Set(
		ctx,
		cacheKey,
		user,
	); err != nil {
		slog.Warn("repository: failed to cache user")
	}

	return user, nil
}
