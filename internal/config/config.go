package config

// AppConfig represents the application configuration.
type AppConfig struct {
	DatabaseURL        	string 	`json:"database_url"`
	LogLevel           	string 	`json:"log_level"`
	APIKey             	string 	`json:"api_key"`
	ServerPort         	int    	`json:"server_port"`
	MaxConnections     	int    	`json:"max_connections"`
	JWTSecret          	string 	`json:"jwt_secret"`
	EmailSMTPHost      	string 	`json:"email_smtp_host"`
	EmailSMTPPort      	int    	`json:"email_smtp_port"`
	EmailSMTPUsername  	string 	`json:"email_smtp_username"`
	EmailSMTPPassword  	string 	`json:"email_smtp_password"`
	EnableDebugLogs    	bool   	`json:"enable_debug_logs"`
	EnableTLS          	bool   	`json:"enable_tls"`
	EnableEmailSending 	bool   	`json:"enable_email_sending"`
}

// NewAppConfig creates a new instance of AppConfig with default values.
func NewAppConfig() *AppConfig {
	return &AppConfig{
		DatabaseURL:        "your_default_database_url",
		LogLevel:           "info",
		APIKey:             "your_default_api_key",
		ServerPort:         8080,
		MaxConnections:     100,
		JWTSecret:          "mysecret",
		EmailSMTPHost:      "smtp.example.com",
		EmailSMTPPort:      587,
		EmailSMTPUsername:  "your_smtp_username",
		EmailSMTPPassword:  "your_smtp_password",
		EnableDebugLogs:    true,
		EnableTLS:          true,
		EnableEmailSending: true,
	}
}
