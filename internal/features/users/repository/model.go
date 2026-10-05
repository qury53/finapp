package users_repository

import (
	"time"

	"github.com/qury53/finapp/internal/core/domain"
)

type UserModel struct {
	ID          int       `db:"id"`
	Version     int       `db:"version"`
	Username    string    `db:"username"`
	Password    string    `db:"password"`
	Email       string    `db:"email"`
	PhoneNumber *string   `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func DomainFromModel(model UserModel) domain.User {
	return domain.User{
		ID:          model.ID,
		Version:     model.Version,
		Username:    model.Username,
		Password:    model.Password,
		Email:       model.Email,
		PhoneNumber: model.PhoneNumber,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func DomainsFromModels(models []UserModel) []domain.User {
	domains := make([]domain.User, len(models))
	for i, model := range models {
		domains[i] = DomainFromModel(model)
	}

	return domains
}
