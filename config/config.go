package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const ENV_PREFIX = "eventsapi"

type Config struct {
	DBHost           string `envconfig:"DB_HOST"`
	DBPort           string `envconfig:"DB_PORT"`
	DBUsername       string `envconfig:"DB_USERNAME"`
	DBName           string `envconfig:"DB_NAME"`
	DBPassword       string `envconfig:"DB_PASSWORD"`
	PasswordHashSalt string `envconfig:"PASSWORD_HASH_SALT"`
	TokenSecret      string `envconfig:"TOKEN_SECRET"`
	Port             string `envconfig:"PORT"`
	SessionKey       string `envconfig:"SESSION_SECRET_KEY"`
}

// Recieve configuration values from env variables
func InitConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	var cfg Config
	if err := envconfig.Process(ENV_PREFIX, &cfg); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	return &cfg, nil
}
