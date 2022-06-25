package routes

import (
	"database/sql"
	database "server/database/sqlc"
	"testing"
)

/*
Test new Route struct
*/
func TestNewRoute(t *testing.T) {
	// Create new Route
	db, err := sql.Open("pgx", "user=mystique09 password=mystique09 dbname=class-manager sslmode=disable")

	if err != nil {
		panic(err)
	}

	qr := database.New(db)

	route := Route{
		DB: qr,
	}

	// Check if Route is not nil
	if route.DB == nil {
		t.Error("NewRoute() returned nil")
	}
}
