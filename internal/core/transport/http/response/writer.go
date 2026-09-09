package transport_http_response

import (
	"net/http"

	"github.com/qury53/finapp/internal/core/domain"
)

// ResponseWriter с хранением статус кода
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Конструктор ResponseWriter
func NewResponseWriter(rw http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: rw,
		statusCode:     domain.StatusCodeUninitialized,
	}
}

// Пишет статус код и сохраняет его в поле statusCode
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

// Возвращает статус код, если он еще не установлен,
// то возвращает -1
func (rw *ResponseWriter) GetStatusCode() int {
	return rw.statusCode
}
