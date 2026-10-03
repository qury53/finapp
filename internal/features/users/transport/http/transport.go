package users_transport_http

import (
	"context"
	"net/http"

	"github.com/qury53/finapp/internal/core/domain"
	transport_http_server "github.com/qury53/finapp/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	UsersService UsersService
}

type UsersService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, userID int) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	DeleteUser(ctx context.Context, userID int) error
	PatchUser(ctx context.Context, userID int, patch domain.UserPatch) (domain.User, error)
}

func NewUsersHTTPHandler(service UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		UsersService: service,
	}
}

func (h *UsersHTTPHandler) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Path:    "/users/",
			Method:  http.MethodPost,
			Handler: h.CreateUser,
		},
		{
			Path:    "/users/{id}",
			Method:  http.MethodGet,
			Handler: h.GetUser,
		},
		{
			Path:    "/users/",
			Method:  http.MethodGet,
			Handler: h.GetUsers,
		},
		{
			Path:    "/users/{id}",
			Method:  http.MethodDelete,
			Handler: h.DeleteUser,
		},
		{
			Path:    "/users/{id}",
			Method:  http.MethodPatch,
			Handler: h.PatchUser,
		},
	}
}
