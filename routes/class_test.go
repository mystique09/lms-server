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

var globalUserId uuid.UUID
var globalAccessToken *jwt.Token
var globalClassroomId uuid.UUID
var globalClassroomInviteCode uuid.UUID

func TestCreateClassroom(t *testing.T) {
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
	parsed_uuid, err := uuid.Parse(signUpUserId)
	if err != nil {
		t.Fail()
	}
	globalUserId = parsed_uuid

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
	parsed_token, err := jwt.Parse(token, keyFunc(testServer.Cfg.JWT_SECRET_KEY))
	if err != nil {
		t.Fail()
	}
	globalAccessToken = parsed_token

	assert.Equal(t, 200, loginRes.Code)

	// create new classroom payload with fields: name, description, rooom, subject, section
	classPayload := `{"name":"testclass","description":"testdescription","room":"testroom","subject":"testsubject","section":"testsection"}`
	createClassReq := httptest.NewRequest("POST", "/", strings.NewReader(classPayload))
	createClassReq.Header.Set("Content-Type", "application/json")
	createClassReq.Header.Set("Authorization", "Bearer "+token)
	createClassRes := httptest.NewRecorder()
	createClassCtx := e.NewContext(createClassReq, createClassRes)
	// set the path params id
	createClassCtx.SetPath("/api/v1/users/:id/classrooms")
	createClassCtx.SetParamNames("id")
	createClassCtx.SetParamValues(signUpUserId)
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

	parsed_uuid, err = uuid.Parse(createMap["id"].(string))
	if err != nil {
		t.Fail()
	}
	globalClassroomId = parsed_uuid
	globalClassroomInviteCode = uuid.MustParse(createMap["invite_code"].(string))
}

func TestUpdateClassroomNameDescriptionSectionAndRoom(t *testing.T) {
	// update payload for TestUpdateClassroomNameDescriptionSectionAndRoom
	newInviteCode := uuid.New()
	updatePayload := fmt.Sprintf(`{"name":"new_testclass","description":"testdescription","room":"testroom","subject":"testsubject","section":"testsection","invite_code":"%v"}`, newInviteCode)
	updateReq := httptest.NewRequest("PUT", "/", strings.NewReader(updatePayload))
	updateReq.Header.Set("Content-Type", "application/json")
	updateReq.Header.Set("Authorization", "Bearer "+globalAccessToken.Raw)
	updateRes := httptest.NewRecorder()
	updateCtx := e.NewContext(updateReq, updateRes)
	updateCtx.SetPath("/api/v1/classrooms/:id")
	updateCtx.SetParamNames("id")
	updateCtx.SetParamValues(globalClassroomId.String())
	updateCtx.Set("user", globalAccessToken)

	// update the classroom

	if assert.NoError(t, testServer.updateClassroom(updateCtx)) {
		// get the response
		updateResult := utils.Response{}
		err := utils.GetJson(updateRes.Body, &updateResult)
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, 200, updateRes.Code)
		// get the classroom id from the response
		updateData := updateResult.Data
		updateMap := updateData.(map[string]interface{})

		// assert the results
		assert.Equal(t, "new_testclass", updateMap["name"])
		assert.Equal(t, "testdescription", updateMap["description"])
		assert.Equal(t, "testroom", updateMap["room"])
		assert.Equal(t, "testsubject", updateMap["subject"])
		assert.Equal(t, "testsection", updateMap["section"])

		parsed_uuid, err := uuid.Parse(updateMap["id"].(string))
		globalClassroomId = parsed_uuid
	}

	globalClassroomInviteCode = newInviteCode
	//testServer.DB.DeleteUser(context.Background(), globalUserId)
}

var joinAccessToken *jwt.Token
var joinUserId string
var joinClassId string

func TestUserJoinClassroom(t *testing.T) {
	// create new user payload: username, password, email
	payload := `{"username":"testuser2020","password":"testpassword","email":"testemail203p@gmail.com"}`
	signupReq := httptest.NewRequest("POST", "/api/v1/signup", strings.NewReader(payload))
	signupReq.Header.Set("Content-Type", "application/json")
	signupRes := httptest.NewRecorder()
	signupCtx := e.NewContext(signupReq, signupRes)
	// set the middleware
	signupCtx.Echo().Use(JwtAuthMiddleware(testServer.Cfg))
	// create new user
	err := testServer.createUser(signupCtx)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	// get the response
	signupResult := utils.Response{}
	err = utils.GetJson(signupRes.Body, &signupResult)
	if err != nil {
		t.Fail()
	}

	// get the user id from the response
	signupData := signupResult.Data
	signupMap := signupData.(map[string]interface{})

	// store the user id to a new varisble called userId
	userId := uuid.MustParse(signupMap["id"].(string))

	// login the new user
	loginPayload := `{"username":"testuser2020","password":"testpassword"}`
	loginReq := httptest.NewRequest("POST", "/api/v1/login", strings.NewReader(loginPayload))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRes := httptest.NewRecorder()
	loginCtx := e.NewContext(loginReq, loginRes)
	// set the middleware
	loginCtx.Echo().Use(JwtAuthMiddleware(testServer.Cfg))
	// login the new user
	err = testServer.loginHandler(loginCtx)
	if err != nil {
		t.Errorf("Error logging in user: %v", err)
	}

	// get the response
	loginResult := utils.Response{}
	err = utils.GetJson(loginRes.Body, &loginResult)
	if err != nil {
		t.Fail()
	}
	// get the access_token
	accessToken := loginResult.Data.(map[string]interface{})["access_token"].(string)
	// join a classroom using the inviteCode = globalClassroomInviteCode
	joinReq := httptest.NewRequest("POST", "/", strings.NewReader(fmt.Sprintf(`{"invite_code":"%v"}`, globalClassroomInviteCode.String())))
	joinReq.Header.Set("Content-Type", "application/json")
	// set the path to /api/v1/users:id/classrooms
	joinReq.Header.Set("Authorization", "Bearer "+accessToken)
	joinRes := httptest.NewRecorder()
	joinCtx := e.NewContext(joinReq, joinRes)
	joinCtx.SetPath("/api/v1/users/:id/classrooms")
	joinCtx.SetParamNames("id")
	joinCtx.SetParamValues(userId.String())

	parse_token, err := jwt.Parse(accessToken, keyFunc(testServer.Cfg.JWT_SECRET_KEY))
	if err != nil {
		t.Fail()
	}
	joinCtx.Set("user", parse_token)

	// join the classroom
	if assert.NoError(t, testServer.joinClassroom(joinCtx)) {
		// get the response
		joinResult := utils.Response{}
		err := utils.GetJson(joinRes.Body, &joinResult)
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, 200, joinRes.Code)
		// get the classroom id from the response
		joinData := joinResult.Data
		joinMap := joinData.(map[string]interface{})

		// assert the results
		assert.Equal(t, userId.String(), joinMap["user_id"].(string))
		assert.Equal(t, globalClassroomId.String(), joinMap["class_id"].(string))
		//testServer.DB.DeleteUser(context.Background(), userId)
		joinUserId = joinMap["user_id"].(string)
		joinClassId = joinMap["class_id"].(string)
		joinAccessToken = parse_token
	}
}

func TestUserLeaveAClassroom(t *testing.T) {

	leaveReq := httptest.NewRequest("DELETE", "/", nil)
	// set the path to /api/v1/users/:id/classrooms/:class_id
	leaveReq.Header.Set("Authorization", "Bearer "+joinAccessToken.Raw)
	leaveRes := httptest.NewRecorder()
	leaveCtx := e.NewContext(leaveReq, leaveRes)
	leaveCtx.SetPath("/api/v1/users/:id/classrooms/:class_id")
	leaveCtx.SetParamNames("id", "class_id")
	leaveCtx.SetParamValues(joinUserId, joinClassId)
	leaveCtx.Set("user", joinAccessToken)
	// leave the classroom
	if assert.NoError(t, testServer.leaveClassroom(leaveCtx)) {
		// get the response
		leaveResult := utils.Response{}
		err := utils.GetJson(leaveRes.Body, &leaveResult)
		if err != nil {
			t.Fail()
		}

		assert.Equal(t, 200, leaveRes.Code)
		// get the classroom id from the response
		leaveData := leaveResult.Data
		leaveMap := leaveData.(map[string]interface{})

		// assert the results
		assert.Equal(t, joinUserId, leaveMap["user_id"].(string))
		assert.Equal(t, joinClassId, leaveMap["class_id"].(string))
		testServer.DB.DeleteUser(context.Background(), uuid.MustParse(joinUserId))
	}
}

func TestDeleteClassroom(t *testing.T) {
	deleteReq := httptest.NewRequest("DELETE", "/", nil)
	deleteReq.Header.Set("Content-Type", "application/json")
	deleteReq.Header.Set("Authorization", "Bearer "+globalAccessToken.Raw)
	deleteRes := httptest.NewRecorder()
	deleteCtx := e.NewContext(deleteReq, deleteRes)
	deleteCtx.SetPath("/api/v1/classrooms/:id")
	deleteCtx.SetParamNames("id")
	deleteCtx.SetParamValues(globalClassroomId.String())
	deleteCtx.Set("user", globalAccessToken)

	// delete the TestDeleteClassroom
	if assert.NoError(t, testServer.deleteClassroom(deleteCtx)) {
		// get the response
		deleteResult := utils.Response{}
		err := utils.GetJson(deleteRes.Body, &deleteResult)
		if err != nil {
			t.Fail()
		}
		deleteData := deleteResult.Data.(map[string]interface{})

		assert.Equal(t, 200, deleteRes.Code)
		assert.Equal(t, globalClassroomId.String(), deleteData["id"])
	}
	testServer.DB.DeleteUser(context.Background(), globalUserId)
}
