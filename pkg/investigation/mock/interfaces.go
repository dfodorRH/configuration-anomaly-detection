// Code generated by MockGen. DO NOT EDIT.
// Source: investigation.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Resolved mocks base method.
func (m *MockService) Resolved() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolved")
	ret0, _ := ret[0].(error)
	return ret0
}

// Resolved indicates an expected call of Resolved.
func (mr *MockServiceMockRecorder) Resolved() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolved", reflect.TypeOf((*MockService)(nil).Resolved))
}

// Triggered mocks base method.
func (m *MockService) Triggered() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Triggered")
	ret0, _ := ret[0].(error)
	return ret0
}

// Triggered indicates an expected call of Triggered.
func (mr *MockServiceMockRecorder) Triggered() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Triggered", reflect.TypeOf((*MockService)(nil).Triggered))
}
