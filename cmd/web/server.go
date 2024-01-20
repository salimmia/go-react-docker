package main

import (
	"fmt"
	"net/http"

	"github.com/salimmia/go-architecture/internal/application"
	"github.com/salimmia/go-architecture/internal/config"
)

func main(){
	app := application.NewApplication(config.NewAppConfig())

	app.HTTPRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "up and running")
	})

	app.HTTPRouter.SERVE(fmt.Sprintf(":%d", app.Config.ServerPort))
}