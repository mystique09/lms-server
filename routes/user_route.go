package routes

import (
	"context"
	"server/app"

	"github.com/labstack/echo/v4"
)

/*
A user router struct.
*/
type UserRouter struct {
	app *app.Router
}

/*
A GetUser method for rt struct to retrive all users.
*/
func (rt *app.Router) GetUser(c echo.Context) error {
	ctx := context.Background()
	users, err := rt.DB.GetUsers(ctx)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, users)
}
