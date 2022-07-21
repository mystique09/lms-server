package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/utils"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Payload struct {
	Status  int
	Data    string
	Message string
}

var authTestUserId string
var authTestToken string

func TestLoginWithOneAccount(t *testing.T) {
	var payload string = `{"username":"mystique07","password":"testpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testServer.loginHandler(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}

		assert.Equal(t, 404, rec.Code)
		assert.Equal(t, "User doesn't exist.", res.Error)
	}
}

func TestSignupRoute(t *testing.T) {
	var userRequestJson string = `{"username":"mystique008","password":"testpassword","email":"testemail2@gmail.com"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(userRequestJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testServer.createUser(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		assert.Equal(t, 200, rec.Code)
		user := res.Data.(map[string]interface{})
		assert.Equal(t, "STUDENT", user["user_role"])
		assert.Equal(t, "PUBLIC", user["visibility"])
		assert.Empty(t, user["password"])
		assert.Equal(t, "testemail2@gmail.com", user["email"])
		authTestUserId = user["id"].(string)
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
		req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(p.Data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, testServer.loginHandler(ctx)) {
			res := utils.Response{}
			if err := utils.GetJson(rec.Body, &res); err != nil {
				t.Fail()
			}

			assert.Equal(t, p.Status, rec.Code)
			assert.Equal(t, p.Message, res.Error)
		}
	}
}

func TestLoginSuccessful(t *testing.T) {
	var payload string = `{"username":"mystique008","password":"testpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, testServer.loginHandler(ctx)) {
		assert.Equal(t, 200, rec.Code)
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		success_resp := res.Data.(map[string]interface{})
		assert.Equal(t, success_resp["message"], "Logged in.")
		assert.NotEmpty(t, success_resp["access_token"])
		assert.NotEmpty(t, success_resp["refresh_token"])
		authTestToken = success_resp["access_token"].(string)
	}
}

func keyFunc(secret_key []byte) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("Unexpected signing method=%v", t.Header["alg"])
		}
		return secret_key, nil
	}
}

func TestDeleteUserAfterSignup(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	parsed_token, err := jwt.Parse(authTestToken, keyFunc(testServer.Cfg.JWT_SECRET_KEY))
	if err != nil {
		t.Fail()
	}

	ctx.SetPath("/api/v1/users/:id")
	ctx.Set("user", parsed_token)
	ctx.SetParamNames("id")
	ctx.SetParamValues(authTestUserId)
	ctx.Echo().Use(JwtAuthMiddleware(server.Cfg))

	if assert.NoError(t, testServer.deleteUser(ctx)) {
		res := utils.Response{}
		if err := utils.GetJson(rec.Body, &res); err != nil {
			t.Fail()
		}
		deleted_user := res.Data.(map[string]interface{})
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, authTestUserId, deleted_user["id"])
	}
}
