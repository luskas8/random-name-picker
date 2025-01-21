package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var RANDOMMER_API_KEY string

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	RANDOMMER_API_KEY = os.Getenv("RANDOMMER_API_KEY")
}
