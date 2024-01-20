package config

// AppConfig represents the application configuration.
type AppConfig struct {
	Debug bool
}

// NewAppConfig creates a new instance of AppConfig with default values.
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Debug: false,
	}
}
