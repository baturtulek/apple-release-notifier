package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error: failed to load the env file")
	}
}
