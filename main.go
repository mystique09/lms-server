package main

import (
	"context"
	"log"
	"net/http"
	"server/config"
	database "server/database/sqlc"
	"server/routes"
	"server/utils"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Setup() routes.Route {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.Init()

	conn := utils.SetupDB(config.DATABASE_URL)
	db := database.New(conn)
	ctx := context.Background()

	return routes.Route{
		DB:  db,
		CTX: ctx,
		Cfg: config,
	}
}

func main() {
	rt := Setup()

	e := echo.New()
	e.Use(routes.LoggerMiddleware())
	e.Use(routes.RateLimitMiddleware(20))
	e.Use(routes.CorsMiddleware(rt.Cfg))

	e.GET("/api/v1", indexHandler)
	e.POST("/api/v1/signup", rt.CreateUser)
	e.POST("/api/v1/login", rt.Login)

	user_route := e.Group("/api/v1/users", routes.JwtAuthMiddleware(rt.Cfg))
	{
		user_route.GET("", rt.GetUsers)
		user_route.GET("/:id", rt.GetUser)
		user_route.PUT("/:id", rt.UpdateUser)
		user_route.DELETE("/:id", rt.DeleteUser)
	}

	e.Logger.Fatal(e.Start(rt.Cfg.PORT))
}

func indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome! This is my backend API for my Class Management System personal project.")
}
