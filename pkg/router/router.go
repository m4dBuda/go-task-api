package router

import (
	handler "go-api/internal/handlers"
	tasksRepository "go-api/internal/repositories/tasks"
	userRepository "go-api/internal/repositories/user"
	"go-api/internal/use-cases/tasks"
	users "go-api/internal/use-cases/user"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *mux.Router {
	// Repositórios para a entidade User
	userRepository := userRepository.NewUserRepository(db)
	createUserUseCase := users.NewCreateUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(createUserUseCase)

	// Repositórios para a entidade Tasks
	tasksRepository := tasksRepository.NewTasksRepository(db)
	getTasksUseCase := tasks.NewGetTasksUseCase(tasksRepository)
	tasksHandler := handler.NewTasksHandler(getTasksUseCase)

	router := mux.NewRouter()

	// Rotas para User
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	// Rotas para Tasks
	router.HandleFunc("/tasks", tasksHandler.GetTasks).Methods("GET")

	return router
}
