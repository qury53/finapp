package income_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

type GetIncomesResponse []IncomeDTOResponse

// GetIncome godoc
// @Summary Gets incomes
// Description Gets incomes from db
// @Tags income
// @Produce json
// @Success 200 {object} GetIncomeResponse "Successfully got incomes"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /income/ [get]
func (h *IncomeHTTPHandler) GetIncomes(w http.ResponseWriter, r *http.Request) {
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
			"transport: failed to get period query param",
			err,
		)

		return
	}

	incomes, err := h.IncomeService.GetIncomes(ctx, userID, period)
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get incomes: %w",
			err,
		)

		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetIncomesResponse(dtosFromDomains(incomes)))
}
