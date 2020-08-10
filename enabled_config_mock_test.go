// Code generated by MockGen. DO NOT EDIT.
// Source: enabled_config.go

// Package orlop_test is a generated GoMock package.
package orlop_test

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHasEnabled is a mock of HasEnabled interface
type MockHasEnabled struct {
	ctrl     *gomock.Controller
	recorder *MockHasEnabledMockRecorder
}

// MockHasEnabledMockRecorder is the mock recorder for MockHasEnabled
type MockHasEnabledMockRecorder struct {
	mock *MockHasEnabled
}

// NewMockHasEnabled creates a new mock instance
func NewMockHasEnabled(ctrl *gomock.Controller) *MockHasEnabled {
	mock := &MockHasEnabled{ctrl: ctrl}
	mock.recorder = &MockHasEnabledMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHasEnabled) EXPECT() *MockHasEnabledMockRecorder {
	return m.recorder
}

// GetEnabled mocks base method
func (m *MockHasEnabled) GetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetEnabled indicates an expected call of GetEnabled
func (mr *MockHasEnabledMockRecorder) GetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnabled", reflect.TypeOf((*MockHasEnabled)(nil).GetEnabled))
}