package transport_http_server

import (
	"net/http"

	transport_http_middleware "github.com/qury53/finapp/internal/core/transport/http/middleware"
)

// Route структура роута
type Route struct {
	Method      int
	Path        string
	Handler     http.HandlerFunc
	Middlewares []transport_http_middleware.Middleware
}

// WithMiddleware применяет мидлвари к роуту
func (r *Route) WithMiddleware() http.Handler {
	return transport_http_middleware.ChainMiddleware(r.Handler, r.Middlewares...)
}
