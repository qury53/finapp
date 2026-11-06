package spending_repository_redis

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *SpendingRedisRepository) PatchSpending(
	ctx context.Context,
	spendingID int,
	spending domain.Spending,
) (domain.Spending, error) {
	updatedSpending, err := r.spendingRepository.PatchSpending(
		ctx,
		spendingID,
		spending,
	)
	if err != nil {
		return domain.Spending{}, err
	}

	spendingCacheKey := r.cacheClient.GenerateCacheKey(
		"spendings",
		"spendingID",
		strconv.Itoa(spendingID),
	)

	if err := r.cacheClient.Del(ctx, spendingCacheKey); err != nil {
		slog.Warn(
			"repository: failed to delete spending cache",
			slog.String("error", err.Error()),
		)
	}

	userID := updatedSpending.UserID

	userListCacheKey := r.cacheClient.GenerateCacheKey(
		"spendings",
		"list",
		"userID",
		strconv.Itoa(userID),
		"since",
		"all",
	)

	if err := r.cacheClient.Del(ctx, userListCacheKey); err != nil {
		slog.Warn(
			"repository: failed to delete spending list cache",
			slog.String("error", err.Error()),
		)
	}

	return updatedSpending, nil
}
