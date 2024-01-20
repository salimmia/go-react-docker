package controller

import (
	"github.com/salimmia/go-architecture/internal/application"
	"github.com/salimmia/go-architecture/internal/repository"
	dbrepo "github.com/salimmia/go-architecture/internal/repository/dbRepo"
)

type Repository struct{
	App *application.Application
	DB repository.Database
}

func NewRepository(app *application.Application, db *repository.DB) *Repository{
	return &Repository{
		App: app,
		DB: dbrepo.NewPostgreSqlDbRepo(app, db.SQL),
	}
}

var Repo *Repository

func NewHandler(r *Repository){
	Repo = r
}