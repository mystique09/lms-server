package routes

import (
	"server/config"
	database "server/database/sqlc"
)

type (
	/*
	   The Route struct to hold the route information.
	*/
	Route struct {
		DB  *database.Queries
		Cfg config.Config
	}

	/*
	   A Response struct to hold the response information.
	*/
	Response struct {
		Status int
		Body   string
	}

	AccessToken struct {
		Token string `json:"access_token"`
	}

	RefreshToken struct {
		Token string `json:"refresh_token"`
	}
)
