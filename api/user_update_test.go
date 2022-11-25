package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	mockdb "server/database/mock"
	database "server/database/sqlc"
	"server/utils"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestUpdateUserApi(t *testing.T) {
	user, password := randomUser(t)

	randomUsername := utils.RandomString(12)
	randomEmail := utils.RandomEmail()

	testCases := []struct {
		name          string
		field         string
		body          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(rec *httptest.ResponseRecorder)
	}{
		{
			name:  "Update username",
			field: "username",
			body:  fmt.Sprintf(`{"username": "%v", "email": "%v", "password": "%v"}`, randomUsername, user.Email, password),
			buildStubs: func(store *mockdb.MockStore) {
				update_username_params := database.UpdateUsernameParams{
					ID:       user.ID,
					Username: randomUsername,
				}
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user, nil)
				updateUserMock := store.EXPECT().UpdateUsername(gomock.Any(), gomock.Eq(update_username_params)).Times(1).Return(user.ID, nil)

				gomock.InOrder(getUserMock, updateUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				var res Response[uuid.UUID]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)
				t.Log(res)
				require.Equal(t, user.ID, res.Data)
				require.Equal(t, 200, rec.Code)
			},
		},
		{
			name:  "Update email",
			field: "email",
			body:  fmt.Sprintf(`{"username": "%v", "email": "%v", "password": "%v"}`, user.Username, randomEmail, password),
			buildStubs: func(store *mockdb.MockStore) {
				update_email_params := database.UpdateUserEmailParams{
					ID:    user.ID,
					Email: randomEmail,
				}
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user, nil)
				updateUserMock := store.EXPECT().UpdateUserEmail(gomock.Any(), gomock.Eq(update_email_params)).Times(1).Return(user.ID, nil)

				gomock.InOrder(getUserMock, updateUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				var res Response[uuid.UUID]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Equal(t, user.ID, res.Data)
				require.Equal(t, 200, rec.Code)
			},
		},
		{
			name:  "Missing email field",
			field: "email",
			body:  fmt.Sprintf(`{"username": "%v", "email": "%v", "password": "%v"}`, randomUsername, "", password),
			buildStubs: func(store *mockdb.MockStore) {
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(0)
				updateUserMock := store.EXPECT().UpdateUserEmail(gomock.Any(), gomock.Any()).Times(0)

				gomock.InOrder(getUserMock, updateUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				var res Response[any]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Contains(t, res.Error, "Email", "required")
				require.Nil(t, res.Data)
				require.Equal(t, 400, rec.Code)
			},
		},
		{
			name:  "Missing username field",
			field: "username",
			body:  fmt.Sprintf(`{"email": "%v", "password": "%v"}`, user.Email, password),
			buildStubs: func(store *mockdb.MockStore) {
				getUserMock := store.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(0)
				updateUserMock := store.EXPECT().UpdateUsername(gomock.Any(), gomock.Any()).Times(0)

				gomock.InOrder(getUserMock, updateUserMock)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				var res Response[any]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Contains(t, res.Error, "required", "Username")
				require.Nil(t, res.Data)
				require.Equal(t, 400, rec.Code)
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
			data := strings.NewReader(tc.body)

			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/users/%v?field=%v", user.ID, tc.field), data)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			paseto_token, tokenPayload, err := server.tokenMaker.CreateToken(user.Username, cfg.AccessTokenDuration)
			require.NoError(t, err)
			require.NotNil(t, tokenPayload)

			req.Header.Set("authorization", fmt.Sprintf("Bearer %v", paseto_token))

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(rec)
		})
	}
}
