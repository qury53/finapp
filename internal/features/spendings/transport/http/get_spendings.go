package spending_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

type GetSpendingsResponse []SpendingDTOResponse

// GetSpending godoc
// @Summary Gets spending
// Description Gets spending from db
// @Tags spending
// @Produce json
// @Success 200 {object} GetSpendingResponse "Successfully got spending"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /spending/ [get]
func (h *SpendingHTTPHandler) GetSpendings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	userID, err := core_http_utils.GetIntQueryParam(r, "user_id")
	if err != nil {
		responseHandler.RespondError("failed to get userID optional param", err)
		return
	}

	period, err := core_http_utils.GetStringQueryParam(r, "period")
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get period query param: %w",
			err,
		)

		return
	}

	spendings, err := h.SpendingService.GetSpendings(ctx, userID, period)
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get spendings: %w",
			err,
		)

		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetSpendingsResponse(dtosFromDomains(spendings)))
}
