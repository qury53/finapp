package spending_repository_redis

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *SpendingRedisRepository) GetSpending(
	ctx context.Context,
	spendingID int,
) (domain.Spending, error) {
	cacheKey := r.cacheClient.GenerateCacheKey(
		"spendings",
		"spendingID",
		strconv.Itoa(spendingID),
	)

	var spending domain.Spending
	ok, err := r.cacheClient.Get(ctx, cacheKey, &spending)
	if ok {
		if err == nil {
			return spending, nil
		} else {
			slog.Warn(
				"repository: failed to unmarshal cache",
				slog.String("error", err.Error()),
			)
		}
	}

	spending, err = r.spendingRepository.GetSpending(ctx, spendingID)
	if err != nil {
		return domain.Spending{}, err
	}

	if err := r.cacheClient.Set(
		ctx,
		cacheKey,
		spending,
	); err != nil {
		slog.Warn(
			"repository: failed to set cache",
			slog.String("error", err.Error()),
		)
	}

	return spending, nil
}
