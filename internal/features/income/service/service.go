package income_service

import (
	"context"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

type IncomeService struct {
	IncomeRepository IncomeRepository
}

type IncomeRepository interface {
	CreateIncome(ctx context.Context, income domain.Income) (domain.Income, error)
	GetIncomes(ctx context.Context, userID *int, period *time.Time) ([]domain.Income, error)
	GetIncome(ctx context.Context, incomeID int) (domain.Income, error)
	DeleteIncome(ctx context.Context, incomeID int) error
	PatchIncome(ctx context.Context, incomeID int, income domain.Income) (domain.Income, error)
}

func NewIncomeService(repository IncomeRepository) *IncomeService {
	return &IncomeService{
		IncomeRepository: repository,
	}
}
