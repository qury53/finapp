package income_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

type GetIncomeResponse IncomeDTOResponse

// GetIncome godoc
// @Summary Gets income
// Description Gets income from db
// @Tags income
// @Produce json
// @Param id path int true "Income ID"
// @Success 200 {object} GetIncomeResponse "Successfully got income"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /income/{id} [get]
func (h *IncomeHTTPHandler) GetIncome(w http.ResponseWriter, r *http.Request) {
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

	income, err := h.IncomeService.GetIncome(ctx, incomeID)
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get income: %w",
			err,
		)

		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetIncomeResponse(dtoFromDomain(income)))
}
