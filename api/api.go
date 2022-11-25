// Package specification LMS API.
//
// # Documentation for the LMS API.
//
// Schemes: http
// Host: localhost:5000
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	database "server/database/sqlc"
	"server/token"
	"server/utils"
	"server/web"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
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

	trace := jaegertracing.New(e, nil)
	defer trace.Close()

	logger := zerolog.New(os.Stdout)

	e.Use(LoggerMiddleware(&logger))
	e.Use(RateLimitMiddleware(20))
	e.Use(CorsMiddleware(&server.cfg))

	e.GET("/", indexPage)
	e.StaticFS("/docs", web.BuildStaticHTTPS())
	e.POST("/api/v1/signup", server.createUser)
	e.POST("/api/v1/login", server.loginHandler)
	e.POST("/api/v1/refresh-token", server.refreshHandler)

	userGroup := e.Group("/api/v1/users", server.authMiddleware)
	userGroup.GET("", server.getUsers)
	userGroup.GET("/:id", server.getUser)
	userGroup.PUT("/:id", server.updateUser)
	userGroup.DELETE("/:id", server.deleteUser)
	// classrooms relationships
	userGroup.GET("/:id/classrooms", server.getClassrooms)
	userGroup.POST("/:id/classrooms", server.joinClassroom)
	userGroup.DELETE("/:id/classrooms/:class_id", server.leaveClassroom)

	// classworks relationships
	userGroup.GET("/:id/classworks", server.getAllUserClassworks)
	userGroup.GET("/:id/classworks/:classwork_id", server.getClassworkById)

	// followers relationships
	userGroup.GET("/:id/followers", server.getFollowers)
	userGroup.GET("/:id/followings", server.getFollowings)
	userGroup.POST("/:id/followings", server.addNewFollower)
	userGroup.DELETE("/:id/followings/:following_id", server.removeFollowing)

	classGroup := e.Group("/api/v1/classrooms", server.authMiddleware)
	classGroup.GET("", server.getAllClassrooms)
	classGroup.POST("", server.createNewClassroom)
	classGroup.GET("/:id", server.getClassroom)
	classGroup.PUT("/:id", server.updateClassroom)
	classGroup.DELETE("/:id", server.deleteClassroom)
	// classworks relationships
	classGroup.GET("/:id/classworks", server.getAllClassworks)
	classGroup.POST("/:id/classworks", server.addNewClasswork)
	classGroup.DELETE("/:id/classworks/:classwork_id", server.deleteClasswork)
	// users relationships
	classGroup.GET("/:id/users", server.getClassroomUsers)
	// posts relationships
	classGroup.GET("/:id/posts", server.getClassroomPosts)

	postGroup := e.Group("/api/v1/posts", server.authMiddleware)
	postGroup.POST("", server.createNewPost)
	postGroup.GET("/:id", server.getOnePost)
	postGroup.PUT("/:id", server.updatePost)
	postGroup.DELETE("/:id", server.deletePost)
	// relationships
	postGroup.GET("/:id/likes", server.getAllPostLikes)
	postGroup.POST("/:id/likes", server.likePost)
	postGroup.DELETE("/:id", server.unlikePost)

	commentGroup := e.Group("/api/v1/comments", server.authMiddleware)
	commentGroup.POST("", server.createNewComment)
	commentGroup.GET("/:id", server.getOneComment)
	commentGroup.PUT("/:id", server.updateComment)
	commentGroup.DELETE("/:id", server.deleteComment)
	// relationships
	commentGroup.GET("/:id/likes", server.getAllCommentLikes)
	commentGroup.POST("/:id/likes", server.likeComment)
	commentGroup.DELETE("/:id", server.unlikeComment)

	server.router = e
}

func indexPage(c echo.Context) error {
	return c.HTML(http.StatusOK, `Hello, welcome! This is a google classroom clone, or an LMS website. Visit the docs <a href="/docs">Here</a>`)
}
