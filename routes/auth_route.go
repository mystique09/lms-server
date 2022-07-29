package routes

import (
	"net/http"
	"server/utils"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Username string `json:"username" validate:"required,gt=6"`
	Password string `json:"password" validate:"required,gt=6"`
}

type AuthSuccessResponse struct {
	Message string `json:"message"`
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func (s *Server) loginHandler(c echo.Context) error {
	var payload AuthRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse("", "One field might be missing, fill in the missing fields."))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	user, err := s.DB.GetUserByUsername(c.Request().Context(), payload.Username)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewResponse(nil, "User doesn't exist."))
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		return c.JSON(http.StatusForbidden, utils.NewResponse(nil, "Incorrect username or password."))
	}

	access_token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(user.ID, user.Username, user.Email, string(user.UserRole)), 5), s.Cfg.JWT_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	refresh_token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(user.ID, user.Username, user.Email, string(user.UserRole)), 60*60*7*31), s.Cfg.JWT_REFRESH_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(AuthSuccessResponse{Message: "Logged in.", Access: access_token, Refresh: refresh_token}, ""))
}

func (s *Server) refreshToken(c echo.Context) error {
	token := c.Get("refresh").(*jwt.Token)
	user := utils.GetPayloadFromJwt(token)
	updated_user, err := s.DB.GetUser(c.Request().Context(), user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	new_access_token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(updated_user.ID, updated_user.Username, updated_user.Email, string(updated_user.UserRole)), 5), s.Cfg.JWT_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, utils.NewResponse(AccessToken{Token: new_access_token}, ""))
}
