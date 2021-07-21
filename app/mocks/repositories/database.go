// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/repositories/database.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	gomock "github.com/golang/mock/gomock"
	database "rabbit-test/app/repositories/database"
	reflect "reflect"
)

// MockDatabase is a mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CreateURL mocks base method
func (m *MockDatabase) CreateURL(request database.CreateShortURLRequest) (*uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateURL", request)
	ret0, _ := ret[0].(*uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateURL indicates an expected call of CreateURL
func (mr *MockDatabaseMockRecorder) CreateURL(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateURL", reflect.TypeOf((*MockDatabase)(nil).CreateURL), request)
}

// CountAllURL mocks base method
func (m *MockDatabase) CountAllURL() (*uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAllURL")
	ret0, _ := ret[0].(*uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAllURL indicates an expected call of CountAllURL
func (mr *MockDatabaseMockRecorder) CountAllURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAllURL", reflect.TypeOf((*MockDatabase)(nil).CountAllURL))
}
