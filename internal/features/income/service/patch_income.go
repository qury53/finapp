package income_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *IncomeService) PatchIncome(
	ctx context.Context,
	incomeID int,
	patch domain.IncomePatch,
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

	if err := income.ApplyPatch(patch); err != nil {
		return domain.Income{}, fmt.Errorf(
			"service: failed to apply patch: %w",
			err,
		)
	}

	if err := income.Validate(); err != nil {
		return domain.Income{}, fmt.Errorf(
			"service: failed to validate patched income: %w",
			err,
		)
	}

	income, err = s.IncomeRepository.PatchIncome(ctx, incomeID, income)
	if err != nil {
		return domain.Income{}, fmt.Errorf(
			"service: failed to patch income: %w",
			err,
		)
	}

	return income, nil
}
