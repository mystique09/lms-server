package routes

import (
	"log"
	"net/http"
	"server/config"
	database "server/database/sqlc"
	"server/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
)

var rt Route

func Setup() Route {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.Init()

	conn := utils.SetupDB(config.DATABASE_URL)
	db := database.New(conn)

	return Route{
		DB:  db,
		Cfg: config,
	}
}

func (rt *Route) IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome! This is my backend API for my Class Management System personal project.")
}

func Launch() {
	rt = Setup()

	e := echo.New()
	e.Use(LoggerMiddleware())
	e.Use(RateLimitMiddleware(20))
	e.Use(CorsMiddleware(rt.Cfg))

	e.GET("/api/v1", rt.IndexHandler)
	e.POST("/api/v1/signup", rt.CreateUser)
	e.POST("/api/v1/login", rt.Login)

	user_route := e.Group("/api/v1/users", JwtAuthMiddleware(rt.Cfg))
	{
		user_route.GET("", rt.GetUsers)
		user_route.GET("/:id", rt.GetUser)
		user_route.PUT("/:id", rt.UpdateUser)
		user_route.DELETE("/:id", rt.DeleteUser)
	}

	e.Logger.Fatal(e.Start(rt.Cfg.PORT))
}
