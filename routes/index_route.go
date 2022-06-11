package routes

import (
	// import the postgres driver
	database "server/database/sqlc"
)

/*
The Route struct to hold the route information.
*/
type Route struct {
	DB *database.Queries
}

/*
A Response struct to hold the response information.
*/
type Response struct {
	Status int
	Body   string
}
