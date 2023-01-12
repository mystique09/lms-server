package middleware

import (
	"log"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationHeaderType = "bearer"
	authorizationPayloadKey = "user"
)

func Todo() {
	log.Println(authorizationHeaderKey, authorizationHeaderType, authorizationPayloadKey)
}
