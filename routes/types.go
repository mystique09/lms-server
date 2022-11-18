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
		return echo.NewHTTPError(http.StatusBadRequest, newResponse[any](nil, err.Error()))
	}
	return nil
}

type Response[T any] struct {
	Data  T      `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

/*
A function to create a new response.
*/
func newResponse[T any](data T, err string) Response[T] {
	return Response[T]{
		Data:  data,
		Error: err,
	}
}
