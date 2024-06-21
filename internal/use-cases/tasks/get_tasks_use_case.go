package tasks

import (
	model "go-api/internal/domain/model/tasks"
	repository "go-api/internal/repositories/tasks"
)

var _ IGetTasksUseCase = (*GetTasksUseCase)(nil)

type GetTasksUseCase struct {
	TasksRepository repository.ITasksRepository
}

func NewGetTasksUseCase(tasksRepository repository.ITasksRepository) *GetTasksUseCase {
	return &GetTasksUseCase{TasksRepository: tasksRepository}
}

func (taskUseCase *GetTasksUseCase) Execute() ([]*model.Tasks, error) {
	return taskUseCase.TasksRepository.FindAll()
}
