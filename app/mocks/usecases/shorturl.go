// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/usecases/shorturl.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	shorturl "rabbit-test/app/usecases/shorturl"
	reflect "reflect"
)

// MockShortURL is a mock of ShortURL interface
type MockShortURL struct {
	ctrl     *gomock.Controller
	recorder *MockShortURLMockRecorder
}

// MockShortURLMockRecorder is the mock recorder for MockShortURL
type MockShortURLMockRecorder struct {
	mock *MockShortURL
}

// NewMockShortURL creates a new mock instance
func NewMockShortURL(ctrl *gomock.Controller) *MockShortURL {
	mock := &MockShortURL{ctrl: ctrl}
	mock.recorder = &MockShortURLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockShortURL) EXPECT() *MockShortURLMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockShortURL) Create(ctx context.Context, fullUrl string, expiry *int) (*shorturl.CreateShortURLResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, fullUrl, expiry)
	ret0, _ := ret[0].(*shorturl.CreateShortURLResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockShortURLMockRecorder) Create(ctx, fullUrl, expiry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockShortURL)(nil).Create), ctx, fullUrl, expiry)
}

// Get mocks base method
func (m *MockShortURL) Get(ctx context.Context, shortCode string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, shortCode)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockShortURLMockRecorder) Get(ctx, shortCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShortURL)(nil).Get), ctx, shortCode)
}
