package income_transport_http

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/qury53/finapp/internal/core/domain"
	core_logger "github.com/qury53/finapp/internal/core/logger"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	core_http_request "github.com/qury53/finapp/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_types "github.com/qury53/finapp/internal/core/transport/http/types"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
	"github.com/shopspring/decimal"
)

var incomeValidStatus = map[string]struct{}{
	"oncely":  {},
	"weekly":  {},
	"monthly": {},
}

type PatchIncomeRequest struct {
	Value       core_http_types.Nullable[string] `json:"value"`
	Title       core_http_types.Nullable[string] `json:"title" validate:"min=1,max=100"`
	Description core_http_types.Nullable[string] `json:"description" validate:"max=500"`
	Status      core_http_types.Nullable[string] `json:"status"`
}

type PatchIncomeResponse IncomeDTOResponse

func (r *PatchIncomeRequest) Validate() error {
	if r.Value.Set {
		if r.Value.Value == nil {
			return fmt.Errorf(
				"value cannot be patched to null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		_, err := decimal.NewFromString(*r.Value.Value)
		if err != nil {
			return fmt.Errorf(
				"value has to be valid decimal: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf(
				"title cannot be patched to null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		titleLength := utf8.RuneCountInString(*r.Title.Value)
		if titleLength < 1 || titleLength > 100 {
			return fmt.Errorf(
				"title length has to be between 1 and 100: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Description.Set {
		if r.Description.Value == nil {
			return fmt.Errorf(
				"description cannot be patched to null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		descriptionLength := utf8.RuneCountInString(*r.Description.Value)
		if descriptionLength < 1 || descriptionLength > 100 {
			return fmt.Errorf(
				"description length has to be between 1 and 100: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Status.Set {
		if r.Status.Value == nil {
			return fmt.Errorf(
				"status has to be not null and be oncely, weekly or monthly: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		_, ok := incomeValidStatus[*r.Status.Value]
		if !ok {
			return fmt.Errorf(
				"income status has to be oncely, weekly or monthly, got : %s: %w",
				*r.Status.Value,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

// PatchIncome godoc
// @Summary Patches income
// Description Patches income in db, income's status can be: oncely, weekly, monthly
// @Tags income
// @Accept json
// @Produce json
// @Param id path int true "Income ID"
// @Param patch body PatchIncomeRequest true "PatchRequestBody"
// @Success 200 {object} PatchIncomeResponse "Patched income response"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} transport_http_response.ErrorResponse "Conflict"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /income/{id} [patch]
func (h *IncomeHTTPHandler) PatchIncome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	incomeID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get income id",
			err,
		)

		return
	}

	var incomePatchRequest PatchIncomeRequest
	if err := core_http_request.DecodeAndValidate(r, &incomePatchRequest); err != nil {
		responseHandler.RespondError(
			"transport: faile to decode and validate income patch request",
			err,
		)

		return
	}

	patch := patchRequestToDomain(incomePatchRequest)

	income, err := h.IncomeService.PatchIncome(ctx, incomeID, patch)

	if err != nil {
		responseHandler.RespondError(
			"transport: failed to patch income",
			err,
		)

		return
	}

	responseHandler.RespondJSON(http.StatusOK, PatchIncomeResponse(dtoFromDomain(income)))
}

func patchRequestToDomain(request PatchIncomeRequest) domain.IncomePatch {
	return domain.IncomePatch{
		Value:       request.Value.ToDomain(),
		Title:       request.Title.ToDomain(),
		Description: request.Description.ToDomain(),
		Status:      request.Status.ToDomain(),
	}
}
