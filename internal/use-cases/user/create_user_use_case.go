package users

import (
	model "go-api/internal/domain/model/user"
	repository "go-api/internal/repositories/user"
)

var _ ICreateUserUseCase = (*CreateUserUseCase)(nil)

type CreateUserUseCase struct {
	UserRepository repository.IUserRepository
}

func NewCreateUserUseCase(userRepository repository.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepository: userRepository}
}

func (uc *CreateUserUseCase) Execute(user *model.User) error {
	return uc.UserRepository.Create(user)
}
