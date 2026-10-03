package spending_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *SpendingService) PatchSpending(
	ctx context.Context,
	spendingID int,
	patch domain.SpendingPatch,
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

	if err := spending.ApplyPatch(patch); err != nil {
		return domain.Spending{}, fmt.Errorf(
			"service: failed to apply patch: %w",
			err,
		)
	}

	if err := spending.Validate(); err != nil {
		return domain.Spending{}, fmt.Errorf(
			"service: failed to validate patched spending: %w",
			err,
		)
	}

	spending, err = s.SpendingRepository.PatchSpending(ctx, spendingID, spending)
	if err != nil {
		return domain.Spending{}, fmt.Errorf(
			"service: failed to patch spending: %w",
			err,
		)
	}

	return spending, nil
}
