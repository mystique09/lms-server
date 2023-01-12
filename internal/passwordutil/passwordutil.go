package passwordutil

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(plain_text string) (string, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(plain_text), bcrypt.DefaultCost)

	if err != nil {
		return plain_text, err
	}
	return string(hashed_password), nil
}

func MatchPassword(p, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, p)
}
