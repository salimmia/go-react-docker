package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController handles HTTP requests related to users.
type UserController struct {}

// NewUserController creates a new instance of UserController.
func NewUserController() *UserController {
    return &UserController{}
}

// CreateUser handles the creation of a new user.
func (uc *UserController) CreateUser(context *gin.Context) {
    // Implementation of creating a user
    context.JSON(http.StatusOK, gin.H{"message": "successfully done"})
}

// GetAllUsers handles the retrieval of all users.
func (uc *UserController) GetAllUsers(context *gin.Context) {
    // Implementation of retrieving all users
}
