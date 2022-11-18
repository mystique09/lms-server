package routes

import (
	"net/http"
	"server/utils"

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

func CorsMiddleware(cfg *utils.Config) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			cfg.FrontendUrl,
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

func JwtAuthMiddleware(cfg *utils.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    cfg.JwtSecretKey,
		Skipper: func(c echo.Context) bool {
			return c.Request().Method == http.MethodGet
		},
	})
}

func RefreshTokenAuthMiddleware(cfg *utils.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    cfg.JwtRefreshSecretKey,
		ContextKey:    "refresh",
	})
}
