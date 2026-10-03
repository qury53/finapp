package spending_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

// DeleteSpending godoc
// @Summary Delete spending
// Description delete spending from db
// @Tags spending
// @Accept json
// @Param id path int true "Spending ID"
// @Success 204 "No content"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /spending/{id} [delete]
func (h *SpendingHTTPHandler) DeleteSpending(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	spendingID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get income id: %w",
			err,
		)

		return
	}

	if err := h.SpendingService.DeleteSpending(ctx, spendingID); err != nil {
		responseHandler.RespondError(
			"transport: failed to delete spending",
			err,
		)

		return
	}

	responseHandler.RespondNoContent()
}
