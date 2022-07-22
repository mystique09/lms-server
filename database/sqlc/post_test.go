package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllPostFromAClass(t *testing.T) {
	user, err := precreateUser("mystique09", "testpassword", "testemail@gmail.com")
	if err != nil {
		t.Fail()
	}
	class, err := preCreateClassroom(user.ID, "TestClassroomName")
	if err != nil {
		t.Fail()
	}
	param := ListAllPostsFromClassParams{
		ClassID: class.ID,
		Offset:  0,
	}

	posts, err := testQueries.ListAllPostsFromClass(context.Background(), param)

	if assert.NoError(t, err) {
		assert.Empty(t, posts)
	}
	postDeleteUser(user.ID)
	postDeleteClassroom(class.ID)
}
