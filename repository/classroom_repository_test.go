package repository

import (
	"context"
	"server/database/mocks/mockdb"
	"server/database/postgresql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateClassroom(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	mockClassroomParam := &postgresql.CreateClassParams{
		ID:          uuid.New(),
		AdminID:     uuid.New(),
		Name:        "Test Classroom",
		Description: "This is a test classroom",
		Section:     "A",
		Room:        "123",
		Subject:     "Math",
		InviteCode:  uuid.New(),
		Visibility:  postgresql.VisibilityPUBLIC,
	}

	mockClassroom := &postgresql.Classroom{
		ID:          mockClassroomParam.ID,
		AdminID:     mockClassroomParam.AdminID,
		Name:        mockClassroomParam.Name,
		Description: mockClassroomParam.Description,
		Section:     mockClassroomParam.Section,
		Room:        mockClassroomParam.Room,
		Subject:     mockClassroomParam.Subject,
		InviteCode:  mockClassroomParam.InviteCode,
		Visibility:  mockClassroomParam.Visibility,
	}

	t.Run("success", func(t *testing.T) {
		store.EXPECT().CreateClass(gomock.Any(), *mockClassroomParam).Times(1).Return(*mockClassroom, nil)
		cr := NewClassroomRepository(store)
		err := cr.Create(context.TODO(), mockClassroomParam)
		require.NoError(t, err)
	})
}
