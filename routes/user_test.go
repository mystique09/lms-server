package routes

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	database "server/database/sqlc"
	"server/utils"
	"testing"

	mockdb "server/database/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetUserRoute(t *testing.T) {
	user := randomUser()

	testCases := []struct {
		name          string
		id            uuid.UUID
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			id:   user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
				requireBodyMatch(t, rec.Body, user)
			},
		},
		{
			name: "NotFound",
			id:   user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(database.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, rec.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)

			rec := httptest.NewRecorder()
			url := fmt.Sprintf("/api/v1/users/%v", tc.id)

			req, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(t, rec)
		})
	}
}

func randomUser() database.User {
	return database.User{
		ID:         uuid.New(),
		Username:   utils.RandomString(8),
		Password:   utils.RandomString(8),
		Email:      utils.RandomString(16),
		UserRole:   database.RoleSTUDENT,
		Visibility: database.VisibilityPUBLIC,
	}
}

func requireBodyMatch(t *testing.T, body *bytes.Buffer, user database.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var getUser database.User
	err = json.Unmarshal(data, &getUser)
	require.NoError(t, err)
	user.Password = ""

	require.Equal(t, user, getUser)
}
