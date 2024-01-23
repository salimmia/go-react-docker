package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/salimmia/go-architecture/internal/application"
	"github.com/salimmia/go-architecture/internal/config"
	"github.com/salimmia/go-architecture/internal/controller"
	"github.com/salimmia/go-architecture/internal/middlewares"
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

	// router := app.HTTPRouter
	// router.Use(middleware.Heartbeat("/users/register"))

	router := chi.NewRouter()
	router.Use(middlewares.CorsMiddleware)


	router.Post("/users/register", controller.Repo.RegistrationUser)
	router.Put("/users/profile/update-user/{user_id}", controller.Repo.UpdateUser)
	router.Post("/users/login", controller.Repo.LogIn)
	router.Get("/users", controller.Repo.Users)
	router.Get("/users/profile/{user_id}", controller.Repo.User)

	// router.(app.Config.ServerPort)
	fmt.Printf("Chi HTTP server running on port %v\n", app.Config.ServerPort)
	http.ListenAndServe(":8080", router)
}