package repository

import (
	"github.com/google/uuid"
	"github.com/salimmia/go-architecture/internal/models"
)

type Database interface{
	RegistrationUser(user *models.User) (*models.UserID, error)
	UpdateUser(user *models.User) (*models.User, error)
	GetUserById(userId uuid.UUID) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}