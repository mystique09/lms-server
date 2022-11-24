package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	mockdb "server/database/mock"
	"server/utils"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		name          string
		payload       string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		{
			name:    "Login Success",
			payload: fmt.Sprintf(`{"username":"%v","password":"%v"}`, user.Username, password),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByUsername(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(user, nil)
				store.EXPECT().CreateSession(gomock.Any(), gomock.Any()).Times(1)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var res Response[authSuccessResponse]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.NotEmpty(t, res.Data.AccessToken)
				require.NotEmpty(t, res.Data.User)
				require.Equal(t, 200, rec.Code)
			},
		},
		{
			name:    "Login failed, mismatch password",
			payload: fmt.Sprintf(`{"username":"%v","password":"%v"}`, user.Username, "invalid password"),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByUsername(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(user, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var res Response[any]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Empty(t, res.Data)
				require.NotEmpty(t, res.Error)
				require.Equal(t, LOGIN_FAILED.Error, res.Error)
				require.Equal(t, 401, rec.Code)
			},
		},
		{
			name:    "Missing password field",
			payload: fmt.Sprintf(`{"username":"%v"}`, user.Username),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var res Response[any]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Empty(t, res.Data)
				require.NotEmpty(t, res.Error)
				require.Contains(t, res.Error, "authRequest.Password")
				require.Equal(t, 400, rec.Code)
			},
		},
		{
			name:    "Missing username field",
			payload: fmt.Sprintf(`{"password":"%v"}`, password),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var res Response[any]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Empty(t, res.Data)
				require.NotEmpty(t, res.Error)
				require.Contains(t, res.Error, "authRequest.Username")
				require.Equal(t, 400, rec.Code)
			},
		},
		{
			name:    "Missing payload",
			payload: "",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var res Response[any]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Empty(t, res.Data)
				require.NotEmpty(t, res.Error)
				require.Contains(t, res.Error, "authRequest.Username", "AuthRequest.Password")
				require.Equal(t, 400, rec.Code)
			},
		},
		{
			name:    "Username length error",
			payload: fmt.Sprintf(`{"usernam":"%v","password":"%v"}`, utils.RandomString(5), password),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var res Response[any]

				body, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				err = json.Unmarshal(body, &res)
				require.NoError(t, err)

				require.Empty(t, res.Data)
				require.NotEmpty(t, res.Error)
				require.Contains(t, res.Error, "authRequest.Username", "required")
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

			url := "/api/v1/login"
			req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(tc.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			require.NoError(t, err)

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(t, rec)
		})
	}
}
