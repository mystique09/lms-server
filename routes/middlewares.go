package routes

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
		},
	)
}

func CorsMiddleware() echo.MiddlewareFunc {
	var FRONTEND_URL string = os.Getenv("FRONTEND_URL")

	if FRONTEND_URL == "" {
		log.Fatal("FRONTEND_URL is not set")
	}

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			FRONTEND_URL,
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
		MaxAge: 86400,
	})
}

func RateLimitMiddleware(limit rate.Limit) echo.MiddlewareFunc {
	return middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(limit))
}

func JwtAuthMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("JWT_SECRET_KEY")),
	})
}
