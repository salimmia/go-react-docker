package main

import (
	"fmt"
	"net/http"

	router "github.com/salimmia/go-architecture/http"
)


var (
	httpRouter 		router.Router		= router.NewMuxRouter()
)

func main(){
	const port = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "up and running")
	})

	httpRouter.SERVE(port)
}