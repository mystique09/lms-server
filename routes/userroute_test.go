package routes

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"server/utils"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var userRequestJson string = `{"username":"mystique09","password":"testpassword","email":"testemail@gmail.com"}`
var testUserId string
var token string

func TestGetUsersRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testRoute.getUsers(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, res.Data)
	}

}

func TestCreateUserRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(userRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testRoute.createUser(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		user := res.Data.(map[string]interface{})
		t.Logf("User %v", user)
		assert.Equal(t, "STUDENT", user["user_role"])
		assert.Empty(t, user["password"])
		assert.Equal(t, "testemail@gmail.com", user["email"])
		testUserId = user["id"].(string)
		t.Log(testUserId)
	}
}

func TestLoginMystique09(t *testing.T) {
	var payload string = `{"username":"mystique09","password":"testpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testRoute.loginRoute(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		success_resp := res.Data.(map[string]interface{})
		assert.Equal(t, success_resp["message"], "Logged in.")
		assert.NotEmpty(t, success_resp["access_token"])
		assert.NotEmpty(t, success_resp["refresh_token"])
		token = success_resp["access_token"].(string)
	}
}

func TestGetUserById(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/v1/users/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(testUserId)

	if assert.NoError(t, testRoute.getUser(ctx)) {
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
	req.Header.Set("authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/v1/users/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(testUserId)

	if assert.NoError(t, testRoute.updateUser(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		t.Log(res)
		updated_user := res.Data.(map[string]interface{})
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEqual(t, "mystique09", updated_user["username"])
	}
}

func TestDeleteUserById(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set("authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/v1/users/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(testUserId)

	if assert.NoError(t, testRoute.deleteUser(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		deleted_user := res.Data.(map[string]interface{})
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, testUserId, deleted_user["id"])
	}
}
