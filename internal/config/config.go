package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDSN() string {
	// Get envs from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")

	// Create DSN (Data Source Name)
	DSN := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT)
	return DSN
}
