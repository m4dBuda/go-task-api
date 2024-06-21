package handler

import (
	"encoding/json"
	model "go-api/internal/domain/model/user"
	usecases "go-api/internal/use-cases/user"
	"net/http"
)

type UserHandler struct {
	CreateUserUseCase usecases.ICreateUserUseCase
}

func NewUserHandler(createUserUseCase usecases.ICreateUserUseCase) *UserHandler {
	return &UserHandler{
		CreateUserUseCase: createUserUseCase,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.CreateUserUseCase.Execute(&newUser); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
