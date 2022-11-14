package utils

import (
	"math/rand"
	"strings"
	"time"
)

const characters = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomString(l int) string {
	var sb strings.Builder

	k := len(characters)

	for i := 0; i < l; i++ {
		char := characters[rand.Intn(k)]
		sb.WriteByte(char)
	}
	return sb.String()
}
