package envirimoents

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(name string, defaultValue string) string {
	err := godotenv.Load()
	if err != nil {
		return defaultValue
	}
	apiKey := os.Getenv(name)
	if apiKey == "" {
		return defaultValue
	} else {
		return apiKey
	}
}

func GetPort() string {
	err := godotenv.Load()
	if err != nil {
		return ":8080"
	}
	apiKey := os.Getenv("APLICATION_PORT")
	if apiKey == "" {
		return ":8080"
	} else {
		return apiKey
	}
}
