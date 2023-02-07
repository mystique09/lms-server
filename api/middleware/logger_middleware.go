package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func LoggerMiddleware(logger *zerolog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogHost:    true,
		LogMethod:  true,
		LogError:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			statusCode := v.Status
			var logRow *zerolog.Event
			if v.Error != nil {
				logRow = logger.Error().Err(v.Error)
			} else {
				logRow = logger.Info()
			}

			logRow.
				Str("host", v.Host).
				Time("time", v.StartTime.UTC()).
				Str("URI", v.URI).
				Int("status", statusCode).
				Str("latency", v.Latency.String()).
				Msg("request")

			return nil
		},
	})
}
