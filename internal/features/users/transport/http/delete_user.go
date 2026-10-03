package users_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/finapp/internal/core/logger"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	transport_http_response "github.com/qury53/finapp/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/finapp/internal/core/transport/http/utils"
)

// DeleteUser godoc
// @Summary Delete user
// Description delete user from db
// @Tags users
// @Accept json
// @Param id path int true "User ID"
// @Success 204 "No content"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [delete]
func (h *UsersHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(w, logger)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError("transport: failed to get userID: %w", core_errors.ErrInvalidArgument)
		return
	}

	if err := h.UsersService.DeleteUser(ctx, userID); err != nil {
		responseHandler.RespondError("transport: failed to delete user: %w", err)
		return
	}

	responseHandler.RespondNoContent()
}
