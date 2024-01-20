package application

import (
	"github.com/salimmia/go-architecture/internal/config"
	"github.com/salimmia/go-architecture/internal/controller"
	"github.com/salimmia/go-architecture/internal/delivery/https"
	"github.com/salimmia/go-architecture/internal/router"
)

// Application represents the main application.
type Application struct {
	Config          *config.AppConfig
	UserController  *controller.UserController
	HTTPRouter      router.Router
}

// NewApplication creates a new instance of the Application with the given configuration.
func NewApplication(config *config.AppConfig) *Application {
	return &Application{
		Config:          config,
		UserController:  controller.NewUserController(),
		HTTPRouter:      https.NewGinRouter(), // or router.NewChiRouter(), or router.NewMuxRouter()
	}
}