package https

import (
	"fmt"
	"net/http"

	"github.com/go-mux/mux"
	"github.com/salimmia/go-architecture/internal/router"
)


type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() router.Router{
	return &muxRouter{}
}

func (c *muxRouter) USE(middleware func(http.Handler) http.Handler) {
	c.USE(middleware)
}

func (*muxRouter) GET(uri string, f func (w http.ResponseWriter, r *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func (w http.ResponseWriter, r *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) PUT(uri string, f func (w http.ResponseWriter, r *http.Request)){
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}

func (*muxRouter) SERVE(port string){
	fmt.Printf("Mux HTTP server running on port %v\n", port)
	http.ListenAndServe(port, muxDispatcher)
}