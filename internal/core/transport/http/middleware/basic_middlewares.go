package transport_http_middleware

import (
	"log/slog"
	"net/http"
	"time"

	uuid "github.com/google/uuid"
	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
)

const (
	requestIDHeader = "X-Request-ID"
)

// Мидлварь отвечающий за отловку паник, вызывает RespondPanic
func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					logger := core_logger.FromContext(r.Context())
					responseHandler := transport_http_response.NewHTTPResponseHandler(
						w,
						logger,
					)

					responseHandler.RespondPanic(p, "got panic")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

// мидлварь в котором прокидывается модернизированный ResponseWriter
// и логгируется http method, status code и время обработки запроса.
func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)
			responseWriter := transport_http_response.NewResponseWriter(w)

			t := time.Now()

			logger.Debug("incoming http request")
			next.ServeHTTP(responseWriter, r)
			logger.Debug(
				"finnished handling HTTP request",
				slog.String("HTTP method", r.Method),
				slog.Int("status code", responseWriter.GetStatusCode()),
				slog.Duration("time spent", time.Since(t)),
			)
		})
	}
}

// мидлварь в котором создается UUID и записывается в заголовок под ключом X-Request-ID
func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

// мидлварь в котором логгер прокидывается в контекст
func Logger(logger *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = core_logger.WithLogger(ctx, logger)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Auth() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//... after auth feature
		})
	}
}

func CORS() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowewOrigins := map[string]struct{}{
				"http://localhost:5050": {},
			}

			origin := r.Header.Get("origin")
			if _, ok := allowewOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PATCH, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
