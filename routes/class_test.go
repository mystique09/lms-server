package routes

import (
	"context"
	"fmt"
	"net/http/httptest"
	"server/utils"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCresteClassroom(t *testing.T) {
	// signup new user
	userPayload := `{"username":"testuser","password":"testpassword","email":"testemailhshs@gmail.com"}`
	signupReq := httptest.NewRequest("POST", "/api/v1/users", strings.NewReader(userPayload))
	signupReq.Header.Set("Content-Type", "application/json")
	signupRes := httptest.NewRecorder()
	signupCtx := e.NewContext(signupReq, signupRes)

	err := testServer.createUser(signupCtx)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}
	assert.Equal(t, 200, signupRes.Code)
	signUpRes := utils.Response{}
	err = utils.GetJson(signupRes.Body, &signUpRes)
	if err != nil {
		t.Errorf("Error parsing response: %v", err)
	}
	signUpMap := signUpRes.Data.(map[string]interface{})
	signUpUserId := fmt.Sprintf("%v", signUpMap["id"].(string))

	// login user
	loginPayload := `{"username":"testuser","password":"testpassword"}`
	loginReq := httptest.NewRequest("POST", "/api/v1/login", strings.NewReader(loginPayload))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRes := httptest.NewRecorder()
	loginCtx := e.NewContext(loginReq, loginRes)

	err = testServer.loginHandler(loginCtx)
	if err != nil {
		t.Errorf("Error logging in user: %v", err)
	}
	// convert the response to a map
	loginResult := utils.Response{}
	err = utils.GetJson(loginRes.Body, &loginResult)
	if err != nil {
		t.Fail()
	}

	// get the token from the response
	loginData := loginResult.Data
	loginMap := loginData.(map[string]interface{})
	token := fmt.Sprintf("%v", loginMap["access_token"])

	assert.Equal(t, 200, loginRes.Code)

	// create new classroom payload with fields: name, description, rooom, subject, section
	classPayload := `{"name":"testclass","description":"testdescription","room":"testroom","subject":"testsubject","section":"testsection"}`
	createClassReq := httptest.NewRequest("POST", "/", strings.NewReader(classPayload))
	createClassReq.Header.Set("Content-Type", "application/json")
	createClassReq.Header.Set("Authorization", "Bearer "+token)
	createClassRes := httptest.NewRecorder()
	createClassCtx := e.NewContext(createClassReq, createClassRes)
	// set the path params id
	createClassCtx.SetParamNames("id")
	createClassCtx.SetParamValues(signUpUserId)
	parsed_token, err := jwt.Parse(token, keyFunc(testServer.Cfg.JWT_SECRET_KEY))
	if err != nil {
		t.Fail()
	}
	createClassCtx.Set("user", parsed_token)

	// set the middleware
	createClassCtx.Echo().Use(JwtAuthMiddleware(testServer.Cfg))

	// create classroom
	err = testServer.createNewClassroom(createClassCtx)
	if err != nil {
		t.Errorf("Error creating classroom: %v", err)
	}

	// get the response
	createClassResult := utils.Response{}
	err = utils.GetJson(createClassRes.Body, &createClassResult)
	if err != nil {
		t.Fail()
	}

	// get the classroom id from the response
	createClassData := createClassResult.Data
	createMap := createClassData.(map[string]interface{})

	// assert the results
	assert.Equal(t, 200, createClassRes.Code)
	assert.Equal(t, "testclass", createMap["name"])
	assert.Equal(t, "testdescription", createMap["description"])
	assert.Equal(t, "testroom", createMap["room"])
	assert.Equal(t, "testsubject", createMap["subject"])
	assert.Equal(t, "testsection", createMap["section"])

	parsed_uuid, err := uuid.Parse(signUpUserId)
	if err != nil {
		t.Fail()
	}
	testServer.DB.DeleteUser(context.Background(), parsed_uuid)
}
