package spending_service

import (
	"context"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

type SpendingService struct {
	SpendingRepository SpendingRepository
}

type SpendingRepository interface {
	CreateSpending(ctx context.Context, income domain.Spending) (domain.Spending, error)
	GetSpendings(
		ctx context.Context,
		userID *int,
		since *time.Time,
	) ([]domain.Spending, error)
	GetSpending(ctx context.Context, spendingID int) (domain.Spending, error)
	DeleteSpending(ctx context.Context, spendingID int) error
	PatchSpending(ctx context.Context, spendingID int, spending domain.Spending) (domain.Spending, error)
}

func NewSpendingService(repository SpendingRepository) *SpendingService {
	return &SpendingService{
		SpendingRepository: repository,
	}
}
