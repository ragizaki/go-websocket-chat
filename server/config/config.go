package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}
	return nil
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
