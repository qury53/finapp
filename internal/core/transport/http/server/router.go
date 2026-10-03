package transport_http_server

import (
	"fmt"
	"net/http"

	transport_http_middleware "github.com/qury53/finapp/internal/core/transport/http/middleware"
)

type APIVersion string

// APIRouter роутен с мидлварями и версионированием
type APIRouter struct {
	*http.ServeMux
	APIVersion  APIVersion
	Middlewares []transport_http_middleware.Middleware
}

// Конструктор APIRouter-а
func NewAPIRouter(
	apiVersion APIVersion,
	middlewares ...transport_http_middleware.Middleware,
) *APIRouter {
	return &APIRouter{
		ServeMux:    http.NewServeMux(),
		APIVersion:  apiVersion,
		Middlewares: middlewares,
	}
}

// применяет мидлвари
func (r *APIRouter) WithMiddleware() http.Handler {
	return transport_http_middleware.ChainMiddleware(r, r.Middlewares...)
}

// Регистрирует роуты и применяет мидлвари(WithMiddleware)
func (r *APIRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		h := route.WithMiddleware()
		r.Handle(pattern, h)
	}
}
