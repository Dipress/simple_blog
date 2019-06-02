// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package update is a generated GoMock package.
package update

import (
	context "context"
	post "github.com/dipress/blog/internal/post"
	user "github.com/dipress/blog/internal/user"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAbillity is a mock of Abillity interface
type MockAbillity struct {
	ctrl     *gomock.Controller
	recorder *MockAbillityMockRecorder
}

// MockAbillityMockRecorder is the mock recorder for MockAbillity
type MockAbillityMockRecorder struct {
	mock *MockAbillity
}

// NewMockAbillity creates a new mock instance
func NewMockAbillity(ctrl *gomock.Controller) *MockAbillity {
	mock := &MockAbillity{ctrl: ctrl}
	mock.recorder = &MockAbillityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAbillity) EXPECT() *MockAbillityMockRecorder {
	return m.recorder
}

// CanUpdate mocks base method
func (m *MockAbillity) CanUpdate(userID int, post *post.Post) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanUpdate", userID, post)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanUpdate indicates an expected call of CanUpdate
func (mr *MockAbillityMockRecorder) CanUpdate(userID, post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanUpdate", reflect.TypeOf((*MockAbillity)(nil).CanUpdate), userID, post)
}

// MockValidater is a mock of Validater interface
type MockValidater struct {
	ctrl     *gomock.Controller
	recorder *MockValidaterMockRecorder
}

// MockValidaterMockRecorder is the mock recorder for MockValidater
type MockValidaterMockRecorder struct {
	mock *MockValidater
}

// NewMockValidater creates a new mock instance
func NewMockValidater(ctrl *gomock.Controller) *MockValidater {
	mock := &MockValidater{ctrl: ctrl}
	mock.recorder = &MockValidaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidater) EXPECT() *MockValidaterMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockValidater) Validate(arg0 context.Context, arg1 *Form) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockValidaterMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidater)(nil).Validate), arg0, arg1)
}

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// FindByUsername mocks base method
func (m *MockRepository) FindByUsername(ctx context.Context, username string, u *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", ctx, username, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindByUsername indicates an expected call of FindByUsername
func (mr *MockRepositoryMockRecorder) FindByUsername(ctx, username, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockRepository)(nil).FindByUsername), ctx, username, u)
}

// FindPost mocks base method
func (m *MockRepository) FindPost(ctx context.Context, id int) (*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPost", ctx, id)
	ret0, _ := ret[0].(*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPost indicates an expected call of FindPost
func (mr *MockRepositoryMockRecorder) FindPost(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPost", reflect.TypeOf((*MockRepository)(nil).FindPost), ctx, id)
}

// UpdatePost mocks base method
func (m *MockRepository) UpdatePost(ctx context.Context, id int, p *post.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", ctx, id, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost
func (mr *MockRepositoryMockRecorder) UpdatePost(ctx, id, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockRepository)(nil).UpdatePost), ctx, id, p)
}
