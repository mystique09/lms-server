package domain

import (
	"server/internal/tokenutil"
	"time"
)

type AccessTokenRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type AccessTokenResponse struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiration time.Time `json:"access_token_expiration"`
}

type AccessTokenUsecase interface {
	ValidateAccessToken(AccessToken string) (*tokenutil.Payload, error)
}
