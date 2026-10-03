package spending_service

import (
	"context"
	"fmt"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *SpendingService) DeleteSpending(
	ctx context.Context,
	spendingID int,
) error {
	if spendingID < 1 {
		return fmt.Errorf(
			"service: spendingID has to be positive: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if err := s.SpendingRepository.DeleteSpending(ctx, spendingID); err != nil {
		return fmt.Errorf(
			"service: failed to delete spending: %w",
			err,
		)
	}

	return nil
}
