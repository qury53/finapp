package statistics_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

type GetUserStatistics StatisticsResponseDTO

var (
	defaultPeriod = "month"
)

func (h *StatisticsHTTPHandler) GetUserStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	period, err := core_http_utils.GetStringQueryParam(r, "period")
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get period",
			err,
		)

		return
	}

	if period == nil {
		period = &defaultPeriod
	}

	userID, err := core_http_utils.GetIntPathValue(r, "user_id")
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get user_id param",
			err,
		)

		return
	}

	statistics, err := h.StatisticsService.GetUserStatistics(ctx, userID, *period)
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get statistics",
			err,
		)

		return
	}

	responseHandler.RespondJSON(
		http.StatusOK,
		GetUserStatistics(dtoFromDomain(statistics)),
	)
}
