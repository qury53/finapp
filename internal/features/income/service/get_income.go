package income_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *IncomeService) GetIncome(
	ctx context.Context,
	incomeID int,
) (domain.Income, error) {
	if incomeID < 1 {
		return domain.Income{}, fmt.Errorf(
			"service: incomeID has to be positive: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	income, err := s.IncomeRepository.GetIncome(ctx, incomeID)
	if err != nil {
		return domain.Income{}, fmt.Errorf(
			"service: failed to get income: %w",
			err,
		)
	}

	return income, nil
}
