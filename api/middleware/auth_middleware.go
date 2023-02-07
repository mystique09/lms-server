package middleware

import (
	"log"
	"net/http"
	"server/domain"
	"server/internal/tokenutil"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationHeaderType = "bearer"
	authorizationPayloadKey = "user"
)

func AuthMiddleware(maker tokenutil.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationHeader := c.Request().Header.Get(authorizationHeaderKey)
			if authorizationHeader == "" {
				return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "authorization header is required"})
			}

			fields := strings.Fields(authorizationHeader)

			if len(fields) < 2 {
				return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "invalid authorization header format"})
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationHeaderType {
				return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "unsupported authorization header type"})
			}

			accessToken := fields[1]
			payload, err := maker.VerifyToken(accessToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			}

			log.Println(payload)
			c.Set(authorizationPayloadKey, payload)
			return next(c)
		}
	}

}
