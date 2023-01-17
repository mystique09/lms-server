package controller

import (
	"server/bootstrap"
	"server/domain"

	"github.com/labstack/echo/v4"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (rtc *RefreshTokenController) RefreshToken(c echo.Context) error {
	var request domain.RefreshTokenRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	refresh_token_payload, err := rtc.RefreshTokenUsecase.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	access_token, access_token_payload, err := rtc.RefreshTokenUsecase.CreateAccessToken(refresh_token_payload.Username, rtc.Env.AccessTokenDuration)
	if err != nil {
		return c.JSON(500, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(200, domain.SuccessResponse[domain.RefreshTokenResponse]{
		Message: "Refresh token success.",
		Data: domain.RefreshTokenResponse{
			AccessToken:           access_token,
			AccessTokenExpiration: access_token_payload.ExpiredAt,
		},
	})
}
