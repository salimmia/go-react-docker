package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of a new user.
func (m *Repository) CreateUser(ctx *gin.Context) {
    // Implementation of creating a user
    ctx.JSON(http.StatusOK, gin.H{"message": "successfully done"})
}

// GetAllUsers handles the retrieval of all users.
func (m *Repository) GetAllUsers(ctx *gin.Context) {
    // Implementation of retrieving all users
}
