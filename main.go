package main

import (
	"log"
	"os"
	"server/app"
	"server/database/sqlc"
	"server/routes"
	"server/utils"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	db := utils.SetupDB(DATABASE_URL)

	qr := database.New(db)

	rt := routes.Route{
		DB: qr,
	}

	config := app.Config{
		Port: ":" + PORT,
	}

	server := echo.New()
	server.GET("/api/v1", indexHandler)
	server.GET("/api/v1/users", rt.GetUsers)

	server.Logger.Fatal(server.Start(config.Port))
}

func indexHandler(c echo.Context) error {
	return c.String(200, "Welcome! This is my backend API for my Class Management System personal project.")
}
