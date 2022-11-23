package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	database "server/database/sqlc"
	"server/utils"
	"strings"
	"testing"

	mockdb "server/database/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      database.CreateUserParams
	password string
}

func (e *eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(database.CreateUserParams)
	if !ok {
		return false
	}

	err := utils.MatchPassword([]byte(e.password), []byte(arg.Password))
	if err != nil {
		return false
	}

	e.arg.ID = arg.ID
	e.arg.Password = arg.Password

	return reflect.DeepEqual(e.arg, arg)
}

func (e *eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and id %v", e.arg, e.password)
}

func EqCreateUserParams(arg *database.CreateUserParams, password string) gomock.Matcher {
	return &eqCreateUserParamsMatcher{*arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	// TODO!
	user, password := randomUser(t)

	testCases := []struct {
		name          string
		body          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(rec *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: fmt.Sprintf(`{"username":"%v", "password": "%v", "email": "%v"}`, user.Username, password, user.Email),
			buildStubs: func(store *mockdb.MockStore) {
				arg := database.CreateUserParams{
					ID:         uuid.New(),
					Username:   user.Username,
					Email:      user.Email,
					Password:   user.Password,
					UserRole:   database.RoleSTUDENT,
					Visibility: database.VisibilityPUBLIC,
				}

				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(&arg, password)).
					Times(1).Return(arg.Username, nil)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
			},
		},
		{
			name: "Missing username field",
			body: fmt.Sprintf(`{"password": "%v", "email": "%v"}`, password, user.Email),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "Missing password field",
			body: fmt.Sprintf(`{"username": "%v", "email": "%v"}`, user.Username, user.Email),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "Missing email field",
			body: fmt.Sprintf(`{"password": "%v", "username": "%v"}`, password, user.Username),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "Invalid email",
			body: fmt.Sprintf(`{"password": "%v", "username": "%v", "email":"invalidemail"}`, password, user.Username),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			name: "Invalid username length",
			body: fmt.Sprintf(`{"password": "%v", "username": "%v", "email":"%v"}`, password, utils.RandomString(4), user.Email),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, rec.Code)
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

			url := "/api/v1/signup"

			req := httptest.NewRequest(http.MethodPost, url, data)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			require.NoError(t, err)

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(rec)
		})
	}
}
