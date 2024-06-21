package userRepository

import (
	model "go-api/internal/domain/model/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ IUserRepository = &UserRepository{}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByID(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	result := r.DB.First(user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Update(user *model.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	result := r.DB.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
