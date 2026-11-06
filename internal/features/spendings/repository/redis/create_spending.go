package spending_repository_redis

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *SpendingRedisRepository) CreateSpending(
	ctx context.Context,
	spending domain.Spending,
) (domain.Spending, error) {
	spending, err := r.spendingRepository.CreateSpending(ctx, spending)
	if err != nil {
		return domain.Spending{}, err
	}

	cacheKey := r.cacheClient.GenerateCacheKey(
		"spendings",
		"spendingID",
		strconv.Itoa(spending.ID),
	)
	if err := r.cacheClient.Set(
		ctx,
		cacheKey,
		spending,
	); err != nil {
		slog.Debug(
			"failed to cache spending",
			slog.Int("spending id", spending.ID),
		)
	}

	return spending, nil
}
