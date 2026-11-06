package spending_repository_redis

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

func (r *SpendingRedisRepository) GetSpendings(
	ctx context.Context,
	userID *int,
	since *time.Time,
) ([]domain.Spending, error) {
	var parts []string

	if userID != nil {
		parts = append(parts, "userID", strconv.Itoa(*userID))
	} else {
		parts = append(parts, "userID", "all")
	}

	if since != nil {
		parts = append(parts, "since", since.Format(time.RFC3339))
	} else {
		parts = append(parts, "since", "all")
	}

	cacheKey := r.cacheClient.GenerateCacheKey(
		append([]string{"spendings", "list"}, parts...)...,
	)

	var spendings []domain.Spending
	ok, err := r.cacheClient.Get(ctx, cacheKey, &spendings)
	if ok {
		if err == nil {
			return spendings, nil
		}

		slog.Warn(
			"repository: failed to unmarshal cache",
			slog.String("error", err.Error()),
		)
	}

	spendings, err = r.spendingRepository.GetSpendings(ctx, userID, since)
	if err != nil {
		return nil, err
	}

	if err := r.cacheClient.Set(
		ctx,
		cacheKey,
		spendings,
	); err != nil {
		slog.Warn(
			"repository: failed to set cache",
			slog.String("error", err.Error()),
		)
	}

	return spendings, nil
}
