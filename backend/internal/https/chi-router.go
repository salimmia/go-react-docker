package https

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/salimmia/go-architecture/internal/router"
)


type ChiRouter struct{
	Router chi.Router
}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() router.Router {
	return &ChiRouter{chi.NewRouter()}
}

func (c *ChiRouter) USE(middleware func(http.Handler) http.Handler) {
	c.Router.Use(middleware)
}

func (c *ChiRouter) GET(uri string, f func (w http.ResponseWriter, r *http.Request)){
	c.Router.Get(uri, f)
}

func (c *ChiRouter) POST(uri string, f func (w http.ResponseWriter, r *http.Request)){
	c.Router.Post(uri, f)
}

func (c *ChiRouter) PUT(uri string, f func (w http.ResponseWriter, r *http.Request)){
	c.Router.Post(uri, f)
}

func (c *ChiRouter) SERVE(port string){
	fmt.Printf("Chi HTTP server running on port %v\n", port)
	http.ListenAndServe(port, c.Router)
}