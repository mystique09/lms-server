package routes

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func TestGetAllUsersApi(t *testing.T) {
	var users []database.User

	for i := 0; i < 20; i++ {
		user, _ := randomUser(t)
		user.Password = ""
		users = append(users, user)
	}

	testCases := []struct {
		name          string
		offset        string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			offset: "1",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUsers(gomock.Any(), gomock.Eq(int32(10))).
					Times(1).
					Return(users, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, 200, rec.Code)

				data, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				var getUsers Response[[]database.User]
				err = json.Unmarshal(data, &getUsers)
				require.NoError(t, err)
				require.Equal(t, newResponse(users), getUsers)
			},
		},
		{
			name:   "OK: Offset is 0",
			offset: "0",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUsers(gomock.Any(), gomock.Eq(int32(0))).
					Times(1).
					Return(users, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, 200, rec.Code)

				data, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				var getUsers Response[[]database.User]
				err = json.Unmarshal(data, &getUsers)
				require.NoError(t, err)
				require.Equal(t, newResponse(users), getUsers)
			},
		},
		{
			name:   "OK: Offset is empty",
			offset: "",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUsers(gomock.Any(), gomock.Eq(int32(0))).
					Times(1).
					Return(users, nil)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, 200, rec.Code)

				data, err := io.ReadAll(rec.Body)
				require.NoError(t, err)

				var getUsers Response[[]database.User]
				err = json.Unmarshal(data, &getUsers)
				require.NoError(t, err)
				require.Equal(t, newResponse(users), getUsers)
			},
		},
		{
			name:   "Invalid offset",
			offset: "i am an invalid offset",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUsers(gomock.Any(), gomock.Eq(0)).
					Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, 400, rec.Code)
			},
		},
		{
			name:   "Negative offset",
			offset: "-1",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetUsers(gomock.Any(), gomock.Eq(0)).Times(0)
			},
			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
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
			url := fmt.Sprintf("/api/v1/users?offset=%v", tc.offset)

			req, err := http.NewRequest(http.MethodGet, url, http.NoBody)

			require.NoError(t, err)

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(t, rec)
		})
	}
}

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

			req, err := http.NewRequest(http.MethodPost, url, data)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			require.NoError(t, err)

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(rec)
		})
	}
}

func TestUpdateUserApi(t *testing.T) {
	// TODO!
}

func TestDeleteUserApi(t *testing.T) {
	// TODO!
}

func randomUser(t *testing.T) (user database.User, password string) {
	password = utils.RandomString(12)
	hashed_password, err := utils.Encrypt(password)
	require.NoError(t, err)

	user = database.User{
		ID:         uuid.New(),
		Username:   utils.RandomString(12),
		Password:   hashed_password,
		Email:      utils.RandomEmail(),
		UserRole:   database.RoleSTUDENT,
		Visibility: database.VisibilityPUBLIC,
	}

	return
}

func requireBodyMatch(t *testing.T, body *bytes.Buffer, user *database.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var getUser Response[*database.User]
	err = json.Unmarshal(data, &getUser)
	log.Println(getUser)
	require.NoError(t, err)

	user.Password = ""

	require.Equal(t, newResponse(user), getUser)
}
