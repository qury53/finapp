package income_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
)

func (s *IncomeService) CreateIncome(
	ctx context.Context,
	income domain.Income,
) (domain.Income, error) {
	if err := income.Validate(); err != nil {
		return domain.Income{}, fmt.Errorf("service: failed to validate income: %w", err)
	}

	income, err := s.IncomeRepository.CreateIncome(ctx, income)
	if err != nil {
		return domain.Income{}, fmt.Errorf("service: failed to create income: %w", err)
	}

	return income, nil
}
