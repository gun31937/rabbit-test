// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/repositories/redis.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRedis is a mock of Redis interface
type MockRedis struct {
	ctrl     *gomock.Controller
	recorder *MockRedisMockRecorder
}

// MockRedisMockRecorder is the mock recorder for MockRedis
type MockRedisMockRecorder struct {
	mock *MockRedis
}

// NewMockRedis creates a new mock instance
func NewMockRedis(ctrl *gomock.Controller) *MockRedis {
	mock := &MockRedis{ctrl: ctrl}
	mock.recorder = &MockRedisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedis) EXPECT() *MockRedisMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockRedis) Set(ctx context.Context, key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockRedisMockRecorder) Set(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedis)(nil).Set), ctx, key, value)
}

// Get mocks base method
func (m *MockRedis) Get(ctx context.Context, key string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRedisMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedis)(nil).Get), ctx, key)
}
