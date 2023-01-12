package controller

import (
	"database/sql"
	"server/bootstrap"
	"server/domain"
	"server/internal/passwordutil"

	"github.com/labstack/echo/v4"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c echo.Context) error {
	var request domain.LoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	user, err := lc.LoginUsecase.GetUserByUsername(c, request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.ErrorResponse{Message: "User doesn't exist."})
		}
		return c.JSON(500, domain.ErrorResponse{Message: err.Error()})
	}

	if err := passwordutil.MatchPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: "Wrong password."})
	}

	access_token, access_token_payload, err := lc.LoginUsecase.CreateAccessToken(user.Username, lc.Env.AccessTokenDuration)
	if err != nil {
		return c.JSON(500, domain.ErrorResponse{Message: err.Error()})
	}

	refresh_token, refresh_token_payload, err := lc.LoginUsecase.CreateRefreshToken(user.Username, lc.Env.RefreshTokenDuration)
	if err != nil {
		return c.JSON(500, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(200, domain.SuccessResponse[domain.LoginResponse]{
		Message: "Login success.",
		Data: domain.LoginResponse{
			User:                   user,
			AccessToken:            access_token,
			AccessTokenExpiration:  access_token_payload.ExpiredAt,
			RefreshToken:           refresh_token,
			RefreshTokenExpiration: refresh_token_payload.ExpiredAt,
		},
	})
}
