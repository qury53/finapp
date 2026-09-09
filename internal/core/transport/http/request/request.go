package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_http_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

// декодирует и валидирует запрос, если со структурой что-то не так, то
// возвращается core_http_errors.ErrInvalidArgument
func DecodeAndValidate(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("failed to decode: %v: %w", err, core_http_errors.ErrInvalidArgument)
	}

	v, ok := dest.(validatable)
	var err error

	if ok {
		err = v.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf("request validation error: %v: %w", err, core_http_errors.ErrInvalidArgument)
	}

	return nil
}
