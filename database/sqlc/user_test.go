package database

import (
	"context"
	"database/sql"
	"server/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createAccount(t *testing.T) User {
	arg := CreateUserParams{
		ID:         uuid.New(),
		Username:   utils.RandomString(8),
		Password:   utils.RandomString(8),
		Email:      utils.RandomString(16),
		UserRole:   RoleSTUDENT,
		Visibility: VisibilityPUBLIC,
	}

	user, err := testQuesries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.NotZero(t, user.ID)
	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.UserRole, RoleSTUDENT)
	require.Equal(t, arg.Visibility, VisibilityPUBLIC)

	return user
}

func TestCreateAccount(t *testing.T) {
	createAccount(t)
}

func TestGetUser(t *testing.T) {
	new_user := createAccount(t)

	user, err := testQuesries.GetUser(context.Background(), new_user.ID)

	require.NoError(t, err)

	require.NotEmpty(t, user)
	require.Equal(t, new_user.ID, user.ID)
	require.Equal(t, new_user.Username, user.Username)
	require.Equal(t, new_user.Password, user.Password)
	require.Equal(t, new_user.Email, user.Email)
	require.Equal(t, new_user.UserRole, user.UserRole)
	require.Equal(t, new_user.Visibility, user.Visibility)
}

func TestUpdateUsername(t *testing.T) {
	user := createAccount(t)

	arg := UpdateUsernameParams{
		Username: utils.RandomString(8),
		ID:       user.ID,
	}

	updated, err := testQuesries.UpdateUsername(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updated)
	require.Equal(t, arg.Username, updated.Username)
}

func TestDeleteAccount(t *testing.T) {
	new_user := createAccount(t)

	deleted, err := testQuesries.DeleteUser(context.Background(), new_user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, deleted)

	user, err := testQuesries.GetUser(context.Background(), deleted.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createAccount(t)
	}

	users, err := testQuesries.GetUsers(context.Background(), 5)

	require.NoError(t, err)
	require.Equal(t, 10, len(users))

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
