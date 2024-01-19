package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type ginRouter struct{}

var (
	ginDispatcher = gin.Default()
)

func NewGinRouter() Router{
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, f func (w http.ResponseWriter, r *http.Request)){
	ginDispatcher.GET(uri, func(ctx *gin.Context) {
		f(ctx.Writer, ctx.Request)
	})
}

func (*ginRouter) POST(uri string, f func (w http.ResponseWriter, r *http.Request)){
	ginDispatcher.POST(uri, func(ctx *gin.Context) {
		f(ctx.Writer, ctx.Request)
	})
}

func (*ginRouter) SERVE(port string){
	fmt.Printf("Gin HTTP server running on port %v\n", port)

	http.ListenAndServe(port, ginDispatcher)
}