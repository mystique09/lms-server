package routes

import (
	"net/http"
	"net/http/httptest"
	"server/utils"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var userRequestJson string = `{"username":"mystique09","password":"testpassword","email":"testemail@gmail.com"}`
var testUserId string

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

func TestDeleteUserById(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
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
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, testUserId, res.Data)
	}
}
