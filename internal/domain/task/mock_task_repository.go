// Code generated by MockGen. DO NOT EDIT.
// Source: task_repository.go
//
// Generated by this command:
//
//	mockgen -package task -source task_repository.go -destination mock_task_repository.go
//

// Package task is a generated GoMock package.
package task

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockRepository) FindAll(ctx context.Context, filter Filter) ([]*Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, filter)
	ret0, _ := ret[0].([]*Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(ctx, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), ctx, filter)
}

// FindByStatus mocks base method.
func (m *MockRepository) FindByStatus(ctx context.Context, userID string, status Status) ([]*Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByStatus", ctx, userID, status)
	ret0, _ := ret[0].([]*Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByStatus indicates an expected call of FindByStatus.
func (mr *MockRepositoryMockRecorder) FindByStatus(ctx, userID, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByStatus", reflect.TypeOf((*MockRepository)(nil).FindByStatus), ctx, userID, status)
}

// Save mocks base method.
func (m *MockRepository) Save(ctx context.Context, userID string, task *Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, userID, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRepositoryMockRecorder) Save(ctx, userID, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRepository)(nil).Save), ctx, userID, task)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, userID string, task *Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, userID, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, userID, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, userID, task)
}

// UpdateStatus mocks base method.
func (m *MockRepository) UpdateStatus(ctx context.Context, userID, taskID string, status Status) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, userID, taskID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockRepositoryMockRecorder) UpdateStatus(ctx, userID, taskID, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockRepository)(nil).UpdateStatus), ctx, userID, taskID, status)
}
