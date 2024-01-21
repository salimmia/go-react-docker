package main

import (
	"log"

	"github.com/salimmia/go-architecture/internal/application"
	"github.com/salimmia/go-architecture/internal/config"
	"github.com/salimmia/go-architecture/internal/controller"
	"github.com/salimmia/go-architecture/internal/repository"
)

func main(){
	cfg, err := config.NewAppConfig()
	if err != nil{
		log.Fatal(err)
		return
	}

	app := application.NewApplication(cfg)
	db, err := repository.ConnectDB(app.Config.DatabaseDriver, app.Config.DatabaseURL)
	if err != nil{
		log.Fatal(err)
		return
	}

	// log.Println(&app.HTTPRouter)

	repo := controller.NewRepository(app, db)
	controller.NewHandler(repo)

	router := app.HTTPRouter
	router.POST("/users/register", controller.Repo.RegistrationUser)
	router.POST("/users/update-user/{user_id}", controller.Repo.UpdateUser)
	router.POST("/users/login", controller.Repo.LogIn)

	router.SERVE(app.Config.ServerPort)
}