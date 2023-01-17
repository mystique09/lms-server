package domain

import (
	"server/internal/tokenutil"
	"time"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenResponse struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiration time.Time `json:"access_token_expiration"`
}

type RefreshTokenUsecase interface {
	ValidateRefreshToken(refreshToken string) (*tokenutil.Payload, error)
	CreateAccessToken(username string, duration time.Duration) (string, *tokenutil.Payload, error)
}
