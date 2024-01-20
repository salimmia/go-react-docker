package application

import (
	"github.com/salimmia/go-architecture/internal/router"
	"github.com/salimmia/go-architecture/internal/usecase"
)

// Application represents the main application.
type Application struct {
    httpRouter   router.Router
    userUseCase  usecase.UserUseCase
}

// NewApplication creates a new instance of the Application.
func NewApplication(httpRouter router.Router, userUseCase usecase.UserUseCase) *Application {
    return &Application{
        httpRouter:  httpRouter,
        userUseCase: userUseCase,
    }
}

// Run starts the application.
func (app *Application) Run(port string) {
	
}