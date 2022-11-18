package token

import "time"

type Maker interface {
	Create(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
