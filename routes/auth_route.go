package routes

import (
	"database/sql"
	"net/http"
	database "server/database/sqlc"
	"server/token"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type authRequest struct {
	Username string `json:"username" validate:"required,gt=6"`
	Password string `json:"password" validate:"required,gt=6"`
}

type AuthRequestBody struct {
	Body authRequest `json:"body"`
}

type authSuccessResponse struct {
	Access string        `json:"access_token"`
	User   database.User `json:"user"`
}

func (s *Server) loginHandler(c echo.Context) error {
	var payload authRequest

	bindErr := c.Bind(&payload)
	if bindErr != nil {
		return c.JSON(400, newError(bindErr.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, newError(err.Error()))
	}

	user, err := s.store.GetUserByUsername(c.Request().Context(), payload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusBadRequest, USER_NOTFOUND)
		}
		return c.JSON(http.StatusInternalServerError, newError(err.Error()))
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, LOGIN_FAILED)
	}

	accessToken, err := s.tokenMaker.CreateToken(user.Username, s.cfg.AccessTokenDuration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newError(err.Error()))
	}

	resp := authSuccessResponse{Access: accessToken, User: user}
	return c.JSON(http.StatusOK, newResponse(resp))
}

func (s *Server) refreshToken(c echo.Context) error {
	refresh := c.Get("refresh").(*jwt.Token)
	user := token.GetPayloadFromJwt(refresh)
	updated_user, err := s.store.GetUser(c.Request().Context(), user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, newError(err.Error()))
	}

	new_access_token, err := token.NewJwtToken(token.NewJwtPayload(updated_user.ID, updated_user.Username, updated_user.Email, string(updated_user.UserRole), 5), []byte(s.cfg.JwtSecretKey))
	if err != nil {
		return c.JSON(http.StatusBadRequest, newError(err.Error()))
	}

	return c.JSON(http.StatusOK, newResponse(AccessToken{
		Token: new_access_token,
	}))
}
