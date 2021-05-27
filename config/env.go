package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("cannot load .env file")
	}
	return os.Getenv(key)
}