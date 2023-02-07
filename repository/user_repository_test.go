package repository

import (
	"context"
	"database/sql"
	"server/database/mocks/mockdb"
	"server/database/postgresql"
	"server/domain"
	"server/internal/stringutil"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	mockUser := &postgresql.CreateUserParams{
		ID:         uuid.New(),
		Username:   stringutil.RandomString(12),
		Email:      stringutil.RandomEmail(),
		Password:   stringutil.RandomString(14),
		UserRole:   postgresql.RoleSTUDENT,
		Visibility: postgresql.VisibilityPUBLIC,
	}

	mockEmptyUser := &postgresql.CreateUserParams{}

	t.Run("success", func(t *testing.T) {
		store.EXPECT().CreateUser(gomock.Any(), gomock.Eq(*mockUser)).Times(1).Return(mockUser.ID.String(), nil)
		ur := NewUserRepository(store)
		err := ur.Create(context.TODO(), mockUser)
		require.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		store.EXPECT().CreateUser(gomock.Any(), gomock.Eq(*mockEmptyUser)).Times(1).Return(mockUser.ID.String(), errors.New("EMPTY USER DATA"))
		ur := NewUserRepository(store)
		err := ur.Create(context.TODO(), mockEmptyUser)
		require.Error(t, err)
	})
}

func TestFetchUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	mockUsers := make([]domain.User, 20)

	t.Run("success", func(t *testing.T) {
		store.EXPECT().GetUsers(gomock.Any(), gomock.Eq(int32(10))).Times(1).Return(mockUsers, nil)
		ur := NewUserRepository(store)
		users, err := ur.Fetch(context.TODO(), 1)
		require.NotEmpty(t, users)
		require.NoError(t, err)
	})
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	mockUser := &domain.User{
		ID: uuid.New(),
	}

	mockEmptyUser := &domain.User{}

	t.Run("success", func(t *testing.T) {
		store.EXPECT().GetUser(gomock.Any(), gomock.Eq(mockUser.ID)).Times(1).Return(*mockUser, nil)
		ur := NewUserRepository(store)
		user, err := ur.GetByID(context.TODO(), mockUser.ID)
		require.NotEmpty(t, user)
		require.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		store.EXPECT().GetUser(gomock.Any(), gomock.Eq(mockEmptyUser.ID)).Times(1).Return(*mockEmptyUser, sql.ErrNoRows)
		ur := NewUserRepository(store)
		user, err := ur.GetByID(context.TODO(), mockEmptyUser.ID)
		require.Empty(t, user)
		require.Error(t, err)
	})
}

func TestGetUserByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	mockUser := &domain.User{
		Username: stringutil.RandomString(12),
	}

	mockEmptyUser := &domain.User{}

	t.Run("success", func(t *testing.T) {
		store.EXPECT().GetUserByUsername(gomock.Any(), gomock.Eq(mockUser.Username)).Times(1).Return(*mockUser, nil)
		ur := NewUserRepository(store)
		user, err := ur.GetByUsername(context.TODO(), mockUser.Username)
		require.NotEmpty(t, user)
		require.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		store.EXPECT().GetUserByUsername(gomock.Any(), gomock.Eq(mockEmptyUser.Username)).Times(1).Return(*mockEmptyUser, sql.ErrNoRows)
		ur := NewUserRepository(store)
		user, err := ur.GetByUsername(context.TODO(), mockEmptyUser.Username)
		require.Empty(t, user)
		require.Error(t, err)
	})
}

func TestUpdateUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	mockUser := &postgresql.UpdateUsernameParams{
		Username: stringutil.RandomString(12),
		ID:       uuid.New(),
	}

	mockEmptyUser := &postgresql.UpdateUsernameParams{}

	t.Run("success", func(t *testing.T) {
		store.EXPECT().UpdateUsername(gomock.Any(), gomock.Eq(*mockUser)).Times(1).Return(mockUser.ID, nil)
		ur := NewUserRepository(store)
		err := ur.UpdateUsername(context.TODO(), mockUser)
		require.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		store.EXPECT().UpdateUsername(gomock.Any(), gomock.Eq(*mockEmptyUser)).Times(1).Return(mockEmptyUser.ID, sql.ErrNoRows)
		ur := NewUserRepository(store)
		err := ur.UpdateUsername(context.TODO(), mockEmptyUser)
		require.Error(t, err)
	})

	t.Run("internal error", func(t *testing.T) {
		store.EXPECT().UpdateUsername(gomock.Any(), gomock.Eq(*mockEmptyUser)).Times(1).Return(mockEmptyUser.ID, sql.ErrConnDone)
		ur := NewUserRepository(store)
		err := ur.UpdateUsername(context.TODO(), mockEmptyUser)
		require.Error(t, err)
	})
}

func TestUpdateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	mockUser := &postgresql.UpdateUserEmailParams{
		Email: stringutil.RandomEmail(),
		ID:    uuid.New(),
	}

	mockEmptyUser := &postgresql.UpdateUserEmailParams{}

	t.Run("success", func(t *testing.T) {
		store.EXPECT().UpdateUserEmail(gomock.Any(), gomock.Eq(*mockUser)).Times(1).Return(mockUser.ID, nil)
		ur := NewUserRepository(store)
		err := ur.UpdateEmail(context.TODO(), mockUser)
		require.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		store.EXPECT().UpdateUserEmail(gomock.Any(), gomock.Eq(*mockEmptyUser)).Times(1).Return(mockEmptyUser.ID, sql.ErrNoRows)
		ur := NewUserRepository(store)
		err := ur.UpdateEmail(context.TODO(), mockEmptyUser)
		require.Error(t, err)
	})

	t.Run("internal error", func(t *testing.T) {
		store.EXPECT().UpdateUserEmail(gomock.Any(), gomock.Eq(*mockEmptyUser)).Times(1).Return(mockEmptyUser.ID, sql.ErrConnDone)
		ur := NewUserRepository(store)
		err := ur.UpdateEmail(context.TODO(), mockEmptyUser)
		require.Error(t, err)
	})
}

func TestUpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	mockUser := &postgresql.UpdateUserPasswordParams{
		Password: stringutil.RandomString(14),
		ID:       uuid.New(),
	}

	mockEmptyUser := &postgresql.UpdateUserPasswordParams{}

	t.Run("success", func(t *testing.T) {
		store.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Eq(*mockUser)).Times(1).Return(mockUser.ID, nil)
		ur := NewUserRepository(store)
		err := ur.UpdatePassword(context.TODO(), mockUser)
		require.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		store.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Eq(*mockEmptyUser)).Times(1).Return(mockEmptyUser.ID, sql.ErrNoRows)
		ur := NewUserRepository(store)
		err := ur.UpdatePassword(context.TODO(), mockEmptyUser)
		require.Error(t, err)
	})

	t.Run("internal error", func(t *testing.T) {
		store.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Eq(*mockEmptyUser)).Times(1).Return(mockEmptyUser.ID, sql.ErrConnDone)
		ur := NewUserRepository(store)
		err := ur.UpdatePassword(context.TODO(), mockEmptyUser)
		require.Error(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	mockUser := uuid.New()
	mockNotFoundUser := uuid.New()

	t.Run("success", func(t *testing.T) {
		store.EXPECT().DeleteUser(gomock.Any(), gomock.Eq(mockUser)).Times(1).Return(mockUser, nil)
		ur := NewUserRepository(store)
		err := ur.Delete(context.TODO(), mockUser)
		require.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		store.EXPECT().DeleteUser(gomock.Any(), gomock.Eq(mockNotFoundUser)).Times(1).Return(uuid.Nil, sql.ErrNoRows)
		ur := NewUserRepository(store)
		err := ur.Delete(context.TODO(), mockNotFoundUser)
		require.Error(t, err)
	})

	t.Run("internal error", func(t *testing.T) {
		store.EXPECT().DeleteUser(gomock.Any(), gomock.Eq(mockNotFoundUser)).Times(1).Return(uuid.Nil, sql.ErrConnDone)
		ur := NewUserRepository(store)
		err := ur.Delete(context.TODO(), mockNotFoundUser)
		require.Error(t, err)
	})
}
