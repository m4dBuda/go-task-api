package userRepository

import (
	model "go-api/internal/domain/model/user"

	"github.com/google/uuid"
)

type IUserRepository interface {
	FindByID(id uuid.UUID) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uuid.UUID) error
	FindAll() ([]*model.User, error)
}
