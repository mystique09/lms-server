package routes

import (
	"os"
	"server/config"
	database "server/database/sqlc"
	"server/utils"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var mockDB *database.Queries
var testRoute Route
var e *echo.Echo

func TestMain(m *testing.M) {
	godotenv.Load("./.development.env")
	cfg := config.Init()
	db := utils.SetupDB(cfg.DATABASE_URL)
	mockDB = database.New(db)
	e = echo.New()

	testRoute = Route{
		DB:  mockDB,
		Cfg: cfg,
	}

	os.Exit(m.Run())
}
