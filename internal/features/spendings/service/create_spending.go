package spending_service

import (
	"context"
	"fmt"

	"github.com/qury53/finapp/internal/core/domain"
)

func (s *SpendingService) CreateSpending(
	ctx context.Context,
	spending domain.Spending,
) (domain.Spending, error) {
	if err := spending.Validate(); err != nil {
		return domain.Spending{}, fmt.Errorf("service: failed to validate spending: %w", err)
	}

	spending, err := s.SpendingRepository.CreateSpending(ctx, spending)
	if err != nil {
		return domain.Spending{}, fmt.Errorf("service: failed to create spending: %w", err)
	}

	return spending, nil
}
