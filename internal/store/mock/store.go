// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dragon-huang0403/todo-go/internal/store (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -destination ./internal/store/mock/store.go github.com/dragon-huang0403/todo-go/internal/store Store
//

// Package mock_store is a generated GoMock package.
package mock_store

import (
	reflect "reflect"

	models "github.com/dragon-huang0403/todo-go/internal/models"
	store "github.com/dragon-huang0403/todo-go/internal/store"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
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

// CreateTask mocks base method.
func (m *MockStore) CreateTask(arg0 store.CreateTaskParams) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockStoreMockRecorder) CreateTask(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockStore)(nil).CreateTask), arg0)
}

// DeleteTask mocks base method.
func (m *MockStore) DeleteTask(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockStoreMockRecorder) DeleteTask(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockStore)(nil).DeleteTask), arg0)
}

// GetTask mocks base method.
func (m *MockStore) GetTask(arg0 uuid.UUID) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockStoreMockRecorder) GetTask(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockStore)(nil).GetTask), arg0)
}

// ListTasks mocks base method.
func (m *MockStore) ListTasks() ([]*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasks")
	ret0, _ := ret[0].([]*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks.
func (mr *MockStoreMockRecorder) ListTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockStore)(nil).ListTasks))
}

// UpdateTask mocks base method.
func (m *MockStore) UpdateTask(arg0 store.UpdateTaskParams) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", arg0)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockStoreMockRecorder) UpdateTask(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockStore)(nil).UpdateTask), arg0)
}
