package dbrepo

import (
	"database/sql"

	"github.com/salimmia/go-architecture/repository"
)

type postgreSqlDbRepo struct{
	DB *sql.DB
}

func NewPostgreSqlDbRepo(db *sql.DB) repository.Repository{
	return &postgreSqlDbRepo{
		DB: db,
	}
}