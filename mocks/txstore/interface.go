// Code generated by MockGen. DO NOT EDIT.
// Source: txstore/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "eth-parser/common"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTxStore is a mock of TxStore interface.
type MockTxStore struct {
	ctrl     *gomock.Controller
	recorder *MockTxStoreMockRecorder
}

// MockTxStoreMockRecorder is the mock recorder for MockTxStore.
type MockTxStoreMockRecorder struct {
	mock *MockTxStore
}

// NewMockTxStore creates a new mock instance.
func NewMockTxStore(ctrl *gomock.Controller) *MockTxStore {
	mock := &MockTxStore{ctrl: ctrl}
	mock.recorder = &MockTxStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxStore) EXPECT() *MockTxStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTxStore) Add(address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockTxStoreMockRecorder) Add(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTxStore)(nil).Add), address)
}

// Append mocks base method.
func (m *MockTxStore) Append(address string, tx common.Transaction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Append", address, tx)
}

// Append indicates an expected call of Append.
func (mr *MockTxStoreMockRecorder) Append(address, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockTxStore)(nil).Append), address, tx)
}

// Keys mocks base method.
func (m *MockTxStore) Keys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockTxStoreMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockTxStore)(nil).Keys))
}

// List mocks base method.
func (m *MockTxStore) List(address string) ([]common.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", address)
	ret0, _ := ret[0].([]common.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTxStoreMockRecorder) List(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTxStore)(nil).List), address)
}
