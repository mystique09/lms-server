package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	database "server/database/sqlc"
	"server/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var cfg utils.Config

func TestMain(m *testing.M) {

	conf, err := utils.LoadConfig("../", "app.sample")

	if err != nil {
		log.Fatal(err)
	}

	cfg = conf

	os.Exit(m.Run())
}

func randomUser(t *testing.T) (user database.User, password string) {
	password = utils.RandomString(12)
	hashed_password, err := utils.Encrypt(password)
	require.NoError(t, err)

	user = database.User{
		ID:         uuid.New(),
		Username:   utils.RandomString(12),
		Password:   hashed_password,
		Email:      utils.RandomEmail(),
		UserRole:   database.RoleSTUDENT,
		Visibility: database.VisibilityPUBLIC,
	}

	return
}

func requireBodyMatch(t *testing.T, body *bytes.Buffer, user *database.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var getUser Response[*database.User]
	err = json.Unmarshal(data, &getUser)
	require.NoError(t, err)

	user.Password = ""

	require.Equal(t, newResponse(user), getUser)
}
