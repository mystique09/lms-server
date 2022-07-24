package routes

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"server/utils"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var userRequestJson string = `{"username":"mystique09","password":"testpassword","email":"testemail@gmail.com"}`
var testUserId string
var token string
var refreshToken string

func TestCreateUserRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(userRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testServer.createUser(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		user := res.Data.(map[string]interface{})
		assert.Equal(t, "STUDENT", user["user_role"])
		assert.Equal(t, "PUBLIC", user["visibility"])
		assert.Equal(t, "mystique09", user["username"])
		assert.Empty(t, user["password"])
		assert.Equal(t, "testemail@gmail.com", user["email"])
		testUserId = user["id"].(string)
	}
}

func TestLoginMystique09(t *testing.T) {
	var payload string = `{"username":"mystique09","password":"testpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.Echo().Use(JwtAuthMiddleware(server.Cfg))

	if assert.NoError(t, testServer.loginHandler(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		success_resp := res.Data.(map[string]interface{})
		assert.Equal(t, success_resp["message"], "Logged in.")
		assert.NotEmpty(t, success_resp["access_token"])
		assert.NotEmpty(t, success_resp["refresh_token"])
		token = success_resp["access_token"].(string)
		refreshToken = success_resp["refresh_token"].(string)
	}
}

func TestGetUsersRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.Echo().Use(JwtAuthMiddleware(server.Cfg))

	if assert.NoError(t, testServer.getUsers(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, res.Data)
	}
}

func TestGetUserById(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/v1/users/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(testUserId)
	ctx.Echo().Use(JwtAuthMiddleware(server.Cfg))

	if assert.NoError(t, testServer.getUser(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		user := res.Data.(map[string]interface{})
		assert.Equal(t, testUserId, user["id"].(string))
	}
}

func TestUpdateUsernameById(t *testing.T) {
	var payload string = `{"username":"updated_username"}`

	q := make(url.Values)
	q.Set("field", "username")

	req := httptest.NewRequest(http.MethodPut, "/?"+q.Encode(), strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/v1/users/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(testUserId)
	ctx.Echo().Use(JwtAuthMiddleware(server.Cfg))
	parsed_token, err := jwt.Parse(token, keyFunc(testServer.Cfg.JWT_SECRET_KEY))
	if err != nil {
		t.Fail()
	}

	ctx.Set("user", parsed_token)
	if assert.NoError(t, testServer.updateUser(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		updated_user := res.Data.(map[string]interface{})
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEqual(t, "mystique09", updated_user["username"])
	}
}

func TestRefreshTokenAfterUpdateUsername(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/refresh", nil)
	//req.Header.Set(echo.HeaderAuthorization, "Bearer "+refreshToken)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	parsed_token, err := jwt.Parse(refreshToken, keyFunc(testServer.Cfg.JWT_REFRESH_SECRET_KEY))
	if err != nil {
		t.Fail()
	}

	ctx.Set("refresh", parsed_token)
	ctx.Echo().Use(RefreshTokenAuthMiddleware(testServer.Cfg))

	if assert.NoError(t, testServer.refreshToken(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		assert.Equal(t, 200, rec.Code)
		data := res.Data.(map[string]interface{})
		token = data["access_token"].(string)
		assert.NotEmpty(t, token)
	}
}

func TestDeleteUserById(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/v1/users/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(testUserId)
	ctx.Echo().Use(JwtAuthMiddleware(server.Cfg))
	parsed_token, err := jwt.Parse(token, keyFunc(testServer.Cfg.JWT_SECRET_KEY))
	if err != nil {
		t.Fail()
	}

	ctx.Set("user", parsed_token)
	if assert.NoError(t, testServer.deleteUser(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
