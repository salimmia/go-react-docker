package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig represents the application configuration.
type AppConfig struct {
	DatabaseURL        	string	`json:"database_url"`
	DatabaseDriver		string 	`json:"driver"`
	LogLevel           	string	`json:"log_level"`
	APIKey             	string 	`json:"api_key"`
	ServerPort         	string  `json:"server_port"`
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
func NewAppConfig() (*AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil{
		return nil, errors.New("failed to read .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil{
		return nil, errors.New("could not parse dbPort")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    dbHost, dbPort, dbUser, dbPassword, dbName)

	return &AppConfig{
		DatabaseURL:        dsn,
		DatabaseDriver: 	os.Getenv("DB_DRIVER"),
		LogLevel:           "info",
		APIKey:             "your_default_api_key",
		ServerPort:         os.Getenv("APP_PORT"),
		MaxConnections:     100,
		JWTSecret:          os.Getenv("JWT_SECRET"),
		EmailSMTPHost:      "smtp.example.com",
		EmailSMTPPort:      587,
		EmailSMTPUsername:  "your_smtp_username",
		EmailSMTPPassword:  "your_smtp_password",
		EnableDebugLogs:    true,
		EnableTLS:          true,
		EnableEmailSending: true,
	}, nil
}
