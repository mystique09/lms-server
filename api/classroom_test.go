package api

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	mockdb "server/database/mock"
	database "server/database/sqlc"
	"server/utils"
	"testing"

	"encoding/json"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func randomClassroom(adminId uuid.UUID) database.Classroom {
	return database.Classroom{
		ID:          uuid.New(),
		AdminID:     adminId,
		Name:        utils.RandomString(12),
		Description: utils.RandomString(24),
		Section:     utils.RandomString(8),
		Subject:     utils.RandomString(6),
		Room:        utils.RandomString(8),
		InviteCode:  uuid.New(),
		Visibility:  database.VisibilityPUBLIC,
	}
}

func TestGetClassrooms(t *testing.T) {
	user, _ := randomUser(t)
	randomUuid := uuid.New()
	var dummyClassrooms []database.Classroom = make([]database.Classroom, 20)

	for i := range dummyClassrooms {
		dummyClassrooms[i] = randomClassroom(user.ID)
	}

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(rec *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  fmt.Sprintf("/api/v1/users/%s/classrooms?page=1", user.ID.String()),
			buildStubs: func(store *mockdb.MockStore) {
				param := database.GetAllClassFromUserParams{
					AdminID: user.ID,
					Offset:  10,
				}
				store.EXPECT().GetAllClassFromUser(gomock.Any(), gomock.Eq(param)).Times(1).Return(dummyClassrooms, nil)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, 200, rec.Code)

				var res Response[[]database.Classroom]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.NotEmpty(t, res.Data)
				require.NotNil(t, res.Data)
				require.Greater(t, len(res.Data), 1)
			},
		},
		{
			name: "Invalid UUID",
			url:  "/api/v1/users/invaliduuid/classrooms",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAllClassFromUser(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, 400, rec.Code)

				var res Response[string]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Contains(t, res.Error, "invalid")
			},
		},
		{
			name: "Invalid page type",
			url:  fmt.Sprintf("/api/v1/users/%s/classrooms?page=invalid", user.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAllClassFromUser(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, 400, rec.Code)

				var res Response[string]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Empty(t, res.Data)
				require.NotNil(t, res.Error)
			},
		},
		{
			name: "Empty page query",
			url:  fmt.Sprintf("/api/v1/users/%s/classrooms", user.ID),
			buildStubs: func(store *mockdb.MockStore) {
				param := database.GetAllClassFromUserParams{
					AdminID: user.ID,
					Offset:  0,
				}
				store.EXPECT().GetAllClassFromUser(gomock.Any(), gomock.Eq(param)).Times(1).Return(dummyClassrooms, nil)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, 200, rec.Code)

				var res Response[[]database.Classroom]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.NotEmpty(t, res.Data)
				require.NotNil(t, res.Data)
				require.Greater(t, len(res.Data), 1)
			},
		},
		{
			name: "Admin doesn't exist",
			url:  fmt.Sprintf("/api/v1/users/%s/classrooms?page=1", randomUuid),
			buildStubs: func(store *mockdb.MockStore) {
				param := database.GetAllClassFromUserParams{
					AdminID: randomUuid,
					Offset:  10,
				}
				store.EXPECT().GetAllClassFromUser(gomock.Any(), gomock.Eq(param)).Times(1).Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, 400, rec.Code)
			},
		},
		// TODO: Add a test for negative page query
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

			req := httptest.NewRequest(http.MethodGet, tc.url, nil)

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(rec)
		})
	}
}
