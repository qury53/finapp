package income_repository_postgres

import (
	"time"

	"github.com/qury53/finapp/internal/core/domain"
	"github.com/shopspring/decimal"
)

type IncomeModel struct {
	ID          int             `db:"id"`
	Version     int             `db:"version"`
	Value       decimal.Decimal `db:"value"`
	Title       string          `db:"title"`
	Description string          `db:"description"`
	Status      string          `db:"status"`
	CreatedAt   time.Time       `db:"created_at"`
	UserID      int             `db:"user_id"`
}

func domainFromModel(model IncomeModel) domain.Income {
	return domain.Income{
		ID:          model.ID,
		Version:     model.Version,
		Value:       model.Value,
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UserID:      model.UserID,
	}
}

func domainsFromModels(models []IncomeModel) []domain.Income {
	incomes := make([]domain.Income, len(models))

	for i, model := range models {
		incomes[i] = domainFromModel(model)
	}

	return incomes
}
