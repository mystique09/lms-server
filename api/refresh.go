package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// swagger:model refreshTokenParameter
type refreshRequest struct {
	// The refresh token in the json body
	//
	// ---
	// unique: true
	// in: body
	// type: string
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// swagger:parameters refreshTokenParameter
type RefreshRequestBody struct {
	// The json payload for login handler
	//
	// ---
	// in: body
	// required: true
	Body refreshRequest `json:"body"`
}

// swagger:model refreshSuccessResponse
type refreshSuccessResponse struct {
	// The access token of response.
	//
	// required: true
	AccessToken string `json:"access_token"`

	// The expiry of access token.
	//
	// required: true
	AccessTokenExpiresAt time.Time `json:"access_token_expiry"`
}

// authentication
func (s *Server) refreshHandler(c echo.Context) error {
	// The refreshToken handler.
	// swagger:operation POST /api/v1/refresh-token auth refreshTokenParameter
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
	//	   description: refresh-token success response
	//	   schema:
	//	     type: object
	//		 	"$ref": "#/definitions/refreshSuccessResponse"
	var payload refreshRequest

	bindErr := c.Bind(&payload)
	if bindErr != nil {
		return c.JSON(400, newError(bindErr.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, newError(err.Error()))
	}

	refreshPayload, err := s.tokenMaker.VerifyToken(payload.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, newError(err.Error()))
	}

	session, err := s.store.GetSession(c.Request().Context(), refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, newError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, newError(err.Error()))
	}

	if session.IsBlocked {
		return c.JSON(http.StatusUnauthorized, newError("blocked session"))
	}

	if session.Username != refreshPayload.Username {
		return c.JSON(http.StatusUnauthorized, newError("invalid session user"))
	}

	if session.RefreshToken != payload.RefreshToken {
		return c.JSON(http.StatusUnauthorized, "mismatched session token")
	}

	if time.Now().After(session.ExpiresAt) {
		return c.JSON(http.StatusUnauthorized, "expired session")
	}

	accessToken, accessTokenPayload, err := s.tokenMaker.CreateToken(refreshPayload.Username, s.cfg.AccessTokenDuration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newError(err.Error()))
	}

	resp := refreshSuccessResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessTokenPayload.ExpiredAt,
	}
	return c.JSON(http.StatusOK, newResponse(resp))
}
