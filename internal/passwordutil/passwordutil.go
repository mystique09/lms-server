package passwordutil

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(plain_text string) (string, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(plain_text), bcrypt.DefaultCost)
	return string(hashed_password), err
}

func MatchPassword(hashed_password, plain_password []byte) error {
	return bcrypt.CompareHashAndPassword(hashed_password, plain_password)
}
