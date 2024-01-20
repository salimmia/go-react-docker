package dbrepo

import (
	"database/sql"

	"github.com/salimmia/go-architecture/internal/models"
	"github.com/salimmia/go-architecture/internal/repository"
)

type postgreSqlDbRepo struct{
	DB *sql.DB
}

func NewPostgreSqlDbRepo(db *sql.DB) repository.Repository{
	return &postgreSqlDbRepo{
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