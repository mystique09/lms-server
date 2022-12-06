package api

import (
	"log"
	"net/http"
	"server/utils"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

func LoggerMiddleware(logger *zerolog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Time("time", v.StartTime.UTC()).
				Str("URI", v.URI).
				Int("status", v.Status).
				Int64("latency", v.Latency.Milliseconds()).
				Msg("request")

			return nil
		},
	})
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
		if authorizationHeader == "" {
			return c.JSON(http.StatusUnauthorized, newError("authorization header is missing"))
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			return c.JSON(http.StatusUnauthorized, newError("invalid authorization header format"))
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationHeaderType {
			return c.JSON(http.StatusUnauthorized, newError("unsupported authorization header type"))
		}

		accessToken := fields[1]
		payload, err := s.tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, newError(err.Error()))
		}

		log.Println(payload)
		c.Set(authorizationPayloadKey, payload)
		return next(c)
	}
}
