package spending_transport_http

import (
	"context"
	"net/http"

	"github.com/qury53/finapp/internal/core/domain"
	transport_http_server "github.com/qury53/finapp/internal/core/transport/http/server"
)

type SpendingHTTPHandler struct {
	SpendingService SpendingService
}

type SpendingService interface {
	CreateSpending(ctx context.Context, spending domain.Spending) (domain.Spending, error)
	GetSpendings(
		ctx context.Context,
		userID *int,
		period *string,
	) ([]domain.Spending, error)
	GetSpending(ctx context.Context, spendingID int) (domain.Spending, error)
	DeleteSpending(ctx context.Context, spendingID int) error
	PatchSpending(ctx context.Context, spendingID int, patch domain.SpendingPatch) (domain.Spending, error)
}

func NewSpendingHTTPHandler(service SpendingService) *SpendingHTTPHandler {
	return &SpendingHTTPHandler{
		SpendingService: service,
	}
}

func (h *SpendingHTTPHandler) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Path:    "/spending/",
			Method:  http.MethodPost,
			Handler: h.CreateSpending,
		},
		{
			Path:    "/spending/",
			Method:  http.MethodGet,
			Handler: h.GetSpendings,
		},
		{
			Path:    "/spending/{id}",
			Method:  http.MethodGet,
			Handler: h.GetSpending,
		},
		{
			Path:    "/spending/{id}",
			Method:  http.MethodDelete,
			Handler: h.DeleteSpending,
		},
		{
			Path:    "/spending/{id}",
			Method:  http.MethodPatch,
			Handler: h.PatchSpending,
		},
	}
}
