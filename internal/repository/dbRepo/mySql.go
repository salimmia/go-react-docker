package dbrepo

import (
	"database/sql"

	"github.com/salimmia/go-architecture/internal/models"
	"github.com/salimmia/go-architecture/internal/repository"
)

type mySqlDbRepo struct{
	DB *sql.DB
}

func NewMySqlDbRepo(db *sql.DB) repository.Repository{
	return &mySqlDbRepo{
		DB: db,
	}
}

func (m *mySqlDbRepo) GetAllUsers() error{
	return nil
}

// CreateUser creates a new user.
func (uc *mySqlDbRepo) CreateUser(user *models.User) error {
	return nil
}