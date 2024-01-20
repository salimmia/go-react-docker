package main

import (
	"fmt"
	"log"
	"net/http"

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

	repo := controller.NewRepository(app, db)
	controller.NewHandler(repo)

	app.HTTPRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "up and running")
	})

	app.HTTPRouter.SERVE(app.Config.ServerPort)
}