package utils

import "github.com/golang-jwt/jwt"

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
			ExpiresAt: 0,
		},
	}
}

func NewJwtToken(payload JwtClaims, secret_key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(secret_key)
}
