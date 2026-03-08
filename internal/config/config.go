package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
	DatabaseURL string
}


func Load() *Config{
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("database url not found")
		return nil
	}

	return &Config{
		Port: port,
		DatabaseURL: databaseURL,
	}
}
