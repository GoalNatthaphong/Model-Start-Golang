package env

import (
	"Goal/configs/logs"
	"errors"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	DB_HOST            string
	DB_USER            string
	DB_PASSWORD        string
	DB_NAME            string
	DB_PORT            string
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("GO_ENV")
	if env != "staging" && env != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			logs.Error("Error loading .env file")
			panic(err)
		}
	}

	requiredEnvVars := []string{
		"DB_HOST",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_PORT",
	}

	for _, envVar := range requiredEnvVars {
		if value := os.Getenv(envVar); value == "" {
			return nil, errors.New("missing required environment variable: " + envVar)
		}
	}

	return &Config{
		DB_HOST:                   	os.Getenv("DB_HOST"),
		DB_USER:                   	os.Getenv("DB_USER"),
		DB_PASSWORD:               	os.Getenv("DB_PASSWORD"),
		DB_NAME:                   	os.Getenv("DB_NAME"),
		DB_PORT:                   	os.Getenv("DB_PORT"),
	}, nil
}