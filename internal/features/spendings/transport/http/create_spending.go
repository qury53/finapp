package spending_transport_http

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

type CreateSpendingRequest struct {
	Value       string `json:"value" validate:"required,min=0,max=999999999"`
	Title       string `json:"title" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"max=500"`
	Status      string `json:"status" validate:"required"`
	Category    string `json:"category" validate:"max=100"`
	UserID      int    `json:"user_id" validate:"required"`
}

type CreateSpendingResponse SpendingDTOResponse

// CreateSpending godoc
// @Summary creates spending, spending's status can be: oncely, weekly, monthly

// Description creates spending in db
// @Tags spending
// @Accept json
// @Produce json
// @Param request body CreateSpendingRequest true "CreateSpending request body"
// @Success 201 {object} CreateSpendingResponse "Successfully created spending"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "NotFound"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /spending/ [post]
func (h *SpendingHTTPHandler) CreateSpending(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	var spendingRequest CreateSpendingRequest

	if err := core_http_request.DecodeAndValidate(
		r,
		&spendingRequest,
	); err != nil {
		responseHandler.RespondError("failed to decode and validate request: %w", err)
		return
	}

	spending, err := domainFromRequest(spendingRequest)
	if err != nil {
		responseHandler.RespondError("transport: %w", err)
		return
	}

	spending, err = h.SpendingService.CreateSpending(ctx, spending)
	if err != nil {
		responseHandler.RespondError(
			"failed to create spending: %w",
			err,
		)

		return
	}

	responseHandler.RespondJSON(http.StatusCreated, CreateSpendingResponse(dtoFromDomain(spending)))
}

func domainFromRequest(request CreateSpendingRequest) (domain.Spending, error) {
	val, err := decimal.NewFromString(request.Value)
	if err != nil {
		return domain.Spending{}, fmt.Errorf(
			"failed to make domain from request: %w; %v",
			core_errors.ErrInvalidArgument,
			err,
		)
	}

	return domain.NewSpendingUninitialized(
		val,
		request.Title,
		request.Description,
		request.Status,
		request.Category,
		request.UserID,
	), nil
}
