package domain

import (
	"server/internal/tokenutil"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User                   User      `json:"user"`
	AccessToken            string    `json:"access_token"`
	AccessTokenExpiration  time.Time `json:"access_token_expiration"`
	RefreshToken           string    `json:"refresh_token"`
	RefreshTokenExpiration time.Time `json:"refresh_token_expiration"`
}

type LoginUsecase interface {
	GetUserByUsername(c echo.Context, username string) (User, error)
	CreateAccessToken(username string, uid uuid.UUID, duration time.Duration) (string, *tokenutil.Payload, error)
	CreateRefreshToken(username string, uid uuid.UUID, duration time.Duration) (string, *tokenutil.Payload, error)
}
