package routes

import (
	"context"
	"server/config"
	database "server/database/sqlc"
)

type (
	/*
	   The Route struct to hold the route information.
	*/
	Route struct {
		DB  *database.Queries
		CTX context.Context
		Cfg config.Config
	}

	/*
	   A Response struct to hold the response information.
	*/
	Response struct {
		Status int
		Body   string
	}
)
