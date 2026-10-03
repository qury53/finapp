package auth_transport_http

import (
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

type UserResponseDTO struct {
	ID          int
	Version     int
	Username    string
	Email       string
	PhoneNumber *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func dtoResponseFromDomain(domain domain.User) UserResponseDTO {
	return UserResponseDTO{
		ID:          domain.ID,
		Version:     domain.Version,
		Username:    domain.Username,
		Email:       domain.Email,
		PhoneNumber: domain.PhoneNumber,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
