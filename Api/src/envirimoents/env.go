package envirimoents

import (
	"os"

	"github.com/joho/godotenv"
)

const returnnull = "NULL"

func GetEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		return returnnull
	}
	apiKey := os.Getenv(name)
	if apiKey == "" {
		return returnnull
	} else {
		return apiKey
	}
}

func GetPort() string {
	err := godotenv.Load()
	if err != nil {
		return returnnull
	}
	apiKey := os.Getenv("APLICATION_PORT")
	if apiKey == "" {
		return ":8080"
	} else {
		return apiKey
	}
}
