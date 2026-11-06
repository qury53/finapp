package users_repository_redis

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *UsersRedisRepository) PatchUser(
	ctx context.Context,
	userID int,
	patchedUser domain.User,
) (domain.User, error) {
	user, err := r.usersRepository.PatchUser(
		ctx,
		userID,
		patchedUser,
	)

	if err != nil {
		return domain.User{}, err
	}

	cacheKey := r.сacheClient.GenerateCacheKey(
		"users",
		"userID",
		strconv.Itoa(userID),
	)

	if err := r.сacheClient.Del(ctx, cacheKey); err != nil {
		slog.Warn(
			"repository: failed to remove expired cache",
			slog.String("error", err.Error()),
		)
	}

	if err := r.сacheClient.Set(ctx, cacheKey, user); err != nil {
		slog.Warn(
			"repository: failed to set cache",
			slog.String("error", err.Error()),
		)
	}

	return user, nil
}
