package spending_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *SpendingService) GetSpending(
	ctx context.Context,
	spendingID int,
) (domain.Spending, error) {
	if spendingID < 1 {
		return domain.Spending{}, fmt.Errorf(
			"service: spendingID has to be positive: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	spending, err := s.SpendingRepository.GetSpending(ctx, spendingID)
	if err != nil {
		return domain.Spending{}, fmt.Errorf(
			"service: failed to get spending: %w",
			err,
		)
	}

	return spending, nil
}
