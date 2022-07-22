package routes

import (
	"log"
	"server/config"
	database "server/database/sqlc"
	"server/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var server Server

func Setup() Server {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.Init()

	conn := utils.SetupDB(config.DATABASE_URL)
	db := database.New(conn)

	return Server{
		DB:  db,
		Cfg: config,
	}
}

func addNewHandleserveroGroup(g *echo.Group, handlers []Handler) {
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
	server = Setup()

	e := echo.New()
	e.Use(LoggerMiddleware())
	e.Use(RateLimitMiddleware(20))
	e.Use(CorsMiddleware(server.Cfg))

	e.GET("/api/v1", server.indexRoute)
	e.POST("/api/v1/signup", server.createUser)
	e.POST("/api/v1/login", server.loginHandler)
	e.POST("/api/v1/refresh", server.refreshToken, RefreshTokenAuthMiddleware(server.Cfg))

	user_group := e.Group("/api/v1/users", JwtAuthMiddleware(server.Cfg))
	{
		user_group.GET("", server.getUsers)
		user_group.GET("/:id", server.getUser)
		user_group.PUT("/:id", server.updateUser)
		user_group.DELETE("/:id", server.deleteUser)
		// relationships
		user_group.GET("/:id/classrooms", server.getClassrooms)
		user_group.POST("/:id/classrooms", server.joinClassroom)
		user_group.DELETE("/:id/classrooms/:class_id", server.leaveClassroom)
		user_group.GET("/:id/followers", server.getFollowers)
		user_group.GET("/:id/followings", server.getFollowings)
		user_group.POST("/:id/followings", server.addNewFollower)
		user_group.DELETE("/:id/followings/:following_id", server.removeFollowing)
	}

	class_group := e.Group("/api/v1/classrooms", JwtAuthMiddleware(server.Cfg))
	{
		class_group.GET("", server.getAllClassrooms)
		class_group.POST("", server.createNewClassroom)
		class_group.GET("/:id", server.getClassroom)
		class_group.PUT("/:id", server.updateClassroom)
		class_group.DELETE("/:id", server.deleteClassroom)
		// relationships
		class_group.GET("/:id/users", server.getClassroomUsers)
		class_group.GET("/:id/posts", server.getClassroomPosts)
	}

	post_group := e.Group("/api/v1/posts", JwtAuthMiddleware(server.Cfg))
	{
		post_group.POST("", server.createNewPost)
		post_group.GET("/:id", server.getOnePost)
		post_group.PUT("/:id", server.updatePost)
		post_group.DELETE("/:id", server.deletePost)
		// relationships
		post_group.GET("/:id/likes", server.getAllPostLikes)
		post_group.POST("/:id/likes", server.likePost)
		post_group.DELETE("/:id", server.unlikePost)
	}

	comment_group := e.Group("/api/v1/posts", JwtAuthMiddleware(server.Cfg))
	{
		comment_group.POST("", server.createNewComment)
		comment_group.GET("/:id", server.getOneComment)
		comment_group.PUT("/:id", server.updateComment)
		comment_group.DELETE("/:id", server.deleteComment)
		// relationships
		comment_group.GET("/:id/likes", server.getAllCommentLikes)
		comment_group.POST("/:id/likes", server.likeComment)
		comment_group.DELETE("/:id", server.unlikeComment)
	}

	e.Logger.Fatal(e.Start(server.Cfg.PORT))
}
