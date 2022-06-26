package config

import (
	"log"
	"os"
)

// A config struct to hold the app configuration.
type Config struct {
	PORT           string
	DATABASE_URL   string
	FRONTEND_URL   string
	JWT_SECRET_KEY []byte
}

func Init() Config {
	var PORT string = os.Getenv("PORT")
	var DATABASE_URL string = os.Getenv("DATABASE_URL")
	var FRONTEND_URL string = os.Getenv("FRONTEND_URL")
	var JWT_SECRET_KEY string = os.Getenv("JWT_SECRET_KEY")

	if PORT == "" {
		PORT = "8080"
	}

	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	if FRONTEND_URL == "" {
		log.Fatal("FRONTEND_URL is not set")
	}

	return Config{
		PORT:           ":" + PORT,
		DATABASE_URL:   DATABASE_URL,
		FRONTEND_URL:   FRONTEND_URL,
		JWT_SECRET_KEY: []byte(JWT_SECRET_KEY),
	}
}
