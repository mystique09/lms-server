package routes

import (
	"net/http"
	"server/utils"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Username string `json:"username" validate:"required,gt=6"`
	Password string `json:"password" validate:"required,gt=6"`
}

type AuthSuccessResponse struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func (s *Server) loginHandler(c echo.Context) error {
	var payload AuthRequest

	bindErr := c.Bind(&payload)
	if bindErr != nil {
		return c.JSON(400, bindErr)
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)
	}

	user, err := s.store.GetUserByUsername(c.Request().Context(), payload.Username)
	if err != nil || user.ID == uuid.Nil {
		return c.JSON(http.StatusNotFound, USER_NOTFOUND)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		return c.JSON(http.StatusForbidden, LOGIN_FAILED)
	}

	access_token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(user.ID, user.Username, user.Email, string(user.UserRole)), 5), s.cfg.JWT_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	refresh_token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(user.ID, user.Username, user.Email, string(user.UserRole)), 60*60*7*31), s.cfg.JWT_REFRESH_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, AuthSuccessResponse{Access: access_token, Refresh: refresh_token})
}

func (s *Server) refreshToken(c echo.Context) error {
	token := c.Get("refresh").(*jwt.Token)
	user := utils.GetPayloadFromJwt(token)
	updated_user, err := s.store.GetUser(c.Request().Context(), user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	new_access_token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(updated_user.ID, updated_user.Username, updated_user.Email, string(updated_user.UserRole)), 5), s.cfg.JWT_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, AccessToken{
		Token: new_access_token,
	})
}
