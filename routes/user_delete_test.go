package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	mockdb "server/database/mock"
	database "server/database/sqlc"
	"server/utils"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestDeleteUserApi(t *testing.T) {
	user, _ := randomUser(t)

	testCases := []struct {
		name          string
		body          string
		tokenPayload  string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(rec *httptest.ResponseRecorder)
	}{
		{
			name:         "Success",
			tokenPayload: user.Username,
			body:         user.ID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user, nil)
				deleteUserMock := store.EXPECT().DeleteUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user.ID, nil)

				gomock.InOrder(getUserMock, deleteUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				var res Response[uuid.UUID]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Equal(t, 200, rec.Code)
				require.Equal(t, newResponse(user.ID), res)
			},
		},
		{
			name:         "Unauthorized",
			tokenPayload: utils.RandomString(12),
			body:         user.ID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user, nil)
				deleteUserMock := store.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Times(0)

				gomock.InOrder(getUserMock, deleteUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				var res Response[string]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Equal(t, 401, rec.Code)
				require.Equal(t, UNAUTHORIZED, res)
			},
		},
		{
			name:         "User ID not found",
			tokenPayload: user.Username,
			body:         user.ID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(database.User{}, sql.ErrNoRows)
				deleteUserMock := store.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Times(0)

				gomock.InOrder(getUserMock, deleteUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				var res Response[string]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Equal(t, 400, rec.Code)
				require.Equal(t, USER_NOTFOUND, res)
			},
		},
		{
			name:         "Internal server error",
			tokenPayload: user.Username,
			body:         user.ID.String(),
			buildStubs: func(store *mockdb.MockStore) {
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(database.User{}, sql.ErrConnDone)
				deleteUserMock := store.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Times(0)

				gomock.InOrder(getUserMock, deleteUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				var res Response[string]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Equal(t, http.StatusInternalServerError, rec.Code)
				require.Equal(t, newError("Something went wrong."), res)
			},
		},
		{
			name:         "Invalid UUID",
			tokenPayload: user.Username,
			body:         "invalidUUID",
			buildStubs: func(store *mockdb.MockStore) {
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(0)
				deleteUserMock := store.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Times(0)

				gomock.InOrder(getUserMock, deleteUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				var res Response[string]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Equal(t, http.StatusBadRequest, rec.Code)
				require.Contains(t, res.Error, "invalid", "UUID")
			},
		},
		{
			name:         "Missing UUID",
			tokenPayload: user.Username,
			body:         "",
			buildStubs: func(store *mockdb.MockStore) {
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(0)
				deleteUserMock := store.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Times(0)

				gomock.InOrder(getUserMock, deleteUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
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

			server, err := NewServer(store, &cfg)
			require.NoError(t, err)

			pasetoToken, err := server.tokenMaker.CreateToken(tc.tokenPayload, server.cfg.AccessTokenDuration)
			require.NoError(t, err)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/users/%v", tc.body), nil)
			req.Header.Set("authorization", fmt.Sprintf("Bearer %v", pasetoToken))

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(rec)
		})
	}
}
