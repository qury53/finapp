package spending_transport_http

import (
	"time"

	"github.com/qury53/finapp/internal/core/domain"
	"github.com/shopspring/decimal"
)

type SpendingDTOResponse struct {
	ID          int             `json:"id"`
	Version     int             `json:"version"`
	Value       decimal.Decimal `json:"value"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Status      string          `json:"status"`
	Category    string          `json:"category"`
	CreatedAt   time.Time       `json:"created_at"`
	UserID      int             `json:"user_id"`
}

func dtoFromDomain(domain domain.Spending) SpendingDTOResponse {
	return SpendingDTOResponse{
		ID:          domain.ID,
		Version:     domain.Version,
		Value:       domain.Value,
		Title:       domain.Title,
		Description: domain.Description,
		Status:      domain.Status,
		Category:    domain.Category,
		CreatedAt:   domain.CreatedAt,
		UserID:      domain.UserID,
	}
}

func dtosFromDomains(domains []domain.Spending) []SpendingDTOResponse {
	dtos := make([]SpendingDTOResponse, len(domains))

	for i, domain := range domains {
		dtos[i] = dtoFromDomain(domain)
	}

	return dtos
}
