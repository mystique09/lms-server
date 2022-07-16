package database

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	args := CreateUserParams{
		ID:       uuid.New(),
		Username: "mystique09",
		Password: "testpassword",
		Email:    "testemail@gmail.com",
		UserRole: RoleSTUDENT,
	}

	user, err := testQueries.CreateUser(context.Background(), args)

	if assert.NoError(t, err) {
		assert.Equal(t, args.Username, user.Username)
		assert.Equal(t, args.Password, user.Password)
		assert.Equal(t, args.Email, user.Email)
		assert.Equal(t, args.UserRole, user.UserRole)
		assert.NotZero(t, user.CreatedAt)
		assert.NotZero(t, user.UpdatedAt)
	}
	testQueries.DeleteUser(context.Background(), user.ID)
}
