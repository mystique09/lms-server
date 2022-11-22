// Package specification of Auth API.
//
// # The purpose of this API is to authenticate user.
//
// Schemes: http
// Host: localhost:5000
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package routes

import (
	"database/sql"
	"net/http"
	database "server/database/sqlc"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// swagger:model
type authRequest struct {
	// The username in the json body
	// unique: true
	// in: body
	// type: string
	Username string `json:"username" validate:"required,gt=6"`

	// The password in the json body
	// in: body
	// unique: true
	// type: string
	Password string `json:"password" validate:"required,gt=6"`
}

// swagger:parameters authUser
type AuthRequestBody struct {
	// The json payload for login handler
	// in: body
	// required: true
	Body authRequest `json:"body"`
}

// swagger:model authSuccessResponse
type authSuccessResponse struct {
	// The access token of response.
	//
	// required: true
	Access string `json:"access_token"`

	// The user object, the user that made the request.
	//
	// required: true
	User database.User `json:"user"`
}

// authentication
func (s *Server) loginHandler(c echo.Context) error {
	// The login handler
	// swagger:operation POST /api/v1/login auth authUser
	//
	// ---
	// consumes:
	// - application/json
	//
	// produces:
	// - application/json
	//
	// responses:
	//   '200':
	//	   description: login success response
	//	   schema:
	//	     type: object
	//		 	"$ref": "#/definitions/authSuccessResponse"
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
