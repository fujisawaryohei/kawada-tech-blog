// Code generated by MockGen. DO NOT EDIT.
// Source: ./web/auth/auth.go

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAuthenticator is a mock of IAuthenticator interface.
type MockIAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthenticatorMockRecorder
}

// MockIAuthenticatorMockRecorder is the mock recorder for MockIAuthenticator.
type MockIAuthenticatorMockRecorder struct {
	mock *MockIAuthenticator
}

// NewMockIAuthenticator creates a new mock instance.
func NewMockIAuthenticator(ctrl *gomock.Controller) *MockIAuthenticator {
	mock := &MockIAuthenticator{ctrl: ctrl}
	mock.recorder = &MockIAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthenticator) EXPECT() *MockIAuthenticatorMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockIAuthenticator) GenerateToken(email string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockIAuthenticatorMockRecorder) GenerateToken(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockIAuthenticator)(nil).GenerateToken), email)
}
