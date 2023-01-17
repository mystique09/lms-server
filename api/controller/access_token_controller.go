package controller

import (
	"server/bootstrap"
	"server/domain"

	"github.com/labstack/echo/v4"
)

type AccessTokenController struct {
	AccessTokenUsecase domain.AccessTokenUsecase
	Env                *bootstrap.Env
}

func (atc *AccessTokenController) ValidateAccessToken(c echo.Context) error {
	var request domain.AccessTokenRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	access_token_payload, err := atc.AccessTokenUsecase.ValidateAccessToken(request.AccessToken)
	if err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(200, domain.SuccessResponse[domain.AccessTokenResponse]{
		Message: "Validate access token success.",
		Data: domain.AccessTokenResponse{
			AccessToken:           request.AccessToken,
			AccessTokenExpiration: access_token_payload.ExpiredAt,
		},
	})
}
