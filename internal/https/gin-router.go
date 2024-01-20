package https

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salimmia/go-architecture/internal/router"
)


type ginRouter struct{
	Engine *gin.Engine
}

func NewGinRouter() router.Router {
	return &ginRouter{
		Engine: gin.Default(),
	}
}

func (r *ginRouter) GET(uri string, f func (w http.ResponseWriter, r *http.Request)){
	r.Engine.GET(uri, func(ctx *gin.Context) {
		f(ctx.Writer, ctx.Request)
	})
}

func (r *ginRouter) POST(uri string, f func (w http.ResponseWriter, r *http.Request)){
	r.Engine.POST(uri, func(ctx *gin.Context) {
		f(ctx.Writer, ctx.Request)
	})
}

func (r *ginRouter) PUT(uri string, f func (w http.ResponseWriter, r *http.Request)){
	r.Engine.PUT(uri, func(ctx *gin.Context) {
		f(ctx.Writer, ctx.Request)
	})
}

func (r *ginRouter) SERVE(port string){
	fmt.Printf("Gin HTTP server running on port %v\n", port)

	http.ListenAndServe(port, r.Engine)
}