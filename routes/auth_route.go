package routes

import (
	"database/sql"
	"net/http"
	"server/token"

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
		return c.JSON(400, newResponse[any](nil, bindErr.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	user, err := s.store.GetUserByUsername(c.Request().Context(), payload.Username)
	if err == sql.ErrNoRows || user.ID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, USER_NOTFOUND)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		return c.JSON(http.StatusForbidden, LOGIN_FAILED)
	}

	access_token, err := token.NewJwtToken(token.NewJwtPayload(user.ID, user.Username, user.Email, string(user.UserRole), 5), []byte(s.cfg.JwtSecretKey))
	if err != nil {
		return c.JSON(402, newResponse[any](nil, err.Error()))
	}

	refresh_token, err := token.NewJwtToken(token.NewJwtPayload(user.ID, user.Username, user.Email, string(user.UserRole), 30), []byte(s.cfg.JwtRefreshSecretKey))
	if err != nil {
		return c.JSON(402, newResponse[any](nil, err.Error()))
	}

	return c.JSON(http.StatusOK, newResponse(AuthSuccessResponse{Access: access_token, Refresh: refresh_token}, ""))
}

func (s *Server) refreshToken(c echo.Context) error {
	refresh := c.Get("refresh").(*jwt.Token)
	user := token.GetPayloadFromJwt(refresh)
	updated_user, err := s.store.GetUser(c.Request().Context(), user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse[any](nil, err.Error()))
	}

	new_access_token, err := token.NewJwtToken(token.NewJwtPayload(updated_user.ID, updated_user.Username, updated_user.Email, string(updated_user.UserRole), 5), []byte(s.cfg.JwtSecretKey))
	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse[any](nil, err.Error()))
	}

	return c.JSON(http.StatusOK, newResponse(AccessToken{
		Token: new_access_token,
	}, ""))
}
