package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	v := r.URL.Query().Get(key)
	if v == "" {
		return nil, nil
	}

	value, err := strconv.Atoi(v)
	if err != nil {
		return nil, fmt.Errorf("query param has to be integer: %w", err)
	}

	return &value, nil
}

func GetStringQueryParam(r *http.Request, key string) (*string, error) {
	v := r.URL.Query().Get(key)
	if v == "" {
		return nil, nil
	}

	return &v, nil
}
