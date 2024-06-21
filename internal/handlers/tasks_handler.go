package handler

import (
	"encoding/json"
	usecases "go-api/internal/use-cases/tasks"
	"net/http"
)

type TasksHandler struct {
	GetTasksUseCase usecases.IGetTasksUseCase
}

func NewTasksHandler(getTasksUseCase usecases.IGetTasksUseCase) *TasksHandler {
	return &TasksHandler{
		GetTasksUseCase: getTasksUseCase,
	}
}

func (h *TasksHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.GetTasksUseCase.Execute()
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
