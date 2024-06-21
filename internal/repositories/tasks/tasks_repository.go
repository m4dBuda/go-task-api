package taskRepository

import (
	model "go-api/internal/domain/model/tasks"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ ITasksRepository = &TasksRepository{}

type TasksRepository struct {
	DB *gorm.DB
}

func NewTasksRepository(db *gorm.DB) *TasksRepository {
	return &TasksRepository{DB: db}
}

func (r *TasksRepository) FindByID(id uuid.UUID) (*model.Tasks, error) {
	task := &model.Tasks{}
	result := r.DB.First(task, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (r *TasksRepository) Create(tasks *model.Tasks) error {
	result := r.DB.Create(tasks)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TasksRepository) Update(tasks *model.Tasks) error {
	result := r.DB.Save(tasks)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TasksRepository) Delete(id uuid.UUID) error {
	result := r.DB.Delete(&model.Tasks{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TasksRepository) FindAll() ([]*model.Tasks, error) {
	var tasks []*model.Tasks
	result := r.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}
