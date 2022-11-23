package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	mockdb "server/database/mock"
	database "server/database/sqlc"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetUserAPI(t *testing.T) {
	user, _ := randomUser(t)

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
				requireBodyMatch(t, rec.Body, &user)
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
		{
			name: "Inter server error",
			id:   user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).Return(database.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rec.Code)
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

			rec := httptest.NewRecorder()
			url := fmt.Sprintf("/api/v1/users/%v", tc.id)

			req, err := http.NewRequest(http.MethodGet, url, http.NoBody)

			require.NoError(t, err)

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(t, rec)
		})
	}
}
