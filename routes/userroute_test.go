package routes

import (
	"net/http"
	"net/http/httptest"
	"server/utils"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var userRequestJson string = `{"username":"mystique09","password":"testpassword","email":"testemail@gmail.com"}`
var testUserId uuid.UUID

func TestGetUsersRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testRoute.getUsers(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEqual(t, userRequestJson, rec.Body.String())
	}

}

func TestCreateUserRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(userRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testRoute.createUser(ctx)) {
		res := Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 1, res.Status)
	}
}
