package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APPENV  string
	APPPORT string
	APPURL  string

	DB_HOST    string
	DB_PORT    string
	DB_USER    string
	DB_PASS    string
	DB_NAME    string
	DB_SSLMODE string

	JWT_SECRET string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found, relying on environment variables")
	}
	config := &Config{
		APPENV:  os.Getenv("APPENV"),
		APPPORT: os.Getenv("APPPORT"),
		APPURL:  os.Getenv("APPURL"),

		DB_HOST:    os.Getenv("DB_HOST"),
		DB_PORT:    os.Getenv("DB_PORT"),
		DB_USER:    os.Getenv("DB_USER"),
		DB_PASS:    os.Getenv("DB_PASS"),
		DB_NAME:    os.Getenv("DB_NAME"),
		DB_SSLMODE: os.Getenv("DB_SSLMODE"),

		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
	if err := validateConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}

func validateConfig(config *Config) error {
	if config.APPENV == "production" {
		if config.DB_NAME == "" {
			return fmt.Errorf("DB_NAME is required in production environment")
		}
		if config.DB_PASS == "" {
			return fmt.Errorf("DB_PASS is required in production environment")
		}
		if config.JWT_SECRET == "" {
			return fmt.Errorf("JWTSecret is required in production environment")
		}
		if len(config.JWT_SECRET) < 32 {
			return fmt.Errorf("JWTSecret should be at least 32 characters long for better security")
		}
		if config.DB_SSLMODE == "disable" {
			log.Println("Warning: DBSSLMode is set to 'disable' in production environment. Consider enabling SSL for better security.")
		}
	}
	return nil
}
