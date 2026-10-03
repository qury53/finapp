package users_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

type GetUserResponse UserResponseDTO

// GetUser godoc
// @Summary Gets user
// Description Gets user from db
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} GetUserResponse "Successfully got user"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError("transport: failed to get userID: %w", err)
		return
	}

	user, err := h.UsersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.RespondError("transport: failed to get user: %w", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetUserResponse(dtoResponseFromDomain(user)))
}
