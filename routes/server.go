package routes

import (
	"fmt"
	"log"
	"os"

	database "server/database/sqlc"
	lmsdocs "server/lms-docs"
	"server/token"
	"server/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

type Server struct {
	store      database.Store
	router     *echo.Echo
	tokenMaker token.Maker
	cfg        utils.Config
	cld        cloudinary.Cloudinary
}

func Launch(cfg *utils.Config) {
	conn := utils.SetupDB(cfg.DBUrl)
	store := database.NewStore(conn, cfg)
	server, err := NewServer(store, cfg)

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.router.Start(cfg.Host))
}

func NewServer(store database.Store, cfg *utils.Config) (*Server, error) {
	cld, err := cloudinary.NewFromURL(cfg.CldUrl)

	if err != nil {
		log.Fatal(err.Error())
	}

	tokenMaker, err := token.NewPasetoMaker(cfg.PasetoSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot make token maker")
	}

	server := &Server{
		store:      store,
		cld:        *cld,
		cfg:        *cfg,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()

	return server, nil
}

// Serves API from routes
func (server *Server) setupRouter() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	logger := zerolog.New(os.Stdout)

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Time("time", v.StartTime.UTC()).
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	e.StaticFS("/docs", lmsdocs.BuildStaticHTTPS())
	e.Use(RateLimitMiddleware(20))
	e.Use(CorsMiddleware(&server.cfg))

	e.GET("/", server.indexRoute)
	e.POST("/api/v1/signup", server.createUser)
	e.POST("/api/v1/login", server.loginHandler)
	e.POST("/api/v1/refresh-token", server.refreshHandler)

	user_group := e.Group("/api/v1/users", server.authMiddleware)
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

	class_group := e.Group("/api/v1/classrooms", server.authMiddleware)
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

	post_group := e.Group("/api/v1/posts", server.authMiddleware)
	post_group.POST("", server.createNewPost)
	post_group.GET("/:id", server.getOnePost)
	post_group.PUT("/:id", server.updatePost)
	post_group.DELETE("/:id", server.deletePost)
	// relationships
	post_group.GET("/:id/likes", server.getAllPostLikes)
	post_group.POST("/:id/likes", server.likePost)
	post_group.DELETE("/:id", server.unlikePost)

	comment_group := e.Group("/api/v1/comments", server.authMiddleware)
	comment_group.POST("", server.createNewComment)
	comment_group.GET("/:id", server.getOneComment)
	comment_group.PUT("/:id", server.updateComment)
	comment_group.DELETE("/:id", server.deleteComment)
	// relationships
	comment_group.GET("/:id/likes", server.getAllCommentLikes)
	comment_group.POST("/:id/likes", server.likeComment)
	comment_group.DELETE("/:id", server.unlikeComment)

	trace := jaegertracing.New(e, nil)
	defer trace.Close()

	server.router = e
}
