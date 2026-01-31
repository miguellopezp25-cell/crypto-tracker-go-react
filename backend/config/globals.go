package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load("config/.env")
	if err != nil {
		return err
	}
	return nil
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
