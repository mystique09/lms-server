package tokenutil

import (
	"server/internal/stringutil"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPasteMaker(t *testing.T) {
	maker, err := NewPasetoMaker(stringutil.RandomString(32))
	require.NoError(t, err)

	username := stringutil.RandomString(8)
	uid := uuid.New()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, tokenPayload, err := maker.CreateToken(username, uid, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotNil(t, tokenPayload)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}

func TestExpirePasteToken(t *testing.T) {
	maker, err := NewPasetoMaker(stringutil.RandomString(32))
	require.NoError(t, err)

	username := stringutil.RandomString(8)
	uid := uuid.New()
	duration := -time.Minute

	token, tokenPayload, err := maker.CreateToken(username, uid, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotNil(t, tokenPayload)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
