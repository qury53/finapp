package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	v := r.PathValue(key)

	if v == "" {
		return -1, fmt.Errorf("userID is not set: %w", core_errors.ErrInvalidArgument)
	}

	value, err := strconv.Atoi(v)
	if err != nil {
		return -1, fmt.Errorf("userID has to be integer: %w", core_errors.ErrInvalidArgument)
	}

	return value, nil
}
