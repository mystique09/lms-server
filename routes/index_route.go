package routes

import (
	// import the postgres driver
	"context"
	"server/config"
	database "server/database/sqlc"
)

/*
The Route struct to hold the route information.
*/
type Route struct {
	DB  *database.Queries
	CTX context.Context
	Cfg config.Config
}

/*
A Response struct to hold the response information.
*/
type Response struct {
	Status int
	Body   string
}
