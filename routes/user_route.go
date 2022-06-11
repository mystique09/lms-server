package routes

import (
	"context"

	"github.com/labstack/echo/v4"
)

/*
A GetUser method for rt struct to retrive all users.
*/
func (rt *Route) GetUsers(c echo.Context) error {
	ctx := context.Background()
	users, err := rt.DB.GetUsers(ctx)

	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, users)
}
