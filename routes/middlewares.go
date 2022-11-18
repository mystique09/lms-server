package routes

import (
	"log"
	"net/http"
	"server/utils"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: `[${time_rfc3339}] [${method}] ${status} ${host}${path} ${latency_human}` + "\n",
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

const (
	authorizationHeaderKey  = "authorization"
	authorizationHeaderType = "bearer"
	authorizationPayloadKey = "user"
)

func (s *Server) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == http.MethodGet {
			return next(c)
		}

		authorizationHeader := c.Request().Header.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			return c.JSON(http.StatusUnauthorized, newResponse[any](nil, "authorization header is missing"))
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			return c.JSON(http.StatusUnauthorized, newResponse[any](nil, "invalid authorization header format"))
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationHeaderType {
			return c.JSON(http.StatusUnauthorized, newResponse[any](nil, "unsupported authorization header type"))
		}

		accessToken := fields[1]
		payload, err := s.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, newResponse[any](nil, err.Error()))
		}

		log.Println(payload)
		c.Set(authorizationPayloadKey, payload)
		return next(c)
	}
}

func RefreshTokenAuthMiddleware(cfg *utils.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    cfg.JwtRefreshSecretKey,
		ContextKey:    "refresh",
	})
}
