package routes

import (
	"log"
	"server/config"
	database "server/database/sqlc"
	"server/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

func addNewHandlerToGroup(g *echo.Group, handlers []Handler) {
	for _, handler := range handlers {
		if handler.Action == "GET" {
			g.GET(handler.Path, handler.HandlerFunc)
		} else if handler.Action == "POST" {
			g.POST(handler.Path, handler.HandlerFunc)
		} else if handler.Action == "UPDATE" {
			g.PUT(handler.Path, handler.HandlerFunc)
		} else if handler.Action == "PATCH" {
			g.PATCH(handler.Path, handler.HandlerFunc)
		} else {
			g.DELETE(handler.Path, handler.HandlerFunc)
		}
	}
}

func Launch() {
	rt = Setup()

	e := echo.New()
	e.Use(LoggerMiddleware())
	e.Use(RateLimitMiddleware(20))
	e.Use(CorsMiddleware(rt.Cfg))

	e.GET("/api/v1", rt.indexRoute)
	e.POST("/api/v1/signup", rt.createUser)
	e.POST("/api/v1/login", rt.loginRoute)
	e.POST("/api/v1/refresh", rt.refreshToken, RefreshTokenAuthMiddleware(rt.Cfg))

	user_group := e.Group("/api/v1/users", JwtAuthMiddleware(rt.Cfg))
	{
		user_group.GET("", rt.getUsers)
		user_group.GET("/:id", rt.getUser)
		user_group.PUT("/:id", rt.updateUser)
		user_group.DELETE("/:id", rt.deleteUser)
		user_group.GET("/:id/classrooms", rt.getClassrooms)
		user_group.GET("/:id/classrooms", rt.getFollowers)
		user_group.GET("/:id/classrooms", rt.getFollowing)
		user_group.POST("/:id/following", rt.createNewFollowing)
		user_group.DELETE("/:id/following/:id", rt.removeFollowing)
	}

	class_group := e.Group("/api/v1/classrooms", JwtAuthMiddleware(rt.Cfg))
	{
		class_group.GET("/:id", rt.getClassroom)
		class_group.POST("", rt.createNewClassroom)
		class_group.PUT("/:id", rt.updateClassroom)
		class_group.DELETE("/:id", rt.deleteClassroom)
		class_group.GET("/:id/users", rt.getClassroomUsers)
		class_group.GET("/:id/posts", rt.getClassroomPosts)
	}

	e.Logger.Fatal(e.Start(rt.Cfg.PORT))
}
