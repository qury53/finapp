package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidate(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("failed to decode: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	v, ok := dest.(validatable)
	var err error

	if ok {
		err = v.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf("request validation error: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return nil
}

func IsE164(phone string) bool {
	// Standard E.164 regex: ^\+[1-9]\d{1,14}$
	re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	return re.MatchString(phone)
}
