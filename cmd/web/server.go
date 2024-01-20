package main

import (
	"fmt"
	"net/http"

	"github.com/salimmia/go-architecture/internal/delivery/https"
	"github.com/salimmia/go-architecture/internal/router"
)


var (
	httpRouter 		router.Router		= https.NewGinRouter()
)

func main(){
	const port = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "up and running")
	})

	httpRouter.SERVE(port)
}