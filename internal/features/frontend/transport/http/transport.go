package frontend_transport_http

import (
	"net/http"

	transport_http_server "github.com/qury53/finapp/internal/core/transport/http/server"
)

type FrontendHTTPHandler struct {
	FrontendService FrontendService
}

type FrontendService interface {
	GetMainPage() ([]byte, error)
	GetProfilePage() ([]byte, error)
}

func NewFrontendHTTPHandler(service FrontendService) *FrontendHTTPHandler {
	return &FrontendHTTPHandler{
		FrontendService: service,
	}
}

func (h *FrontendHTTPHandler) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Path:    "/",
			Handler: h.GetMainPage,
		},
		{
			Path:    "/profile/",
			Method:  http.MethodGet,
			Handler: h.GetProfilePage,
		},
	}
}
