// Code generated by MockGen. DO NOT EDIT.
// Source: server/domain (interfaces: ClassroomRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	postgresql "server/database/postgresql"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockClassroomRepository is a mock of ClassroomRepository interface.
type MockClassroomRepository struct {
	ctrl     *gomock.Controller
	recorder *MockClassroomRepositoryMockRecorder
}

// MockClassroomRepositoryMockRecorder is the mock recorder for MockClassroomRepository.
type MockClassroomRepositoryMockRecorder struct {
	mock *MockClassroomRepository
}

// NewMockClassroomRepository creates a new mock instance.
func NewMockClassroomRepository(ctrl *gomock.Controller) *MockClassroomRepository {
	mock := &MockClassroomRepository{ctrl: ctrl}
	mock.recorder = &MockClassroomRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClassroomRepository) EXPECT() *MockClassroomRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClassroomRepository) Create(arg0 context.Context, arg1 *postgresql.CreateClassParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockClassroomRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClassroomRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockClassroomRepository) Delete(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClassroomRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClassroomRepository)(nil).Delete), arg0, arg1)
}

// Fetch mocks base method.
func (m *MockClassroomRepository) Fetch(arg0 context.Context, arg1 int32) ([]postgresql.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0, arg1)
	ret0, _ := ret[0].([]postgresql.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockClassroomRepositoryMockRecorder) Fetch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockClassroomRepository)(nil).Fetch), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockClassroomRepository) GetByID(arg0 context.Context, arg1 uuid.UUID) (postgresql.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(postgresql.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockClassroomRepositoryMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockClassroomRepository)(nil).GetByID), arg0, arg1)
}

// GetByInviteCode mocks base method.
func (m *MockClassroomRepository) GetByInviteCode(arg0 context.Context, arg1 uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByInviteCode", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByInviteCode indicates an expected call of GetByInviteCode.
func (mr *MockClassroomRepositoryMockRecorder) GetByInviteCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByInviteCode", reflect.TypeOf((*MockClassroomRepository)(nil).GetByInviteCode), arg0, arg1)
}

// GetClasroomsByUser mocks base method.
func (m *MockClassroomRepository) GetClasroomsByUser(arg0 context.Context, arg1 postgresql.GetAllClassFromUserParams) ([]postgresql.Classroom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClasroomsByUser", arg0, arg1)
	ret0, _ := ret[0].([]postgresql.Classroom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClasroomsByUser indicates an expected call of GetClasroomsByUser.
func (mr *MockClassroomRepositoryMockRecorder) GetClasroomsByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClasroomsByUser", reflect.TypeOf((*MockClassroomRepository)(nil).GetClasroomsByUser), arg0, arg1)
}

// GetClassroomMembers mocks base method.
func (m *MockClassroomRepository) GetClassroomMembers(arg0 context.Context, arg1 uuid.UUID) ([]postgresql.ClassroomMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassroomMembers", arg0, arg1)
	ret0, _ := ret[0].([]postgresql.ClassroomMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassroomMembers indicates an expected call of GetClassroomMembers.
func (mr *MockClassroomRepositoryMockRecorder) GetClassroomMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassroomMembers", reflect.TypeOf((*MockClassroomRepository)(nil).GetClassroomMembers), arg0, arg1)
}

// UpdateClassroomDescription mocks base method.
func (m *MockClassroomRepository) UpdateClassroomDescription(arg0 context.Context, arg1 *postgresql.UpdateClassroomDescriptionParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomDescription", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClassroomDescription indicates an expected call of UpdateClassroomDescription.
func (mr *MockClassroomRepositoryMockRecorder) UpdateClassroomDescription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomDescription", reflect.TypeOf((*MockClassroomRepository)(nil).UpdateClassroomDescription), arg0, arg1)
}

// UpdateClassroomInviteCode mocks base method.
func (m *MockClassroomRepository) UpdateClassroomInviteCode(arg0 context.Context, arg1 *postgresql.UpdateClassroomInviteCodeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomInviteCode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClassroomInviteCode indicates an expected call of UpdateClassroomInviteCode.
func (mr *MockClassroomRepositoryMockRecorder) UpdateClassroomInviteCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomInviteCode", reflect.TypeOf((*MockClassroomRepository)(nil).UpdateClassroomInviteCode), arg0, arg1)
}

// UpdateClassroomName mocks base method.
func (m *MockClassroomRepository) UpdateClassroomName(arg0 context.Context, arg1 *postgresql.UpdateClassroomNameParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomName", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClassroomName indicates an expected call of UpdateClassroomName.
func (mr *MockClassroomRepositoryMockRecorder) UpdateClassroomName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomName", reflect.TypeOf((*MockClassroomRepository)(nil).UpdateClassroomName), arg0, arg1)
}

// UpdateClassroomRoom mocks base method.
func (m *MockClassroomRepository) UpdateClassroomRoom(arg0 context.Context, arg1 *postgresql.UpdateClassroomRoomParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomRoom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClassroomRoom indicates an expected call of UpdateClassroomRoom.
func (mr *MockClassroomRepositoryMockRecorder) UpdateClassroomRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomRoom", reflect.TypeOf((*MockClassroomRepository)(nil).UpdateClassroomRoom), arg0, arg1)
}

// UpdateClassroomSection mocks base method.
func (m *MockClassroomRepository) UpdateClassroomSection(arg0 context.Context, arg1 *postgresql.UpdateClassroomSectionParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomSection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClassroomSection indicates an expected call of UpdateClassroomSection.
func (mr *MockClassroomRepositoryMockRecorder) UpdateClassroomSection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomSection", reflect.TypeOf((*MockClassroomRepository)(nil).UpdateClassroomSection), arg0, arg1)
}

// UpdateClassroomSubject mocks base method.
func (m *MockClassroomRepository) UpdateClassroomSubject(arg0 context.Context, arg1 *postgresql.UpdateClassroomSubjectParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClassroomSubject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClassroomSubject indicates an expected call of UpdateClassroomSubject.
func (mr *MockClassroomRepositoryMockRecorder) UpdateClassroomSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClassroomSubject", reflect.TypeOf((*MockClassroomRepository)(nil).UpdateClassroomSubject), arg0, arg1)
}
