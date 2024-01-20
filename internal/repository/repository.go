package repository

import "github.com/salimmia/go-architecture/internal/models"

type Repository interface{
	GetAllUsers() error
	CreateUser(user *models.User) error
}