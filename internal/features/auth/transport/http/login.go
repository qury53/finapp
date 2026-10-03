package auth_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	core_http_request "github.com/qury53/finapp/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// Login godoc
// @Summary Login user
// Description checks if username and password are valid and returns tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "LoginUser request body"
// @Success 200 {object} LoginResponse "Successfully logged in user"
// @Failure 401 {object} transport_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /auth/login/ [post]
func (h *AuthHTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	var request LoginRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.RespondError(
			"transport: failed to decode and validate request",
			err,
		)

		return
	}

	access, refresh, err := h.AuthService.Login(ctx, request.Username, request.Password)
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to login",
			err,
		)

		return
	}

	h.writeTokensResponse(w, access, refresh, responseHandler)
}
