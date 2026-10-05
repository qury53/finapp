package users_repository_redis

import (
	"context"
	"log/slog"
	"strconv"
)

func (r *UsersRedisRepository) DeleteUser(
	ctx context.Context,
	userID int,
) error {
	cacheKey := r.сacheClient.GenerateCacheKey(
		"users",
		"userID",
		strconv.Itoa(userID),
	)

	if err := r.сacheClient.Del(ctx, cacheKey); err != nil {
		slog.Warn(
			"repository: failed to delete cache",
			slog.String("error", err.Error()),
		)
	}

	if err := r.usersRepository.DeleteUser(
		ctx,
		userID,
	); err != nil {
		return err
	}

	return nil
}
