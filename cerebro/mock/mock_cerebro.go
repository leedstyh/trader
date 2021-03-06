// Code generated by MockGen. DO NOT EDIT.
// Source: ./cerebro.go

// Package mock_cerebro is a generated GoMock package.
package mock_cerebro

import (
	store "github.com/gobenpark/trader/store"
	strategy "github.com/gobenpark/trader/strategy"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCerebroker is a mock of Cerebroker interface
type MockCerebroker struct {
	ctrl     *gomock.Controller
	recorder *MockCerebrokerMockRecorder
}

// MockCerebrokerMockRecorder is the mock recorder for MockCerebroker
type MockCerebrokerMockRecorder struct {
	mock *MockCerebroker
}

// NewMockCerebroker creates a new mock instance
func NewMockCerebroker(ctrl *gomock.Controller) *MockCerebroker {
	mock := &MockCerebroker{ctrl: ctrl}
	mock.recorder = &MockCerebrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCerebroker) EXPECT() *MockCerebrokerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockCerebroker) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockCerebrokerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCerebroker)(nil).Start))
}

// Stop mocks base method
func (m *MockCerebroker) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockCerebrokerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCerebroker)(nil).Stop))
}

// AddStore mocks base method
func (m *MockCerebroker) AddStore(arg0 store.Storer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddStore", arg0)
}

// AddStore indicates an expected call of AddStore
func (mr *MockCerebrokerMockRecorder) AddStore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStore", reflect.TypeOf((*MockCerebroker)(nil).AddStore), arg0)
}

// AddStrategy mocks base method
func (m *MockCerebroker) AddStrategy(arg0 strategy.Strategy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddStrategy", arg0)
}

// AddStrategy indicates an expected call of AddStrategy
func (mr *MockCerebrokerMockRecorder) AddStrategy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStrategy", reflect.TypeOf((*MockCerebroker)(nil).AddStrategy), arg0)
}

// Resample mocks base method
func (m *MockCerebroker) Resample() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Resample")
}

// Resample indicates an expected call of Resample
func (mr *MockCerebrokerMockRecorder) Resample() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resample", reflect.TypeOf((*MockCerebroker)(nil).Resample))
}
