package routes

import (
	"net/http"
	"server/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	/*
	   The Route struct to hold the route information.
	*/

	/*
	   A Response struct to hold the response information.
	*/
	Response struct {
		Status int
		Body   string
	}

	AccessToken struct {
		Token string `json:"access_token"`
	}

	RefreshToken struct {
		Token string `json:"refresh_token"`
	}

	Handler struct {
		Path        string
		Action      string
		HandlerFunc echo.HandlerFunc
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}
	return nil
}
