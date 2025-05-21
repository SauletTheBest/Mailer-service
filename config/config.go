package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadAPIKey() string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MAILERSEND_API_KEY")
}
