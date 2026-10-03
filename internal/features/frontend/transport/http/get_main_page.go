package frontend_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
)

func (h *FrontendHTTPHandler) GetMainPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	file, err := h.FrontendService.GetMainPage()
	if err != nil {
		responseHandler.RespondError(
			"transport: failed to get main page",
			err,
		)

		return
	}

	responseHandler.RespondFile(http.StatusOK, file)
}
