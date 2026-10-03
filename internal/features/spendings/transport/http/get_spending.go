package spending_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

type GetSpendingResponse SpendingDTOResponse

// GetSpending godoc
// @Summary Gets spending
// Description Gets spending from db
// @Tags spending
// @Produce json
// @Param id path int true "Spending ID"
// @Success 200 {object} GetSpendingResponse "Successfully got spending"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /spending/{id} [get]
func (h *SpendingHTTPHandler) GetSpending(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	spendingID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get spending id: %w",
			err,
		)

		return
	}

	spending, err := h.SpendingService.GetSpending(ctx, spendingID)
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get spending: %w",
			err,
		)

		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetSpendingResponse(dtoFromDomain(spending)))
}
