package tokenutil

import (
	"time"
	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(username string, uid uuid.UUID, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
