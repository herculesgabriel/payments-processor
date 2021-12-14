// Code generated by MockGen. DO NOT EDIT.
// Source: domain/adapter/broker/interface.go

// Package mock_broker is a generated GoMock package.
package mock_broker

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProducerInterface is a mock of ProducerInterface interface.
type MockProducerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProducerInterfaceMockRecorder
}

// MockProducerInterfaceMockRecorder is the mock recorder for MockProducerInterface.
type MockProducerInterfaceMockRecorder struct {
	mock *MockProducerInterface
}

// NewMockProducerInterface creates a new mock instance.
func NewMockProducerInterface(ctrl *gomock.Controller) *MockProducerInterface {
	mock := &MockProducerInterface{ctrl: ctrl}
	mock.recorder = &MockProducerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducerInterface) EXPECT() *MockProducerInterfaceMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockProducerInterface) Publish(message interface{}, key []byte, topic string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", message, key, topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockProducerInterfaceMockRecorder) Publish(message, key, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockProducerInterface)(nil).Publish), message, key, topic)
}
