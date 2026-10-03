package users_transport_http

import (
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

type UserRequestDTO struct {
	ID          int
	Version     int
	Username    string
	Password    string
	Email       string
	PhoneNumber *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserResponseDTO struct {
	ID          int
	Version     int
	Username    string
	Email       string
	PhoneNumber *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func domainFromRequestDTO(dto UserRequestDTO) domain.User {
	return domain.User{
		ID:          dto.ID,
		Version:     dto.Version,
		Username:    dto.Username,
		Password:    dto.Password,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
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

func dtosFromDomains(domains []domain.User) []UserResponseDTO {
	dtos := make([]UserResponseDTO, len(domains))

	for i, domain := range domains {
		dtos[i] = dtoResponseFromDomain(domain)
	}

	return dtos
}
