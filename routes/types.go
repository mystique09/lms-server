package routes

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
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
		return echo.NewHTTPError(http.StatusBadRequest, newResponse(nil, err.Error()))
	}
	return nil
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

/*
A function to create a new response.
*/
func newResponse(data interface{}, err string) Response {
	return Response{
		Data:  data,
		Error: err,
	}
}
