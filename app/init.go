package app

import (
	"server/database/sqlc"
)

/*
A Router struct to hold the database connection information and other router config.
*/
type Router struct {
	DB *database.Queries
}

// A config struct to hold the app configuration.
type Config struct {
	Port   string
	DBHost string
	DBUser string
	DBPass string
	DBName string
}
