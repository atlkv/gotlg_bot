package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

type Config struct {}

var envs = [2]string {"local", "production"}

func NewConfig(env string) (*Config, error) {
	return NewConfigWithEnv(env)
}

func NewConfigWithEnv(env string) (*Config, error) {
	if err := loadEnvironment(env); err != nil {
		return nil, err
	}
	return &Config{}, nil
}

func loadEnvironment(env string) error {
	env = strings.ToLower(env)
	var fileName string

	switch env {
		case "local":
			fileName = ".env.local"
		case "production":
			fileName = ".env.production"
		default:
			return errors.New(
				fmt.Sprintf("env not found: %s", env))
	}

	if err := godotenv.Load(fileName); err != nil {
		return errors.New(
			fmt.Sprintf("environment file: %s not found! Error message: %s", fileName, err))
	}

	return nil
}

func (c *Config) GetTelegramToken() (string, error) {
	if token, ok := os.LookupEnv("TELEGRAM_TOKEN"); ok {
		return token, nil
	}
	
	return "", errors.New("environment variable TELEGRAM_TOKEN not found in .env")
}