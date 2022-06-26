package routes

import (
	"net/http"
	"server/utils"

	"github.com/labstack/echo/v4"
)

type AuthData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (rt *Route) Login(c echo.Context) error {
	var payload AuthData

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", err.Error()))
	}

	// check if payloaf is not empty
	if payload.Username == "" || payload.Password == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", "Username and password are required"))
	}

	// check if user exists
	user, err := rt.DB.GetUserByUsername(rt.CTX, payload.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", err.Error()))
	}

	/*
		if user.Username == "" {
			return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", "User not found"))
		}
	*/
	// check if password is correct
	if payload.Password != user.Password {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", "Password mismatch"))
	}

	// create jwt token
	token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(user.Username, user.Email, string(user.UserRole))), rt.Cfg.JWT_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(1, token, ""))
}
