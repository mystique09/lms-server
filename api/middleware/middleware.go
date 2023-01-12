package middleware

import (
	"log"
	"server/bootstrap"

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

func CorsMiddleware(cfg *bootstrap.Env) echo.MiddlewareFunc {
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

func Todo() {
	log.Println(authorizationHeaderKey, authorizationHeaderType, authorizationPayloadKey)
}
