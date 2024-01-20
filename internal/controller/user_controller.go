package controller

import (
	"net/http"

	"github.com/salimmia/go-architecture/internal/usecase"
)

// UserController handles HTTP requests related to users.
type UserController struct {
    userUseCase usecase.UserUseCase
}

// NewUserController creates a new instance of UserController.
func NewUserController(userUseCase usecase.UserUseCase) *UserController {
    return &UserController{
        userUseCase: userUseCase,
    }
}

// CreateUser handles the creation of a new user.
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    // Implementation of creating a user
}

// GetAllUsers handles the retrieval of all users.
func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    // Implementation of retrieving all users
}
