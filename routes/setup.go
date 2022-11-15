package routes

import (
	"log"

	"server/config"
	database "server/database/sqlc"
	"server/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

type Server struct {
	store  database.Store
	router *echo.Echo
	cfg    config.Config
	cld    cloudinary.Cloudinary
}

func NewServer(store database.Store) *Server {
	cfg := config.Init()

	cld, err := cloudinary.NewFromURL(cfg.CLD_URL)

	if err != nil {
		log.Fatal(err.Error())
	}

	server := &Server{store: store, cld: *cld, cfg: cfg}
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	trace := jaegertracing.New(e, nil)
	defer trace.Close()

	e.Use(LoggerMiddleware())
	e.Use(RateLimitMiddleware(20))
	e.Use(CorsMiddleware(server.cfg))

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "indexPage", nil)
	})

	e.GET("/api/v1", server.indexRoute)
	e.POST("/api/v1/signup", server.createUser)
	e.POST("/api/v1/login", server.loginHandler)
	e.POST("/api/v1/refresh", server.refreshToken, RefreshTokenAuthMiddleware(server.cfg))

	user_group := e.Group("/api/v1/users", JwtAuthMiddleware(server.cfg))
	{
		user_group.GET("", server.getUsers)
		user_group.GET("/:id", server.getUser)
		user_group.PUT("/:id", server.updateUser)
		user_group.DELETE("/:id", server.deleteUser)
		// classrooms relationships
		user_group.GET("/:id/classrooms", server.getClassrooms)
		user_group.POST("/:id/classrooms", server.joinClassroom)
		user_group.DELETE("/:id/classrooms/:class_id", server.leaveClassroom)

		// classworks relationships
		user_group.GET("/:id/classworks", server.getAllUserClassworks)
		user_group.GET("/:id/classworks/:classwork_id", server.getClassworkById)

		// followers relationships
		user_group.GET("/:id/followers", server.getFollowers)
		user_group.GET("/:id/followings", server.getFollowings)
		user_group.POST("/:id/followings", server.addNewFollower)
		user_group.DELETE("/:id/followings/:following_id", server.removeFollowing)
	}

	class_group := e.Group("/api/v1/classrooms", JwtAuthMiddleware(server.cfg))
	{
		class_group.GET("", server.getAllClassrooms)
		class_group.POST("", server.createNewClassroom)
		class_group.GET("/:id", server.getClassroom)
		class_group.PUT("/:id", server.updateClassroom)
		class_group.DELETE("/:id", server.deleteClassroom)
		// classworks relationships
		class_group.GET("/:id/classworks", server.getAllClassworks)
		class_group.POST("/:id/classworks", server.addNewClasswork)
		class_group.DELETE("/:id/classworks/:classwork_id", server.deleteClasswork)
		// users relationships
		class_group.GET("/:id/users", server.getClassroomUsers)
		// posts relationships
		class_group.GET("/:id/posts", server.getClassroomPosts)
	}

	post_group := e.Group("/api/v1/posts", JwtAuthMiddleware(server.cfg))
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

	comment_group := e.Group("/api/v1/comments", JwtAuthMiddleware(server.cfg))
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

	server.router = e

	return server
}

func Launch() {
	err := godotenv.Load()
	cfg := config.Init()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := utils.SetupDB(cfg.DATABASE_URL)
	store := database.NewStore(conn, cfg)
	server := NewServer(store)

	log.Fatal(server.router.Start(cfg.PORT))
}
