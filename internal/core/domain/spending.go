package domain

import (
	"fmt"
	"time"
	"unicode/utf8"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	"github.com/shopspring/decimal"
)

type Spending struct {
	ID          int
	Version     int
	Value       decimal.Decimal
	Title       string
	Description string
	Status      string
	Category    string
	CreatedAt   time.Time
	UserID      int
}

type SpendingPatch struct {
	Value       Nullable[string]
	Title       Nullable[string]
	Description Nullable[string]
	Status      Nullable[string]
	Category    Nullable[string]
}

func NewSpendingUninitialized(
	value decimal.Decimal,
	title string,
	description string,
	status string,
	category string,
	userID int,
) Spending {
	return Spending{
		ID:          IncomeIDUninitialized,
		Version:     IncomeVersionUninitialized,
		Value:       value,
		Title:       title,
		Description: description,
		Status:      status,
		Category:    category,
		CreatedAt:   time.Now(),
		UserID:      userID,
	}
}

func (i *Spending) Validate() error {
	if i.Value.LessThan(decimal.Zero) {
		return fmt.Errorf(
			"value cannot be negative; %w",
			core_errors.ErrInvalidArgument,
		)
	}

	maxVal, _ := decimal.NewFromString("9999999999")
	if i.Value.GreaterThan(maxVal) {
		return fmt.Errorf(
			"value cannot be more then 9999999999: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if i.Title == "" {
		return fmt.Errorf(
			"title cannot be blank string: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	titleLength := utf8.RuneCountInString(i.Title)
	if titleLength < 1 || titleLength > 100 {
		return fmt.Errorf(
			"title length should be between 1 and 100: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	descriptionLength := utf8.RuneCountInString(i.Description)
	if descriptionLength > 500 {
		return fmt.Errorf(
			"description has to be less then 500: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if _, ok := validStatus[i.Status]; !ok {
		return fmt.Errorf(
			"status is not valid, given: %s, :%w",
			i.Status,
			core_errors.ErrInvalidArgument,
		)
	}

	if i.Category == "" {
		i.Category = "Без категории"
	}

	categoryLength := utf8.RuneCountInString(i.Category)

	if categoryLength > 100 {
		return fmt.Errorf(
			"category length should be less than 100: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (r *SpendingPatch) Validate() error {
	if r.Value.Set {
		if r.Value.Value == nil {
			return fmt.Errorf(
				"value cannot be patched to null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		_, err := decimal.NewFromString(*r.Value.Value)
		if err != nil {
			return fmt.Errorf(
				"value has to be valid decimal: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf(
				"title cannot be patched to null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		titleLength := utf8.RuneCountInString(*r.Title.Value)
		if titleLength < 1 || titleLength > 100 {
			return fmt.Errorf(
				"title length has to be between 1 and 100: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Description.Set {
		if r.Description.Value == nil {
			return fmt.Errorf(
				"description cannot be patched to null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		descriptionLength := utf8.RuneCountInString(*r.Description.Value)
		if descriptionLength < 1 || descriptionLength > 100 {
			return fmt.Errorf(
				"description length has to be between 1 and 100: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Status.Set {
		if r.Status.Value == nil {
			return fmt.Errorf(
				"status has to be not null and be oncely, weekly or monthly: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		_, ok := validStatus[*r.Status.Value]
		if !ok {
			return fmt.Errorf(
				"income status has to be oncely, weekly or monthly, got : %s: %w",
				*r.Status.Value,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Category.Set {
		if r.Category.Value == nil || *r.Category.Value == "" {
			*r.Category.Value = "Без категории"
		}

		categoryLength := utf8.RuneCountInString(*r.Category.Value)
		if categoryLength > 100 {
			return fmt.Errorf(
				"category length has to be less than 100: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

func (i *Spending) ApplyPatch(patch SpendingPatch) error {
	tmp := *i

	if patch.Value.Set {
		v, err := decimal.NewFromString(*patch.Value.Value)
		if err != nil {
			return fmt.Errorf(
				"value is not valid integer: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		tmp.Value = v
	}

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = *patch.Description.Value
	}

	if patch.Status.Set {
		tmp.Status = *patch.Status.Value
	}

	if patch.Category.Set {
		tmp.Category = *patch.Category.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf(
			"failed to validate patched income: %w",
			err,
		)
	}

	*i = tmp

	return nil
}
