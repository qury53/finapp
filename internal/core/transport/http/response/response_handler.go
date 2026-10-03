package transport_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

// Структура хэндлера, который хэндлит ответы пользователю
type HTTPResponseHandler struct {
	w      http.ResponseWriter
	logger *core_logger.Logger
}

// коструктор ResponseHandler
func NewHTTPResponseHandler(
	w http.ResponseWriter,
	logger *core_logger.Logger,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		w:      w,
		logger: logger,
	}
}

func (h *HTTPResponseHandler) RespondFile(statusCode int, file []byte) {
	h.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.w.WriteHeader(statusCode)
	_, err := h.w.Write(file)
	if err != nil {
		h.logger.Error("failed to write file", slog.String("error", err.Error()))
	}
}

// возвращает статус код и json
func (h *HTTPResponseHandler) RespondJSON(statusCode int, data any) {
	h.w.WriteHeader(statusCode)
	if err := json.NewEncoder(h.w).Encode(data); err != nil {
		h.logger.Error("failed to respond json")
	}
}

func (h *HTTPResponseHandler) RespondStatusCode(statusCode int) {
	h.w.WriteHeader(statusCode)
}

// возвращает сообщение о панике и статус код 500
func (h *HTTPResponseHandler) RespondPanic(p any, message string) {
	h.logger.Error(message, "got panic", p)
	h.errorRespond(
		http.StatusInternalServerError,
		message,
		fmt.Errorf("internal server error"),
	)
}

// возвращает NoContent
func (h *HTTPResponseHandler) RespondNoContent() {
	h.w.WriteHeader(http.StatusNoContent)
}

// проводит валидацию ошибки по статус коду и пишет ответ
func (h *HTTPResponseHandler) RespondError(msg string, err error) {
	var (
		logFunc    func(msg string, args ...any)
		statusCode int
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = slog.Debug
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = slog.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = slog.Warn
	case errors.Is(err, core_errors.ErrUnauthorized):
		statusCode = http.StatusUnauthorized
		logFunc = slog.Debug
	default:
		statusCode = http.StatusInternalServerError
		logFunc = slog.Error
	}

	logFunc(msg, "error", err)
	h.errorRespond(statusCode, msg, err)
}

// возвращает ответ ошибки
func (h *HTTPResponseHandler) errorRespond(statusCode int, msg string, err error) {
	h.w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{
		Error:   err.Error(),
		Message: msg,
	}

	if err := json.NewEncoder(h.w).Encode(&errorResponse); err != nil {
		h.logger.Error("failed to write response", "err", err, "msg", msg)
	}
}
