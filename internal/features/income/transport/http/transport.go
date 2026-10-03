package income_transport_http

import (
	"context"
	"net/http"

	"github.com/qury53/finapp/internal/core/domain"
	transport_http_server "github.com/qury53/finapp/internal/core/transport/http/server"
)

type IncomeHTTPHandler struct {
	IncomeService IncomeService
}

type IncomeService interface {
	CreateIncome(ctx context.Context, income domain.Income) (domain.Income, error)
	GetIncomes(
		ctx context.Context,
		userID *int,
		period *string,
	) ([]domain.Income, error)
	GetIncome(ctx context.Context, incomeID int) (domain.Income, error)
	DeleteIncome(ctx context.Context, incomeID int) error
	PatchIncome(ctx context.Context, incomeID int, patch domain.IncomePatch) (domain.Income, error)
}

func NewIncomeHTTPHandler(service IncomeService) *IncomeHTTPHandler {
	return &IncomeHTTPHandler{
		IncomeService: service,
	}
}

func (h *IncomeHTTPHandler) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Path:    "/income/",
			Method:  http.MethodPost,
			Handler: h.CreateIncome,
		},
		{
			Path:    "/income/",
			Method:  http.MethodGet,
			Handler: h.GetIncomes,
		},
		{
			Path:    "/income/{id}",
			Method:  http.MethodGet,
			Handler: h.GetIncome,
		},
		{
			Path:    "/income/{id}",
			Method:  http.MethodDelete,
			Handler: h.DeleteIncome,
		},
		{
			Path:    "/income/{id}",
			Method:  http.MethodPatch,
			Handler: h.PatchIncome,
		},
	}
}
