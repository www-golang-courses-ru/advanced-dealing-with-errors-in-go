// Code generated by MockGen. DO NOT EDIT.
// Source: transport.go

// Package commandexecutormocks is a generated GoMock package.
package commandexecutormocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	commandexecutor "github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/tasks/07-working-with-errors-in-concurrency/command-executor"
)

// MockITransport is a mock of ITransport interface.
type MockITransport struct {
	ctrl     *gomock.Controller
	recorder *MockITransportMockRecorder
}

// MockITransportMockRecorder is the mock recorder for MockITransport.
type MockITransportMockRecorder struct {
	mock *MockITransport
}

// NewMockITransport creates a new mock instance.
func NewMockITransport(ctrl *gomock.Controller) *MockITransport {
	mock := &MockITransport{ctrl: ctrl}
	mock.recorder = &MockITransportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransport) EXPECT() *MockITransportMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockITransport) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockITransportMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockITransport)(nil).Context))
}

// Recv mocks base method.
func (m *MockITransport) Recv() (*commandexecutor.ProtoCommand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*commandexecutor.ProtoCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockITransportMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockITransport)(nil).Recv))
}

// Send mocks base method.
func (m *MockITransport) Send(c *commandexecutor.ProtoCommandResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockITransportMockRecorder) Send(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockITransport)(nil).Send), c)
}
