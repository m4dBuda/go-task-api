package tasks

import (
	model "go-api/internal/domain/model/tasks"
)

type IGetTasksUseCase interface {
	Execute() ([]*model.Tasks, error)
}
