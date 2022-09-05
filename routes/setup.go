package routes

import (
	"io"
	"log"
	"text/template"

	//"net/http"
	"server/config"
	database "server/database/sqlc"

	//"server/frontend"
	"server/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var server Server

func Setup() Server {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.Init()

	conn := utils.SetupDB(config.DATABASE_URL)
	db := database.New(conn)
	cld, err := cloudinary.NewFromURL(config.CLD_URL)

	return Server{
		DB:  db,
		Cfg: config,
		Cld: cld,
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
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(LoggerMiddleware())
	e.Use(RateLimitMiddleware(20))
	e.Use(CorsMiddleware(server.Cfg))

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}
	/*
		// remove this if deploy to production
		e.StaticFS("frontend/build", frontend.BuildHTTPFS())

		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			//Filesystem: http.FS(frontend.BuildHTTPFS()),
			Root:  "frontend/build",
			HTML5: true,
		}))
	*/
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "indexPage", nil)
	})
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
		// classrokms relationships
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

	class_group := e.Group("/api/v1/classrooms", JwtAuthMiddleware(server.Cfg))
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

	comment_group := e.Group("/api/v1/comments", JwtAuthMiddleware(server.Cfg))
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
