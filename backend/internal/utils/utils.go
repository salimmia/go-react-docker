package utils

import "github.com/salimmia/go-architecture/internal/application"

var app *application.Application

// NewHelpers sets up app config for helpers
func NewHelpers(a *application.Application) {
	app = a
}