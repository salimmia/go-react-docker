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
	return &Application{
		Config:          config,
		HTTPRouter:      https.NewMuxRouter(), // or router.NewChiRouter(), or router.NewMuxRouter()
	}
}