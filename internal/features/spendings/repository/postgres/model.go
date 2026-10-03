package spending_repository_postgres

import (
	"time"

	"github.com/qury53/finapp/internal/core/domain"
	"github.com/shopspring/decimal"
)

type SpendingModel struct {
	ID          int             `db:"id"`
	Version     int             `db:"version"`
	Value       decimal.Decimal `db:"value"`
	Title       string          `db:"title"`
	Description string          `db:"description"`
	Status      string          `db:"status"`
	Category    string          `db:"category"`
	CreatedAt   time.Time       `db:"created_at"`
	UserID      int             `db:"user_id"`
}

func domainFromModel(model SpendingModel) domain.Spending {
	return domain.Spending{
		ID:          model.ID,
		Version:     model.Version,
		Value:       model.Value,
		Title:       model.Title,
		Description: model.Description,
		Status:      model.Status,
		Category:    model.Category,
		CreatedAt:   model.CreatedAt,
		UserID:      model.UserID,
	}
}

func domainsFromModels(models []SpendingModel) []domain.Spending {
	spendings := make([]domain.Spending, len(models))

	for i, model := range models {
		spendings[i] = domainFromModel(model)
	}

	return spendings
}
