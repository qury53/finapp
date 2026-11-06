package spending_repository_redis

import (
	"context"
	"log/slog"
	"strconv"
)

func (r *SpendingRedisRepository) DeleteSpending(
	ctx context.Context,
	spendingID int,
) error {
	if err := r.spendingRepository.DeleteSpending(ctx, spendingID); err != nil {
		return err
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

	return nil
}
