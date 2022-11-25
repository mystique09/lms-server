package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	const STR_LEN = 12
	randomString := RandomString(STR_LEN)

	require.NotNil(t, randomString)
	require.NotEmpty(t, randomString)
	require.Equal(t, STR_LEN, len(randomString))
}

func TestRandomEmail(t *testing.T) {
	randomEmail := RandomEmail()

	require.NotNil(t, randomEmail)
	require.NotEmpty(t, randomEmail)
	require.Contains(t, randomEmail, "@")
}
