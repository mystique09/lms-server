package token

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const minSecretKey = 32

type (
	JwtUserPayload struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
		Role     string    `json:"role"`
		jwt.StandardClaims
	}
)

func NewJwtPayload(id uuid.UUID, username, email, role string, duration time.Duration) *JwtUserPayload {
	return &JwtUserPayload{
		ID:       id,
		Username: username,
		Email:    email,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * duration).Unix(),
		},
	}
}

func NewJwtToken(payload *JwtUserPayload, secret_key []byte) (string, error) {
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

type JwtMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKey {
		return nil, fmt.Errorf("invalid key sizeL must be at least %d characters", minSecretKey)
	}
	return &JwtMaker{secretKey}, nil
}

func (maker *JwtMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))

	return token, payload, err
}

func (maker *JwtMaker) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	})

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
