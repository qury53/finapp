package income_service

import (
	"context"
	"fmt"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (s *IncomeService) GetIncomes(
	ctx context.Context,
	userID *int,
	period *string,
) ([]domain.Income, error) {
	if userID != nil && *userID < 1 {
		return nil, fmt.Errorf(
			"service: userID has to be positive: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	now := time.Now()

	var since *time.Time
	if period != nil {
		switch *period {
		case "year":
			sd := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
			since = &sd
		case "month":
			sd := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			since = &sd
		case "week":
			weekday := int(now.Weekday())
			daysToSubtract := weekday - 1
			if weekday == 0 {
				daysToSubtract = 6
			}

			dayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			sd := dayStart.AddDate(0, 0, -daysToSubtract)
			since = &sd
		case "day":
			sd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
			since = &sd

		default:
			return nil, fmt.Errorf(
				"service: invalid period format, got: %s: %w",
				*period,
				core_errors.ErrInvalidArgument,
			)
		}
	}
	incomes, err := s.IncomeRepository.GetIncomes(ctx, userID, since)
	if err != nil {
		return nil, fmt.Errorf(
			"service: failed to get incomes: %w",
			err,
		)
	}

	return incomes, nil
}
