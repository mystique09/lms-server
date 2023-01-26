package controller

import (
	"server/bootstrap"
	"server/domain"
	"server/internal/passwordutil"
	"strings"

	"github.com/labstack/echo/v4"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c echo.Context) error {
	var request domain.SignupRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	hashesPassword, err := passwordutil.Encrypt(request.Password)
	if err != nil {
		return c.JSON(500, domain.ErrorResponse{Message: err.Error()})
	}

	if err := sc.SignupUsecase.CreateUser(c, request.Username, request.Email, hashesPassword); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			if strings.Contains(err.Error(), "email") {
				return c.JSON(400, domain.ErrorResponse{Message: "Email already used by another user."})
			}

			if strings.Contains(err.Error(), "username") {
				return c.JSON(400, domain.ErrorResponse{Message: "Username already used by another user."})
			}
		}
		return c.JSON(500, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(200, domain.SuccessResponse[domain.SignupResponse]{
		Message: "Signup success.",
		Data: domain.SignupResponse{
			User: domain.NewUser{
				Username: request.Username,
				Email:    request.Email,
			},
		},
	})
}
