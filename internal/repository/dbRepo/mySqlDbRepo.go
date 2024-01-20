package dbrepo

import (
	"database/sql"

	"github.com/salimmia/go-architecture/repository"
)


type mySqlDbRepo struct{
	DB *sql.DB
}

func NewMySqlDbRepo(db *sql.DB) repository.Repository{
	return &mySqlDbRepo{
		DB: db,
	}
}