package dbrepo

import (
	"database/sql"

	"github.com/salimmia/go-architecture/internal/application"
	"github.com/salimmia/go-architecture/internal/models"
	"github.com/salimmia/go-architecture/internal/repository"
)

type postgreSqlDbRepo struct{
	App *application.Application
	DB *sql.DB
}

func NewPostgreSqlDbRepo(app *application.Application, db *sql.DB) repository.Database{
	return &postgreSqlDbRepo{
		App: app,
		DB: db,
	}
}

func (m *postgreSqlDbRepo) GetAllUsers() error{
	return nil
}

// CreateUser creates a new user.
func (uc *postgreSqlDbRepo) CreateUser(user *models.User) error {
	return nil
}