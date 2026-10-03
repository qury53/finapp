package domain

import (
	"fmt"
	"time"
	"unicode/utf8"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
	core_http_request "github.com/qury53/finapp/internal/core/transport/http/request"
)

// Доменная сущность пользователя
type User struct {
	ID          int
	Version     int
	Username    string
	Password    string
	Email       string
	PhoneNumber *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserPatch struct {
	Password    Nullable[string]
	Email       Nullable[string]
	PhoneNumber Nullable[string]
}

// Создает нового неинциализрованного пользователя, ID = -1, VERSION = -1
func NewUserUninitialized(
	username string,
	password string,
	email string,
	phone_number *string,
) User {
	t := time.Now()

	return User{
		ID:          UserIDUninitialized,
		Version:     UserVersionUninitialized,
		Username:    username,
		Password:    password,
		Email:       email,
		PhoneNumber: phone_number,
		CreatedAt:   time.Now(),
		UpdatedAt:   t,
	}
}

// Валидирует domain.User
// username длинна(3-32)
// password длинна(6-72)
// email длинна (3-255)
// phone_number (7-15)
func (u *User) Validate() error {
	usernameLength := utf8.RuneCountInString(u.Username)
	if usernameLength < 3 || usernameLength > 32 {
		return fmt.Errorf("username has to be between 3 and 32: %w", core_errors.ErrInvalidArgument)
	}

	passwordLength := utf8.RuneCountInString(u.Password)
	if passwordLength < 6 || passwordLength > 128 {
		return fmt.Errorf("password has to be between 6 and 72: %w", core_errors.ErrInvalidArgument)
	}

	emailLength := utf8.RuneCountInString(u.Email)
	if emailLength < 3 || emailLength > 255 {
		return fmt.Errorf(
			"email has to be between 3 and 255: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	phoneLength := utf8.RuneCountInString(*u.PhoneNumber)
	if phoneLength < 7 || phoneLength > 15 {
		return fmt.Errorf(
			"phone_number length has to bet between 7 and 15: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if ok := core_http_request.IsE164(*u.PhoneNumber); !ok {
		return fmt.Errorf(
			"phone_number does not match pattern: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	tmp := *u

	if patch.Password.Set {
		tmp.Password = *patch.Password.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.Password.Value
	}
	if patch.Email.Set {
		tmp.Email = *patch.Email.Value
	}

	*u = tmp

	return nil
}

func (p *UserPatch) Validate() error {
	if p.Password.Set && p.Password.Value == nil {
		return fmt.Errorf("password cannot be patched to nil: %w", core_errors.ErrInvalidArgument)
	}

	if p.Email.Set && p.Email.Value == nil {
		return fmt.Errorf("email cannot be patched to nil: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
