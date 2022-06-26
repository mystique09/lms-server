package main

import (
	"context"
	"log"
	"os"
	"server/app"
	database "server/database/sqlc"
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

	conn := utils.SetupDB(DATABASE_URL)
	db := database.New(conn)
	ctx := context.Background()

	rt := routes.Route{
		DB:  db,
		CTX: ctx,
	}

	config := app.Config{
		Port: ":" + PORT,
	}

	server := echo.New()
	server.Use(routes.LoggerMiddleware())
	server.Use(routes.RateLimitMiddleware())
	server.GET("/api/v1", indexHandler)
	server.POST("/api/v1/auth", rt.CreateUser)

	user_route := server.Group("/api/v1/users")
	{
		user_route.GET("", rt.GetUsers)
		user_route.GET("/:id", rt.GetUser)
		user_route.PUT("/:id", rt.UpdateUser)
		server.DELETE("/:id", rt.DeleteUser)
	}

	server.Logger.Fatal(server.Start(config.Port))
}

func indexHandler(c echo.Context) error {
	return c.String(200, "Welcome! This is my backend API for my Class Management System personal project.")
}
