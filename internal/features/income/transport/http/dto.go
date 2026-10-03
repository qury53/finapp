package income_transport_http

import (
	"time"

	"github.com/qury53/finapp/internal/core/domain"
	"github.com/shopspring/decimal"
)

type IncomeDTOResponse struct {
	ID          int             `json:"id"`
	Version     int             `json:"version"`
	Value       decimal.Decimal `json:"value"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Status      string          `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UserID      int             `json:"user_id"`
}

func dtoFromDomain(domain domain.Income) IncomeDTOResponse {
	return IncomeDTOResponse{
		ID:          domain.ID,
		Version:     domain.Version,
		Value:       domain.Value,
		Title:       domain.Title,
		Description: domain.Description,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UserID:      domain.UserID,
	}
}

func dtosFromDomains(domains []domain.Income) []IncomeDTOResponse {
	dtos := make([]IncomeDTOResponse, len(domains))

	for i, domain := range domains {
		dtos[i] = dtoFromDomain(domain)
	}

	return dtos
}
