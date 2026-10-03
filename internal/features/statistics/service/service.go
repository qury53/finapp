package statistics_service

import (
	"context"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

type StatisticsService struct {
	CacheClient    CacheClient
	IncomeReader   IncomeReader
	SpendingReader SpendingReader
	UserReader     UserReader
}

type CacheClient interface {
	Get(ctx context.Context, key string, dst any) (bool, error)
	Set(ctx context.Context, key string, value any) error
	Del(ctx context.Context, key string) error
}

type IncomeReader interface {
	GetIncomes(ctx context.Context, userID *int, since *time.Time) ([]domain.Income, error)
}

type SpendingReader interface {
	GetSpendings(ctx context.Context, userID *int, since *time.Time) ([]domain.Spending, error)
}

type UserReader interface {
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
}

func NewStatisticsService(
	cacheClient CacheClient,
	incomeReader IncomeReader,
	spendingReader SpendingReader,
	userReader UserReader,
) *StatisticsService {
	return &StatisticsService{
		CacheClient:    cacheClient,
		IncomeReader:   incomeReader,
		SpendingReader: spendingReader,
		UserReader:     userReader,
	}
}
