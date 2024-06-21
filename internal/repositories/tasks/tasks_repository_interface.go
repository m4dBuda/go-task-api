package taskRepository

import (
	model "go-api/internal/domain/model/tasks"

	"github.com/google/uuid"
)

type ITasksRepository interface {
	FindByID(id uuid.UUID) (*model.Tasks, error)
	Create(user *model.Tasks) error
	Update(user *model.Tasks) error
	Delete(id uuid.UUID) error
	FindAll() ([]*model.Tasks, error)
}
