package users_transport_http

import (
	"fmt"
	"net/http"
	"net/mail"
	"unicode/utf8"

	"github.com/qury53/finapp/internal/core/domain"
	core_logger "github.com/qury53/finapp/internal/core/logger"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	core_http_request "github.com/qury53/finapp/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_types "github.com/qury53/finapp/internal/core/transport/http/types"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	Password    core_http_types.Nullable[string] `json:"password" validate:"required,min=3,max=32"`
	Email       core_http_types.Nullable[string] `json:"email" validate:"required,min=3,max=255"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" validate:"omitempty,e164"`
}

type PatchUserResponse UserResponseDTO

func (r *PatchUserRequest) Validate() error {
	if r.Password.Set {
		if r.Password.Value == nil {
			return fmt.Errorf("password cannot be patched to nil: %w", core_errors.ErrInvalidArgument)
		}

		passwordLen := utf8.RuneCountInString(*r.Password.Value)
		if passwordLen < 3 || passwordLen > 32 {
			return fmt.Errorf("password length has to be between 3 and 32: %w", core_errors.ErrInvalidArgument)
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			if !core_http_request.IsE164(*r.PhoneNumber.Value) {
				return fmt.Errorf("phone number does not match e164 format: %w", core_errors.ErrInvalidArgument)
			}
		}
	}

	if r.Email.Set {
		_, err := mail.ParseAddress(*r.Email.Value)
		if err != nil {
			return fmt.Errorf("invalid email format: %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

// PatchUser godoc
// @Summary Patches user
// Description Patches user in db
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param patch body PatchUserRequest true "PatchRequestBody"
// @Success 200 {object} PatchUserResponse "Patched user response"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} transport_http_response.ErrorResponse "Conflict"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [patch]
func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("4")
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError("transport: failed to get userID: %w", err)
		return
	}

	var patchUserRequest PatchUserRequest
	if err := core_http_request.DecodeAndValidate(r, &patchUserRequest); err != nil {
		responseHandler.RespondError("transport: failed to get request: %w", err)
		return
	}

	patch := patchFromRequest(patchUserRequest)
	fmt.Println("3")
	user, err := h.UsersService.PatchUser(ctx, userID, patch)
	if err != nil {
		responseHandler.RespondError("transport: failed to patch user: %w", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, PatchUserResponse(dtoResponseFromDomain(user)))
}

func patchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		Password:    request.Password.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
		Email:       request.Email.ToDomain(),
	}
}
