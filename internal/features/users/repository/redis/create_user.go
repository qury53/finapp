package users_repository_redis

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *UsersRedisRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	user, err := r.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	cacheKey := r.сacheClient.GenerateCacheKey(
		"users",
		"userID",
		strconv.Itoa(user.ID),
	)

	if err := r.сacheClient.Set(
		ctx,
		cacheKey,
		user,
	); err != nil {
		slog.Warn(
			"repository: failed to cache user",
			slog.String("error", err.Error()),
		)
	}

	return user, nil
}
