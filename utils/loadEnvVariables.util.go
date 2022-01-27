package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// Loads Environment Variables
func LoadEnvironmentVariables() {
	err := godotenv.Load(AppRootPath + "/.env")

	if err != nil {
		log.Fatal("ERROR: Failed to load the env file: ", err)
	}
}
