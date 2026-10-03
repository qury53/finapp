package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/qury53/finapp/internal/core/domain"
	transport_http_server "github.com/qury53/finapp/internal/core/transport/http/server"
)

type AuthHTTPHandler struct {
	AuthService AuthService
}

type AuthService interface {
	Login(ctx context.Context, username string, password string) (string, string, error)
	Register(ctx context.Context, user domain.User) (domain.User, string, string, error)
	GenerateAccessToken(userID int) (string, error)
	GenerateRefreshToken(userID int) (string, error)
	GetRefreshTokenTTL() time.Duration
	ValidateRefreshToken(ctx context.Context, token string) (int, bool, error)
}

func NewAuthHTTPHandler(service AuthService) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		AuthService: service,
	}
}

func (h *AuthHTTPHandler) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Path:    "/auth/login/",
			Method:  http.MethodPost,
			Handler: h.Login,
		},
		{
			Path:    "/auth/register/",
			Method:  http.MethodPost,
			Handler: h.Register,
		},
		{
			Path:    "/auth/refresh/",
			Method:  http.MethodPost,
			Handler: h.Refresh,
		},
	}
}
