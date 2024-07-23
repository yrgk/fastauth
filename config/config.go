package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string
	Port        string
	SecretKey   string
}

func GetConfig() Config {
	// dir, err := os.Getwd()
    // if err != nil {
    //     log.Fatalf("Error getting working directory: %v", err)
    // }
    // log.Println("Current working directory:", dir)

	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env file not found: %s", err)
	}

	// Getting values from environment file
	DSN := os.Getenv("DSN")
	port := os.Getenv("PORT")
	secretKey := os.Getenv("SECRET_KEY")

	return Config{
		DSN: DSN,
		Port: port,
		SecretKey: secretKey,
	}
}
