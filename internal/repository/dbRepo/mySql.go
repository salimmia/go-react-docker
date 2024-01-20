package dbrepo

import (
	"database/sql"

	"github.com/salimmia/go-architecture/internal/application"
	"github.com/salimmia/go-architecture/internal/models"
	"github.com/salimmia/go-architecture/internal/repository"
)

type mySqlDbRepo struct{
	DB *sql.DB
	App *application.Application
}

func NewMySqlDbRepo(db *sql.DB, app *application.Application) repository.Database{
	return &mySqlDbRepo{
		DB: db,
		App: app,
	}
}

func (m *mySqlDbRepo) GetAllUsers() error{
	return nil
}

// CreateUser creates a new user.
func (uc *mySqlDbRepo) CreateUser(user *models.User) error {
	return nil
}