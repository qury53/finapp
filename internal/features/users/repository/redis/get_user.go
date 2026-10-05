package users_repository_redis

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *UsersRedisRepository) GetUser(
	ctx context.Context,
	userID int,
) (domain.User, error) {
	cacheKey := r.сacheClient.GenerateCacheKey(
		"users",
		"userID",
		strconv.Itoa(userID),
	)

	var user domain.User
	ok, err := r.сacheClient.Get(ctx, cacheKey, &user)
	if ok {
		if err == nil {
			return user, nil
		} else {
			slog.Warn("repository: failed to unmarshal cache", "error", err)
		}
	}

	if err != nil {
		slog.Warn(
			"repository: failed to get user cache",
			slog.String("error", err.Error()),
		)
	}

	user, err = r.usersRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
