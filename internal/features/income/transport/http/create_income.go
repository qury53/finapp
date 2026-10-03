package income_transport_http

import (
	"fmt"
	"net/http"

	"github.com/qury53/finapp/internal/core/domain"
	core_logger "github.com/qury53/finapp/internal/core/logger"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	core_http_request "github.com/qury53/finapp/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	"github.com/shopspring/decimal"
)

type CreateIncomeRequest struct {
	Value       string `json:"value" validate:"required,min=0,max=999999999" example:"1999"`
	Title       string `json:"title" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"max=500"`
	Status      string `json:"status" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"`
}

type CreateIncomeResponse IncomeDTOResponse

// CreateIncome godoc
// @Summary creates income
// Description creates income in db, income's status can be: oncely, weekly, monthly
// @Tags income
// @Accept json
// @Produce json
// @Param request body CreateIncomeRequest true "CreateIncome request body"
// @Success 201 {object} CreateIncomeResponse "Successfully created income"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "NotFound"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /income/ [post]
func (h *IncomeHTTPHandler) CreateIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	var incomeRequest CreateIncomeRequest

	if err := core_http_request.DecodeAndValidate(
		r,
		&incomeRequest,
	); err != nil {
		responseHandler.RespondError("failed to decode and validate request: %w", err)
		return
	}

	income, err := domainFromRequest(incomeRequest)
	if err != nil {
		responseHandler.RespondError("transport: %w", err)
		return
	}

	income, err = h.IncomeService.CreateIncome(ctx, income)
	if err != nil {
		responseHandler.RespondError(
			"failed to create income: %w",
			err,
		)

		return
	}

	responseHandler.RespondJSON(http.StatusCreated, CreateIncomeResponse(dtoFromDomain(income)))
}

func domainFromRequest(request CreateIncomeRequest) (domain.Income, error) {
	val, err := decimal.NewFromString(request.Value)
	if err != nil {
		return domain.Income{}, fmt.Errorf(
			"failed to make domain from request: %w; %v",
			core_errors.ErrInvalidArgument,
			err,
		)
	}

	return domain.NewIncomeUninitialized(
		val,
		request.Title,
		request.Description,
		request.Status,
		request.UserID,
	), nil
}
