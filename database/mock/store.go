// Code generated by MockGen. DO NOT EDIT.
// Source: server/database/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"
	database "server/database/sqlc"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddNewClassroomMember mocks base method.
func (m *MockStore) AddNewClassroomMember(arg0 context.Context, arg1 database.AddNewClassroomMemberParams) (database.ClassroomMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewClassroomMember", arg0, arg1)
	ret0, _ := ret[0].(database.ClassroomMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNewClassroomMember indicates an expected call of AddNewClassroomMember.
func (mr *MockStoreMockRecorder) AddNewClassroomMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewClassroomMember", reflect.TypeOf((*MockStore)(nil).AddNewClassroomMember), arg0, arg1)
}

// CreateClass mocks base method.
func (m *MockStore) CreateClass(arg0 context.Context, arg1 database.CreateClassParams) (database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClass", arg0, arg1)
	ret0, _ := ret[0].(database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClass indicates an expected call of CreateClass.
func (mr *MockStoreMockRecorder) CreateClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClass", reflect.TypeOf((*MockStore)(nil).CreateClass), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 database.CreateUserParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteClass mocks base method.
func (m *MockStore) DeleteClass(arg0 context.Context, arg1 uuid.UUID) (database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClass", arg0, arg1)
	ret0, _ := ret[0].(database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteClass indicates an expected call of DeleteClass.
func (mr *MockStoreMockRecorder) DeleteClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClass", reflect.TypeOf((*MockStore)(nil).DeleteClass), arg0, arg1)
}

// DeleteClassworkFromClass mocks base method.
func (m *MockStore) DeleteClassworkFromClass(arg0 context.Context, arg1 database.DeleteClassworkFromClassParams) (database.ClassWork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClassworkFromClass", arg0, arg1)
	ret0, _ := ret[0].(database.ClassWork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteClassworkFromClass indicates an expected call of DeleteClassworkFromClass.
func (mr *MockStoreMockRecorder) DeleteClassworkFromClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClassworkFromClass", reflect.TypeOf((*MockStore)(nil).DeleteClassworkFromClass), arg0, arg1)
}

// DeleteCommentFromPost mocks base method.
func (m *MockStore) DeleteCommentFromPost(arg0 context.Context, arg1 database.DeleteCommentFromPostParams) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommentFromPost", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCommentFromPost indicates an expected call of DeleteCommentFromPost.
func (mr *MockStoreMockRecorder) DeleteCommentFromPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentFromPost", reflect.TypeOf((*MockStore)(nil).DeleteCommentFromPost), arg0, arg1)
}

// DeletePostFromClass mocks base method.
func (m *MockStore) DeletePostFromClass(arg0 context.Context, arg1 database.DeletePostFromClassParams) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostFromClass", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePostFromClass indicates an expected call of DeletePostFromClass.
func (mr *MockStoreMockRecorder) DeletePostFromClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostFromClass", reflect.TypeOf((*MockStore)(nil).DeletePostFromClass), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// FollowUser mocks base method.
func (m *MockStore) FollowUser(arg0 context.Context, arg1 database.FollowUserParams) (database.UserFollow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FollowUser", arg0, arg1)
	ret0, _ := ret[0].(database.UserFollow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FollowUser indicates an expected call of FollowUser.
func (mr *MockStoreMockRecorder) FollowUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FollowUser", reflect.TypeOf((*MockStore)(nil).FollowUser), arg0, arg1)
}

// GetAllClassFromUser mocks base method.
func (m *MockStore) GetAllClassFromUser(arg0 context.Context, arg1 database.GetAllClassFromUserParams) ([]database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllClassFromUser", arg0, arg1)
	ret0, _ := ret[0].([]database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllClassFromUser indicates an expected call of GetAllClassFromUser.
func (mr *MockStoreMockRecorder) GetAllClassFromUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllClassFromUser", reflect.TypeOf((*MockStore)(nil).GetAllClassFromUser), arg0, arg1)
}

// GetAllClassroomMembers mocks base method.
func (m *MockStore) GetAllClassroomMembers(arg0 context.Context, arg1 database.GetAllClassroomMembersParams) ([]database.GetAllClassroomMembersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllClassroomMembers", arg0, arg1)
	ret0, _ := ret[0].([]database.GetAllClassroomMembersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllClassroomMembers indicates an expected call of GetAllClassroomMembers.
func (mr *MockStoreMockRecorder) GetAllClassroomMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllClassroomMembers", reflect.TypeOf((*MockStore)(nil).GetAllClassroomMembers), arg0, arg1)
}

// GetAllCommentLikes mocks base method.
func (m *MockStore) GetAllCommentLikes(arg0 context.Context, arg1 uuid.UUID) ([]database.CommentLike, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCommentLikes", arg0, arg1)
	ret0, _ := ret[0].([]database.CommentLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCommentLikes indicates an expected call of GetAllCommentLikes.
func (mr *MockStoreMockRecorder) GetAllCommentLikes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCommentLikes", reflect.TypeOf((*MockStore)(nil).GetAllCommentLikes), arg0, arg1)
}

// GetAllCommentsFromPost mocks base method.
func (m *MockStore) GetAllCommentsFromPost(arg0 context.Context, arg1 database.GetAllCommentsFromPostParams) ([]database.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCommentsFromPost", arg0, arg1)
	ret0, _ := ret[0].([]database.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCommentsFromPost indicates an expected call of GetAllCommentsFromPost.
func (mr *MockStoreMockRecorder) GetAllCommentsFromPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCommentsFromPost", reflect.TypeOf((*MockStore)(nil).GetAllCommentsFromPost), arg0, arg1)
}

// GetAllFollowers mocks base method.
func (m *MockStore) GetAllFollowers(arg0 context.Context, arg1 database.GetAllFollowersParams) ([]database.GetAllFollowersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFollowers", arg0, arg1)
	ret0, _ := ret[0].([]database.GetAllFollowersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFollowers indicates an expected call of GetAllFollowers.
func (mr *MockStoreMockRecorder) GetAllFollowers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFollowers", reflect.TypeOf((*MockStore)(nil).GetAllFollowers), arg0, arg1)
}

// GetAllFollowing mocks base method.
func (m *MockStore) GetAllFollowing(arg0 context.Context, arg1 database.GetAllFollowingParams) ([]database.GetAllFollowingRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFollowing", arg0, arg1)
	ret0, _ := ret[0].([]database.GetAllFollowingRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFollowing indicates an expected call of GetAllFollowing.
func (mr *MockStoreMockRecorder) GetAllFollowing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFollowing", reflect.TypeOf((*MockStore)(nil).GetAllFollowing), arg0, arg1)
}

// GetAllJoinedClassrooms mocks base method.
func (m *MockStore) GetAllJoinedClassrooms(arg0 context.Context, arg1 database.GetAllJoinedClassroomsParams) ([]database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllJoinedClassrooms", arg0, arg1)
	ret0, _ := ret[0].([]database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJoinedClassrooms indicates an expected call of GetAllJoinedClassrooms.
func (mr *MockStoreMockRecorder) GetAllJoinedClassrooms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJoinedClassrooms", reflect.TypeOf((*MockStore)(nil).GetAllJoinedClassrooms), arg0, arg1)
}

// GetAllPostLikes mocks base method.
func (m *MockStore) GetAllPostLikes(arg0 context.Context, arg1 uuid.UUID) ([]database.PostLike, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPostLikes", arg0, arg1)
	ret0, _ := ret[0].([]database.PostLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPostLikes indicates an expected call of GetAllPostLikes.
func (mr *MockStoreMockRecorder) GetAllPostLikes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPostLikes", reflect.TypeOf((*MockStore)(nil).GetAllPostLikes), arg0, arg1)
}

// GetClass mocks base method.
func (m *MockStore) GetClass(arg0 context.Context, arg1 uuid.UUID) (database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClass", arg0, arg1)
	ret0, _ := ret[0].(database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClass indicates an expected call of GetClass.
func (mr *MockStoreMockRecorder) GetClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClass", reflect.TypeOf((*MockStore)(nil).GetClass), arg0, arg1)
}

// GetClassWork mocks base method.
func (m *MockStore) GetClassWork(arg0 context.Context, arg1 database.GetClassWorkParams) (database.ClassWork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassWork", arg0, arg1)
	ret0, _ := ret[0].(database.ClassWork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassWork indicates an expected call of GetClassWork.
func (mr *MockStoreMockRecorder) GetClassWork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassWork", reflect.TypeOf((*MockStore)(nil).GetClassWork), arg0, arg1)
}

// GetClassroomMemberById mocks base method.
func (m *MockStore) GetClassroomMemberById(arg0 context.Context, arg1 database.GetClassroomMemberByIdParams) (database.ClassroomMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassroomMemberById", arg0, arg1)
	ret0, _ := ret[0].(database.ClassroomMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassroomMemberById indicates an expected call of GetClassroomMemberById.
func (mr *MockStoreMockRecorder) GetClassroomMemberById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassroomMemberById", reflect.TypeOf((*MockStore)(nil).GetClassroomMemberById), arg0, arg1)
}

// GetClassroomWithInviteCode mocks base method.
func (m *MockStore) GetClassroomWithInviteCode(arg0 context.Context, arg1 uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassroomWithInviteCode", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassroomWithInviteCode indicates an expected call of GetClassroomWithInviteCode.
func (mr *MockStoreMockRecorder) GetClassroomWithInviteCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassroomWithInviteCode", reflect.TypeOf((*MockStore)(nil).GetClassroomWithInviteCode), arg0, arg1)
}

// GetFollowerById mocks base method.
func (m *MockStore) GetFollowerById(arg0 context.Context, arg1 uuid.UUID) (database.GetFollowerByIdRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowerById", arg0, arg1)
	ret0, _ := ret[0].(database.GetFollowerByIdRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowerById indicates an expected call of GetFollowerById.
func (mr *MockStoreMockRecorder) GetFollowerById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowerById", reflect.TypeOf((*MockStore)(nil).GetFollowerById), arg0, arg1)
}

// GetOneFollower mocks base method.
func (m *MockStore) GetOneFollower(arg0 context.Context, arg1 database.GetOneFollowerParams) (database.UserFollow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneFollower", arg0, arg1)
	ret0, _ := ret[0].(database.UserFollow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneFollower indicates an expected call of GetOneFollower.
func (mr *MockStoreMockRecorder) GetOneFollower(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneFollower", reflect.TypeOf((*MockStore)(nil).GetOneFollower), arg0, arg1)
}

// GetOnePost mocks base method.
func (m *MockStore) GetOnePost(arg0 context.Context, arg1 uuid.UUID) (database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOnePost", arg0, arg1)
	ret0, _ := ret[0].(database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOnePost indicates an expected call of GetOnePost.
func (mr *MockStoreMockRecorder) GetOnePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOnePost", reflect.TypeOf((*MockStore)(nil).GetOnePost), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 uuid.UUID) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 context.Context, arg1 string) (database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockStore) GetUsers(arg0 context.Context, arg1 int32) ([]database.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].([]database.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockStoreMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStore)(nil).GetUsers), arg0, arg1)
}

// InsertNewClasswork mocks base method.
func (m *MockStore) InsertNewClasswork(arg0 context.Context, arg1 database.InsertNewClassworkParams) (database.ClassWork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewClasswork", arg0, arg1)
	ret0, _ := ret[0].(database.ClassWork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNewClasswork indicates an expected call of InsertNewClasswork.
func (mr *MockStoreMockRecorder) InsertNewClasswork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewClasswork", reflect.TypeOf((*MockStore)(nil).InsertNewClasswork), arg0, arg1)
}

// InsertNewCommentInPost mocks base method.
func (m *MockStore) InsertNewCommentInPost(arg0 context.Context, arg1 database.InsertNewCommentInPostParams) (database.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewCommentInPost", arg0, arg1)
	ret0, _ := ret[0].(database.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNewCommentInPost indicates an expected call of InsertNewCommentInPost.
func (mr *MockStoreMockRecorder) InsertNewCommentInPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewCommentInPost", reflect.TypeOf((*MockStore)(nil).InsertNewCommentInPost), arg0, arg1)
}

// InsertNewPost mocks base method.
func (m *MockStore) InsertNewPost(arg0 context.Context, arg1 database.InsertNewPostParams) (database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewPost", arg0, arg1)
	ret0, _ := ret[0].(database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNewPost indicates an expected call of InsertNewPost.
func (mr *MockStoreMockRecorder) InsertNewPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewPost", reflect.TypeOf((*MockStore)(nil).InsertNewPost), arg0, arg1)
}

// LeaveClassroom mocks base method.
func (m *MockStore) LeaveClassroom(arg0 context.Context, arg1 database.LeaveClassroomParams) (database.ClassroomMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveClassroom", arg0, arg1)
	ret0, _ := ret[0].(database.ClassroomMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaveClassroom indicates an expected call of LeaveClassroom.
func (mr *MockStoreMockRecorder) LeaveClassroom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveClassroom", reflect.TypeOf((*MockStore)(nil).LeaveClassroom), arg0, arg1)
}

// LikeComment mocks base method.
func (m *MockStore) LikeComment(arg0 context.Context, arg1 database.LikeCommentParams) (database.CommentLike, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikeComment", arg0, arg1)
	ret0, _ := ret[0].(database.CommentLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LikeComment indicates an expected call of LikeComment.
func (mr *MockStoreMockRecorder) LikeComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikeComment", reflect.TypeOf((*MockStore)(nil).LikeComment), arg0, arg1)
}

// LikePost mocks base method.
func (m *MockStore) LikePost(arg0 context.Context, arg1 database.LikePostParams) (database.PostLike, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikePost", arg0, arg1)
	ret0, _ := ret[0].(database.PostLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LikePost indicates an expected call of LikePost.
func (mr *MockStoreMockRecorder) LikePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikePost", reflect.TypeOf((*MockStore)(nil).LikePost), arg0, arg1)
}

// ListAllPostsByUser mocks base method.
func (m *MockStore) ListAllPostsByUser(arg0 context.Context, arg1 database.ListAllPostsByUserParams) ([]database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllPostsByUser", arg0, arg1)
	ret0, _ := ret[0].([]database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllPostsByUser indicates an expected call of ListAllPostsByUser.
func (mr *MockStoreMockRecorder) ListAllPostsByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllPostsByUser", reflect.TypeOf((*MockStore)(nil).ListAllPostsByUser), arg0, arg1)
}

// ListAllPostsFromClass mocks base method.
func (m *MockStore) ListAllPostsFromClass(arg0 context.Context, arg1 database.ListAllPostsFromClassParams) ([]database.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllPostsFromClass", arg0, arg1)
	ret0, _ := ret[0].([]database.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllPostsFromClass indicates an expected call of ListAllPostsFromClass.
func (mr *MockStoreMockRecorder) ListAllPostsFromClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllPostsFromClass", reflect.TypeOf((*MockStore)(nil).ListAllPostsFromClass), arg0, arg1)
}

// ListAllPublicClass mocks base method.
func (m *MockStore) ListAllPublicClass(arg0 context.Context, arg1 int32) ([]database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllPublicClass", arg0, arg1)
	ret0, _ := ret[0].([]database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllPublicClass indicates an expected call of ListAllPublicClass.
func (mr *MockStoreMockRecorder) ListAllPublicClass(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllPublicClass", reflect.TypeOf((*MockStore)(nil).ListAllPublicClass), arg0, arg1)
}

// ListClassworkAdmin mocks base method.
func (m *MockStore) ListClassworkAdmin(arg0 context.Context, arg1 database.ListClassworkAdminParams) ([]database.ClassWork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClassworkAdmin", arg0, arg1)
	ret0, _ := ret[0].([]database.ClassWork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClassworkAdmin indicates an expected call of ListClassworkAdmin.
func (mr *MockStoreMockRecorder) ListClassworkAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClassworkAdmin", reflect.TypeOf((*MockStore)(nil).ListClassworkAdmin), arg0, arg1)
}

// ListSubmittedClassworks mocks base method.
func (m *MockStore) ListSubmittedClassworks(arg0 context.Context, arg1 database.ListSubmittedClassworksParams) ([]database.ClassWork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubmittedClassworks", arg0, arg1)
	ret0, _ := ret[0].([]database.ClassWork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubmittedClassworks indicates an expected call of ListSubmittedClassworks.
func (mr *MockStoreMockRecorder) ListSubmittedClassworks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubmittedClassworks", reflect.TypeOf((*MockStore)(nil).ListSubmittedClassworks), arg0, arg1)
}

// UnfollowUser mocks base method.
func (m *MockStore) UnfollowUser(arg0 context.Context, arg1 uuid.UUID) (database.UserFollow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnfollowUser", arg0, arg1)
	ret0, _ := ret[0].(database.UserFollow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnfollowUser indicates an expected call of UnfollowUser.
func (mr *MockStoreMockRecorder) UnfollowUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnfollowUser", reflect.TypeOf((*MockStore)(nil).UnfollowUser), arg0, arg1)
}

// UnlikeComment mocks base method.
func (m *MockStore) UnlikeComment(arg0 context.Context, arg1 database.UnlikeCommentParams) (database.CommentLike, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlikeComment", arg0, arg1)
	ret0, _ := ret[0].(database.CommentLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnlikeComment indicates an expected call of UnlikeComment.
func (mr *MockStoreMockRecorder) UnlikeComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlikeComment", reflect.TypeOf((*MockStore)(nil).UnlikeComment), arg0, arg1)
}

// UnlikePost mocks base method.
func (m *MockStore) UnlikePost(arg0 context.Context, arg1 database.UnlikePostParams) (database.PostLike, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlikePost", arg0, arg1)
	ret0, _ := ret[0].(database.PostLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnlikePost indicates an expected call of UnlikePost.
func (mr *MockStoreMockRecorder) UnlikePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlikePost", reflect.TypeOf((*MockStore)(nil).UnlikePost), arg0, arg1)
}

// UpdateAClassworkMark mocks base method.
func (m *MockStore) UpdateAClassworkMark(arg0 context.Context, arg1 database.UpdateAClassworkMarkParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAClassworkMark", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAClassworkMark indicates an expected call of UpdateAClassworkMark.
func (mr *MockStoreMockRecorder) UpdateAClassworkMark(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAClassworkMark", reflect.TypeOf((*MockStore)(nil).UpdateAClassworkMark), arg0, arg1)
}

// UpdateClassroomDescription mocks base method.
func (m *MockStore) UpdateClassroomDescription(arg0 context.Context, arg1 database.UpdateClassroomDescriptionParams) (database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomDescription", arg0, arg1)
	ret0, _ := ret[0].(database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClassroomDescription indicates an expected call of UpdateClassroomDescription.
func (mr *MockStoreMockRecorder) UpdateClassroomDescription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomDescription", reflect.TypeOf((*MockStore)(nil).UpdateClassroomDescription), arg0, arg1)
}

// UpdateClassroomInviteCode mocks base method.
func (m *MockStore) UpdateClassroomInviteCode(arg0 context.Context, arg1 database.UpdateClassroomInviteCodeParams) (database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomInviteCode", arg0, arg1)
	ret0, _ := ret[0].(database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClassroomInviteCode indicates an expected call of UpdateClassroomInviteCode.
func (mr *MockStoreMockRecorder) UpdateClassroomInviteCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomInviteCode", reflect.TypeOf((*MockStore)(nil).UpdateClassroomInviteCode), arg0, arg1)
}

// UpdateClassroomName mocks base method.
func (m *MockStore) UpdateClassroomName(arg0 context.Context, arg1 database.UpdateClassroomNameParams) (database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomName", arg0, arg1)
	ret0, _ := ret[0].(database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClassroomName indicates an expected call of UpdateClassroomName.
func (mr *MockStoreMockRecorder) UpdateClassroomName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomName", reflect.TypeOf((*MockStore)(nil).UpdateClassroomName), arg0, arg1)
}

// UpdateClassroomRoom mocks base method.
func (m *MockStore) UpdateClassroomRoom(arg0 context.Context, arg1 database.UpdateClassroomRoomParams) (database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomRoom", arg0, arg1)
	ret0, _ := ret[0].(database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClassroomRoom indicates an expected call of UpdateClassroomRoom.
func (mr *MockStoreMockRecorder) UpdateClassroomRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomRoom", reflect.TypeOf((*MockStore)(nil).UpdateClassroomRoom), arg0, arg1)
}

// UpdateClassroomSection mocks base method.
func (m *MockStore) UpdateClassroomSection(arg0 context.Context, arg1 database.UpdateClassroomSectionParams) (database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomSection", arg0, arg1)
	ret0, _ := ret[0].(database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClassroomSection indicates an expected call of UpdateClassroomSection.
func (mr *MockStoreMockRecorder) UpdateClassroomSection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomSection", reflect.TypeOf((*MockStore)(nil).UpdateClassroomSection), arg0, arg1)
}

// UpdateClassroomSubject mocks base method.
func (m *MockStore) UpdateClassroomSubject(arg0 context.Context, arg1 database.UpdateClassroomSubjectParams) (database.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomSubject", arg0, arg1)
	ret0, _ := ret[0].(database.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClassroomSubject indicates an expected call of UpdateClassroomSubject.
func (mr *MockStoreMockRecorder) UpdateClassroomSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomSubject", reflect.TypeOf((*MockStore)(nil).UpdateClassroomSubject), arg0, arg1)
}

// UpdateCommentContentInPost mocks base method.
func (m *MockStore) UpdateCommentContentInPost(arg0 context.Context, arg1 database.UpdateCommentContentInPostParams) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCommentContentInPost", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCommentContentInPost indicates an expected call of UpdateCommentContentInPost.
func (mr *MockStoreMockRecorder) UpdateCommentContentInPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommentContentInPost", reflect.TypeOf((*MockStore)(nil).UpdateCommentContentInPost), arg0, arg1)
}

// UpdatePostContent mocks base method.
func (m *MockStore) UpdatePostContent(arg0 context.Context, arg1 database.UpdatePostContentParams) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePostContent", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePostContent indicates an expected call of UpdatePostContent.
func (mr *MockStoreMockRecorder) UpdatePostContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePostContent", reflect.TypeOf((*MockStore)(nil).UpdatePostContent), arg0, arg1)
}

// UpdateUserEmail mocks base method.
func (m *MockStore) UpdateUserEmail(arg0 context.Context, arg1 database.UpdateUserEmailParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserEmail", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserEmail indicates an expected call of UpdateUserEmail.
func (mr *MockStoreMockRecorder) UpdateUserEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserEmail", reflect.TypeOf((*MockStore)(nil).UpdateUserEmail), arg0, arg1)
}

// UpdateUserPassword mocks base method.
func (m *MockStore) UpdateUserPassword(arg0 context.Context, arg1 database.UpdateUserPasswordParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockStoreMockRecorder) UpdateUserPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockStore)(nil).UpdateUserPassword), arg0, arg1)
}

// UpdateUsername mocks base method.
func (m *MockStore) UpdateUsername(arg0 context.Context, arg1 database.UpdateUsernameParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsername", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUsername indicates an expected call of UpdateUsername.
func (mr *MockStoreMockRecorder) UpdateUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsername", reflect.TypeOf((*MockStore)(nil).UpdateUsername), arg0, arg1)
}