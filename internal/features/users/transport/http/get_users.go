package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

type GetUsersResponse []UserResponseDTO

// GetUsers godoc
// @Summary Gets users
// Description Gets users from db
// @Tags users
// @Produce json
// @Success 200 {object} GetUsersResponse "Successfully got users"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /users/ [get]
func (h *UsersHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	limit, offset, err := getLimitOffset(r)
	if err != nil {
		responseHandler.RespondError("transport: failed to get query params: %w", err)
		return
	}

	users, err := h.UsersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.RespondError("transport: failed to get users: %w", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetUsersResponse(dtosFromDomains(users)))
}

func getLimitOffset(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get limit query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get offset query param: %w", err)
	}

	return limit, offset, nil
}
