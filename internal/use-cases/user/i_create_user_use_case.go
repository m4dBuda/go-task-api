package users

import (
	model "go-api/internal/domain/model/user"
)

type ICreateUserUseCase interface {
	Execute(user *model.User) error
}
