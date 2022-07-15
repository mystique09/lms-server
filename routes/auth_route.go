package routes

import (
	"net/http"
	"server/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (rt *Route) Login(c echo.Context) error {
	var payload AuthRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", err.Error()))
	}

	if payload.Username == "" || payload.Password == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", "Username and password are required"))
	}

	user, err := rt.DB.GetUserByUsername(rt.CTX, payload.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", "User doesn't exist."))
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", "Password mismatch"))
	}

	token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(user.Username, user.Email, string(user.UserRole))), rt.Cfg.JWT_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, "", err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(1, token, ""))
}
