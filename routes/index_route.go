package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (rt *Route) indexRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome! This is my backend API for my Class Management System personal project.")
}
