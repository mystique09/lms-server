package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	JwtClaims struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
		jwt.StandardClaims
	}

	JwtUserPayload struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}
)

func NewJwtPayload(username, email, role string) JwtUserPayload {
	return JwtUserPayload{
		Username: username,
		Email:    email,
		Role:     role,
	}
}

func NewJwtClaims(payload JwtUserPayload) JwtClaims {
	return JwtClaims{
		Username: payload.Username,
		Email:    payload.Email,
		Role:     payload.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		},
	}
}

func NewJwtToken(payload JwtClaims, secret_key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(secret_key)
}

func GetPayloadFromJwt(token *jwt.Token) JwtUserPayload {
	var claims jwt.MapClaims = token.Claims.(jwt.MapClaims)
	return JwtUserPayload{
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
		Role:     claims["role"].(string),
	}
}
