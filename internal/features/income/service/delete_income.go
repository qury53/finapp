package income_service

import (
	"context"
	"fmt"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *IncomeService) DeleteIncome(
	ctx context.Context,
	incomeID int,
) error {
	if incomeID < 1 {
		return fmt.Errorf(
			"service: incomeID has to be positive: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if err := s.IncomeRepository.DeleteIncome(ctx, incomeID); err != nil {
		return fmt.Errorf(
			"service: failed to delete income: %w",
			err,
		)
	}

	return nil
}
