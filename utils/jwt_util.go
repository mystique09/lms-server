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

func NewJwtClaims(username, email, role string) *JwtClaims {
	return &JwtClaims{
		Username: username,
		Email:    email,
		Role:     role,
	}
}
