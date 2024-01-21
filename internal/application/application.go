package application

import (
	"github.com/salimmia/go-architecture/internal/config"
	"github.com/salimmia/go-architecture/internal/https"
	"github.com/salimmia/go-architecture/internal/router"
)

// Application represents the main application.
type Application struct {
	Config          *config.AppConfig
	HTTPRouter      router.Router
}

// NewApplication creates a new instance of the Application with the given configuration.
func NewApplication(config *config.AppConfig) *Application {
	var router router.Router

	switch config.Router{
	case "mux":
		router = https.NewMuxRouter()
	case "gin":
		router = https.NewGinRouter()
	default:
		router = https.NewChiRouter()
	}

	return &Application{
		Config:          config,
		HTTPRouter:   	router, // or router.NewChiRouter(), or router.NewMuxRouter()
	}
}