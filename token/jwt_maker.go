package token

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type (
	JwtUserPayload struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
		Role     string    `json:"role"`
		jwt.StandardClaims
	}
)

func NewJwtPayload(id uuid.UUID, username, email, role string, duration time.Duration) JwtUserPayload {
	return JwtUserPayload{
		ID:       id,
		Username: username,
		Email:    email,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * duration).Unix(),
		},
	}
}

func NewJwtToken(payload JwtUserPayload, secret_key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(secret_key)
}

func ExtractTokenFromHeader(header http.Header) string {
	authorization := strings.Split(header.Get("authorization"), " ")
	token := authorization[1]
	return token
}

func GetPayloadFromJwt(token *jwt.Token) JwtUserPayload {
	var claims jwt.MapClaims = token.Claims.(jwt.MapClaims)
	parsed, _ := uuid.Parse(claims["id"].(string))

	return JwtUserPayload{
		ID:       parsed,
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
		Role:     claims["role"].(string),
	}
}
