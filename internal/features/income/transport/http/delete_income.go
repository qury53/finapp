package income_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

// DeleteIncome godoc
// @Summary Delete income
// Description delete income from db
// @Tags income
// @Accept json
// @Param id path int true "Income ID"
// @Success 204 "No content"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /income/{id} [delete]
func (h *IncomeHTTPHandler) DeleteIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	incomeID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get income id: %w",
			err,
		)

		return
	}

	if err := h.IncomeService.DeleteIncome(ctx, incomeID); err != nil {
		responseHandler.RespondError(
			"transport: failed to delete income",
			err,
		)

		return
	}

	responseHandler.RespondNoContent()
}
