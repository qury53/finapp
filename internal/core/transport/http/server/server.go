package transport_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_middleware "github.com/qury53/finapp/internal/core/transport/http/middleware"
)

// структура HTTP сервера
type HTTPServer struct {
	mux         *http.ServeMux
	config      *Config
	logger      *core_logger.Logger
	middlewares []transport_http_middleware.Middleware
}

// Конструктор HTTP сервера
func NewHTTPServer(
	config *Config,
	logger *core_logger.Logger,
	middlewares ...transport_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:    http.NewServeMux(),
		config: config,
		logger: logger,
	}
}

// Регистрирует роутры и применяет к ним мидлвари
func (h *HTTPServer) RegisterAPIRouters(routers ...APIRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.APIVersion)
		r := router.WithMiddleware()
		h.mux.Handle(prefix+"/", http.StripPrefix(prefix, r))
	}
}

// Запуск HTTP сервера
func (h *HTTPServer) Run(ctx context.Context) error {
	mux := transport_http_middleware.ChainMiddleware(h.mux, h.middlewares...)
	server := &http.Server{
		Handler: mux,
		Addr:    h.config.HTTPAddr,
	}

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(ctx, h.config.ShutdownTime)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to shut down server: %w", err)
		}

	case err := <-ch:
		if err != nil {
			return fmt.Errorf("server error: %w", err)
		}
	}

	return nil
}
