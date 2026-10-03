package statistics_transport_http

import (
	"context"
	"net/http"

	"github.com/qury53/finapp/internal/core/domain"
	transport_http_server "github.com/qury53/finapp/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	StatisticsService StatisticsService
}

type StatisticsService interface {
	GetUserStatistics(ctx context.Context, userID int, period string) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(service StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		StatisticsService: service,
	}
}

func (h *StatisticsHTTPHandler) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Path:    "/statistics/{user_id}",
			Method:  http.MethodGet,
			Handler: h.GetUserStatistics,
		},
	}
}
