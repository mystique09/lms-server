package database

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var user_id uuid.UUID

func TestCreateUser(t *testing.T) {
	args := CreateUserParams{
		ID:         uuid.New(),
		Username:   "mystique09",
		Password:   "testpassword",
		Email:      "testemail@gmail.com",
		UserRole:   RoleSTUDENT,
		Visibility: VisibilityPUBLIC,
	}

	user, err := testQueries.CreateUser(context.Background(), args)

	if assert.NoError(t, err) {
		assert.Equal(t, args.Username, user.Username)
		assert.Equal(t, args.Password, user.Password)
		assert.Equal(t, args.Email, user.Email)
		assert.Equal(t, args.UserRole, user.UserRole)
		assert.NotZero(t, user.CreatedAt)
		assert.NotZero(t, user.UpdatedAt)
		user_id = user.ID
	}
	//testQueries.DeleteUser(context.Background(), user.ID)
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
	deleted, err := testQueries.DeleteUser(context.Background(), user_id)

	if assert.NoError(t, err) {
		assert.Equal(t, user_id, deleted.ID)
	}
}
