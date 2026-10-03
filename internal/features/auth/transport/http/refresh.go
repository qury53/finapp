package auth_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
)

// Refresh godoc
// @Summary Refresh access token
// @Description Uses refresh cookie to issue a new access token and rotate the refresh cookie
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} LoginResponse "Successfully refreshed tokens"
// @Failure 401 {object} transport_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /auth/refresh/ [post]
func (h *AuthHTTPHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get cookie",
			core_errors.ErrUnauthorized,
		)

		return
	}

	userID, valid, err := h.AuthService.ValidateRefreshToken(ctx, cookie.Value)
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to validate refresh token",
			err,
		)

		return
	}

	if !valid {
		responseHandler.RespondStatusCode(http.StatusUnauthorized)
		return
	}

	access, err := h.AuthService.GenerateAccessToken(userID)
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to generate access token",
			err,
		)

		return
	}

	refresh, err := h.AuthService.GenerateRefreshToken(userID)
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to generate refresh token",
			err,
		)

		return
	}

	h.writeTokensResponse(w, access, refresh, responseHandler)
}
