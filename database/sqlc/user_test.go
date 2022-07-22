package database

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var user_id uuid.UUID

func TestCreateUser(t *testing.T) {
	username := "mystique09"
	password := "testpassword"
	email := "testemail@gmail.com"
	user, err := precreateUser(username, password, email)

	if assert.NoError(t, err) {
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, RoleSTUDENT, user.UserRole)
		assert.NotZero(t, user.CreatedAt)
		assert.NotZero(t, user.UpdatedAt)
		user_id = user.ID
	}
}

func TestGetUser(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), user_id)

	if assert.NoError(t, err) {
		assert.Equal(t, user_id, user.ID)
		assert.Equal(t, "mystique09", user.Username)
		assert.Equal(t, "testemail@gmail.com", user.Email)
		assert.Equal(t, RoleSTUDENT, user.UserRole)
		assert.Equal(t, VisibilityPUBLIC, user.Visibility)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	args := UpdateUserPasswordParams{
		ID:       user_id,
		Password: "new_testpassword",
	}

	updated, err := testQueries.UpdateUserPassword(context.Background(), args)

	if assert.NoError(t, err) {
		assert.Equal(t, "new_testpassword", updated.Password)
		assert.NotEqual(t, "testpassword", updated.Password)
	}
}

func TestUpdateUserEmail(t *testing.T) {
	args := UpdateUserEmailParams{
		ID:    user_id,
		Email: "new_testemail@gmail.com",
	}

	updated, err := testQueries.UpdateUserEmail(context.Background(), args)

	if assert.NoError(t, err) {
		assert.Equal(t, "new_testemail@gmail.com", updated.Email)
		assert.NotEqual(t, "testemail@gmail.com", updated.Email)
	}
}

func TestUpdateUserName(t *testing.T) {
	args := UpdateUsernameParams{
		ID:       user_id,
		Username: "new_mystique09",
	}

	updated, err := testQueries.UpdateUsername(context.Background(), args)

	if assert.NoError(t, err) {
		assert.Equal(t, "new_mystique09", updated.Username)
		assert.NotEqual(t, "mystique09", updated.Username)
	}
}

func TestDeleteUser(t *testing.T) {
	postDeleteUser(user_id)

	user, err := testQueries.GetUser(context.Background(), user_id)

	if assert.Error(t, err) {
		assert.Equal(t, uuid.Nil, user.ID)
	}
}
