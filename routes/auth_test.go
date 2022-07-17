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

type Payload struct {
	Status  int
	Data    string
	Message string
}

func TestLoginWithOneAccount(t *testing.T) {
	var payload string = `{"username":"mystique07","password":"testpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testRoute.loginRoute(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}

		assert.Equal(t, 404, rec.Code)
		assert.Equal(t, "User doesn't exist.", res.Error)
	}
}

// loop through each payload and Test
func TestLoginWithManyInvalidAccounts(t *testing.T) {
	var payloads []Payload = []Payload{
		{
			Status:  404,
			Data:    `{"username":"jzudhsjsj","password":"jsjsjejeu"}`,
			Message: "User doesn't exist.",
		},
		{
			Status:  404,
			Data:    `{"username":"jzudhsjsj","password":"jsjsjejeu"}`,
			Message: "User doesn't exist.",
		},
		{
			Status:  400,
			Data:    `{"username":"jzudhsjsj"}`,
			Message: "One field might be missing, fill in the missing fields.",
		},
		{
			Status:  400,
			Data:    `{"username":"jzudhsjsj", "password":""}`,
			Message: "One field might be missing, fill in the missing fields.",
		},
		{
			Status:  403,
			Data:    `{"username":"mystique008","password":"testpasswor"}`,
			Message: "Incorrect username or password.",
		},
	}

	// loop through each payload and Test
	for _, p := range payloads {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(p.Data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, testRoute.loginRoute(ctx)) {
			res := utils.Response{}
			if err := utils.GetJson(rec.Body, &res); err != nil {
				t.Fail()
			}

			assert.Equal(t, p.Status, rec.Code)
			assert.Equal(t, p.Message, res.Error)
		}
	}
}

func TestLoginSucfessful(t *testing.T) {
	var payload string = `{"username":"mystique007","password":"testpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testRoute.loginRoute(ctx)) {
		assert.Equal(t, 200, rec.Code)
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		success_resp := res.Data.(map[string]interface{})
		assert.Equal(t, success_resp["message"], "Logged in.")
		assert.NotEmpty(t, success_resp["access_token"])
		assert.NotEmpty(t, success_resp["refresh_token"])
	}
}
