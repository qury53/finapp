package statistics_service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	"github.com/shopspring/decimal"
)

var validPeriods = map[string]struct{}{
	"day":   {},
	"week":  {},
	"month": {},
	"year":  {},
}

func (s *StatisticsService) GetUserStatistics(
	ctx context.Context,
	userID int,
	period string,
) (domain.Statistics, error) {
	if userID < 1 {
		return domain.Statistics{}, fmt.Errorf(
			"userID has to be positive: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	_, ok := validPeriods[period]
	if !ok {
		return domain.Statistics{}, fmt.Errorf(
			"service: statistics period is not valid, got: %s: %w",
			period,
			core_errors.ErrInvalidArgument,
		)
	}

	statistics := domain.Statistics{}
	cacheKey := statisticsCacheKey(userID, period)

	ok, err := s.CacheClient.Get(
		ctx,
		cacheKey,
		&statistics,
	)

	if ok {
		if err != nil {
			fmt.Println("failed to unmarshal cache")
		} else {
			fmt.Println("got from cache")
			return statistics, nil
		}
	}

	since, err := getTimeFromPeriod(&period)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			"service: %w",
			err,
		)
	}

	incomes, err := s.IncomeReader.GetIncomes(ctx, &userID, since)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			"service: failed to get incomes for user with ID: %d: %w",
			userID,
			err,
		)
	}

	spendings, err := s.SpendingReader.GetSpendings(ctx, &userID, since)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			"service: failed to get incomes for user with ID: %d: %w",
			userID,
			err,
		)
	}

	var totalIncome decimal.Decimal = decimal.Zero
	for _, income := range incomes {
		totalIncome = totalIncome.Add(income.Value)
	}

	slog.Debug("totalIncome", slog.String("income", totalIncome.String()))

	var totalSpending decimal.Decimal = decimal.Zero
	for _, spending := range spendings {
		totalSpending = totalSpending.Add(spending.Value)
	}

	slog.Debug("totalSpending", slog.String("spending", totalSpending.String()))

	statistics.TotalIncome = totalIncome
	statistics.Balance = totalIncome.Sub(totalSpending)
	statistics.TotalSpending = totalSpending

	var leftPercentage decimal.Decimal

	if totalIncome.IsZero() {
		leftPercentage = decimal.Zero
	} else {
		leftPercentage = statistics.
			Balance.
			Div(totalIncome).
			Mul(decimal.NewFromInt(100))
	}

	statistics.MoneyLeftPrecentage = leftPercentage.StringFixed(2)
	statistics.MoneySpentPrecentage = decimal.NewFromInt(100).Sub(leftPercentage).StringFixed(2)
	err = s.CacheClient.Set(ctx, cacheKey, statistics)
	fmt.Println(err)
	return statistics, nil
}

func getTimeFromPeriod(period *string) (*time.Time, error) {
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
				"failed to convert period to time: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return since, nil
}

func statisticsCacheKey(userID int, period string) string {
	return fmt.Sprintf("statistics:userID:%d:period:%s", userID, period)
}
