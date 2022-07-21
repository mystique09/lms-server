package database

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"github.com/stretchr/testify/assert"
)

var class_id uuid.UUID
var class_invcode uuid.UUID

func TestCreateOneNewClassroom(t *testing.T) {
	uid := uuid.New()
	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		ID:         uid,
		Username:   "mystique08",
		Password:   "testpassword",
		Email:      "testemail233@gmail.com",
		UserRole:   RoleSTUDENT,
		Visibility: VisibilityPUBLIC,
	})

	if err != nil {
		t.Fail()
	}

	args := CreateClassParams{
		ID:          uuid.New(),
		AdminID:     user.ID,
		Description: random.New().String(20, charsets),
		Room:        random.New().String(20, charsets),
		Section:     random.New().String(20, charsets),
		Subject:     random.New().String(20, charsets),
		InviteCode:  uuid.New(),
		Visibility:  VisibilityPUBLIC,
	}

	classroom, err := testQueries.CreateClass(context.Background(), args)

	if assert.NoError(t, err) {
		assert.Equal(t, uid, classroom.AdminID)
		assert.Len(t, classroom.Description, 20)
		assert.Len(t, classroom.Room, 20)
		assert.Len(t, classroom.Section, 20)
		assert.Len(t, classroom.Subject, 20)
		assert.NotNil(t, classroom.InviteCode)
		assert.Equal(t, classroom.Visibility, VisibilityPUBLIC)
	}

	user_id = uid
	class_id = classroom.ID
	class_invcode = classroom.InviteCode
}

func TestCreateNewMoreClassrooms(t *testing.T) {
	args := []CreateClassParams{
		{
			ID:          uuid.New(),
			AdminID:     user_id,
			Description: random.New().String(20, charsets),
			Room:        random.New().String(20, charsets),
			Section:     random.New().String(20, charsets),
			Subject:     random.New().String(20, charsets),
			InviteCode:  uuid.New(),
			Visibility:  VisibilityPUBLIC,
		},
		{
			ID:          uuid.New(),
			AdminID:     user_id,
			Description: random.New().String(20, charsets),
			Room:        random.New().String(20, charsets),
			Section:     random.New().String(20, charsets),
			Subject:     random.New().String(20, charsets),
			InviteCode:  uuid.New(),
			Visibility:  VisibilityPUBLIC,
		},
		{
			ID:          uuid.New(),
			AdminID:     user_id,
			Description: random.New().String(20, charsets),
			Room:        random.New().String(20, charsets),
			Section:     random.New().String(20, charsets),
			Subject:     random.New().String(20, charsets),
			InviteCode:  uuid.New(),
			Visibility:  VisibilityPUBLIC,
		},
	}

	for i, c := range args {
		class, err := testQueries.CreateClass(context.Background(), c)
		t.Logf("Test case %v for id %v", i+1, c.ID)
		if assert.NoError(t, err) {
			assert.Equal(t, class.AdminID, c.AdminID)
			assert.Len(t, class.Description, 20)
			assert.Len(t, class.Room, 20)
			assert.Len(t, class.Section, 20)
			assert.Len(t, class.Subject, 20)
			assert.NotNil(t, class.InviteCode)
			assert.Equal(t, class.Visibility, VisibilityPUBLIC)

		}
	}
	//testQueries.DeleteUser(context.Background(), user_id)
}

func TestUpdateOneClassroom(t *testing.T) {
	args := UpdateClassParams{
		ID:         class_id,
		Name:       random.New().String(15, charsets),
		Subject:    random.New().String(12, charsets),
		Section:    random.New().String(8, charsets),
		Room:       random.New().String(5, charsets),
		InviteCode: uuid.New(),
	}

	updated, err := testQueries.UpdateClass(context.Background(), args)

	if assert.NoError(t, err) {
		assert.NotEqual(t, updated.InviteCode, class_invcode)
		assert.Len(t, updated.Name, 15)
		assert.Len(t, updated.Subject, 12)
		assert.Len(t, updated.Section, 8)
		assert.Len(t, updated.Room, 5)
	}
	//testQueries.DeleteUser(context.Background(), user_id)
}

func TestDeleteClassroom(t *testing.T) {
	deleted, err := testQueries.DeleteClass(context.Background(), class_id)

	if assert.NoError(t, err) {
		assert.Equal(t, deleted.ID, class_id)
	}

	class, err := testQueries.GetClass(context.Background(), class_id)

	if assert.Error(t, err) {
		assert.Equal(t, class.ID, uuid.Nil)
	}

	testQueries.DeleteUser(context.Background(), user_id)
}
