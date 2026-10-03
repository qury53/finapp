package users_transport_http

import (
	"net/http"

	"github.com/qury53/finapp/internal/core/domain"
	core_logger "github.com/qury53/finapp/internal/core/logger"
	core_http_request "github.com/qury53/finapp/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	Username    string  `json:"username" validate:"min=3,max=32" example:"cool_username"`
	Password    string  `json:"password" validate:"min=6,max=72" example:"cool_safe_password"`
	Email       string  `json:"email" validate:"min=3,max=255" example:"cool_email@gmail.com"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=7,max=15" example:"+79281001010"`
}

type CreateUserResponse UserResponseDTO

// CreateUser godoc
// @Summary creates user
// Description creates user in database
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "CreateUser request body"
// @Success 201 {object} CreateUserResponse "Successfully created user"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /users/ [post]
func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	var createUserRequest CreateUserRequest
	if err := core_http_request.DecodeAndValidate(r, &createUserRequest); err != nil {
		responseHandler.RespondError("transport: failed to decode and validate request", err)
		return
	}

	user, err := h.UsersService.CreateUser(ctx, requestToDomain(createUserRequest))
	if err != nil {
		responseHandler.RespondError("tansport: failed to create user", err)
		return
	}

	responseHandler.RespondJSON(http.StatusCreated, CreateUserResponse(dtoResponseFromDomain(user)))
}

func requestToDomain(request CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(
		request.Username,
		request.Password,
		request.Email,
		request.PhoneNumber,
	)
}
