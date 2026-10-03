package auth_transport_http

import (
	"net/http"
	"time"

	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
)

func (h *AuthHTTPHandler) writeTokensResponse(
	w http.ResponseWriter,
	accessToken string,
	refreshToken string,
	responseHandler *transport_http_response.HTTPResponseHandler,
) {
	h.writeRefresh(w, refreshToken)

	responseHandler.RespondJSON(
		http.StatusOK,
		LoginResponse{
			AccessToken: accessToken,
			TokenType:   "Bearer",
		},
	)
}

func (h *AuthHTTPHandler) writeRefresh(
	w http.ResponseWriter,
	refreshToken string,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(h.AuthService.GetRefreshTokenTTL() / time.Second),
	})

}
