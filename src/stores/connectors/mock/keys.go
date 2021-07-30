// Code generated by MockGen. DO NOT EDIT.
// Source: keys.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	entities "github.com/consensys/quorum-key-manager/src/stores/store/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockKeysConnector is a mock of KeysConnector interface
type MockKeysConnector struct {
	ctrl     *gomock.Controller
	recorder *MockKeysConnectorMockRecorder
}

// MockKeysConnectorMockRecorder is the mock recorder for MockKeysConnector
type MockKeysConnectorMockRecorder struct {
	mock *MockKeysConnector
}

// NewMockKeysConnector creates a new mock instance
func NewMockKeysConnector(ctrl *gomock.Controller) *MockKeysConnector {
	mock := &MockKeysConnector{ctrl: ctrl}
	mock.recorder = &MockKeysConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeysConnector) EXPECT() *MockKeysConnectorMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockKeysConnector) Create(ctx context.Context, id string, alg *entities.Algorithm, attr *entities.Attributes) (*entities.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, id, alg, attr)
	ret0, _ := ret[0].(*entities.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockKeysConnectorMockRecorder) Create(ctx, id, alg, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockKeysConnector)(nil).Create), ctx, id, alg, attr)
}

// Import mocks base method
func (m *MockKeysConnector) Import(ctx context.Context, id string, privKey []byte, alg *entities.Algorithm, attr *entities.Attributes) (*entities.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", ctx, id, privKey, alg, attr)
	ret0, _ := ret[0].(*entities.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Import indicates an expected call of Import
func (mr *MockKeysConnectorMockRecorder) Import(ctx, id, privKey, alg, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockKeysConnector)(nil).Import), ctx, id, privKey, alg, attr)
}

// Get mocks base method
func (m *MockKeysConnector) Get(ctx context.Context, id string) (*entities.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*entities.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockKeysConnectorMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKeysConnector)(nil).Get), ctx, id)
}

// List mocks base method
func (m *MockKeysConnector) List(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockKeysConnectorMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKeysConnector)(nil).List), ctx)
}

// Update mocks base method
func (m *MockKeysConnector) Update(ctx context.Context, id string, attr *entities.Attributes) (*entities.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, attr)
	ret0, _ := ret[0].(*entities.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockKeysConnectorMockRecorder) Update(ctx, id, attr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockKeysConnector)(nil).Update), ctx, id, attr)
}

// Delete mocks base method
func (m *MockKeysConnector) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockKeysConnectorMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKeysConnector)(nil).Delete), ctx, id)
}

// GetDeleted mocks base method
func (m *MockKeysConnector) GetDeleted(ctx context.Context, id string) (*entities.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeleted", ctx, id)
	ret0, _ := ret[0].(*entities.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeleted indicates an expected call of GetDeleted
func (mr *MockKeysConnectorMockRecorder) GetDeleted(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeleted", reflect.TypeOf((*MockKeysConnector)(nil).GetDeleted), ctx, id)
}

// ListDeleted mocks base method
func (m *MockKeysConnector) ListDeleted(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeleted", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeleted indicates an expected call of ListDeleted
func (mr *MockKeysConnectorMockRecorder) ListDeleted(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeleted", reflect.TypeOf((*MockKeysConnector)(nil).ListDeleted), ctx)
}

// Restore mocks base method
func (m *MockKeysConnector) Restore(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore
func (mr *MockKeysConnectorMockRecorder) Restore(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockKeysConnector)(nil).Restore), ctx, id)
}

// Destroy mocks base method
func (m *MockKeysConnector) Destroy(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockKeysConnectorMockRecorder) Destroy(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockKeysConnector)(nil).Destroy), ctx, id)
}

// Sign mocks base method
func (m *MockKeysConnector) Sign(ctx context.Context, id string, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", ctx, id, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockKeysConnectorMockRecorder) Sign(ctx, id, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockKeysConnector)(nil).Sign), ctx, id, data)
}

// Verify mocks base method
func (m *MockKeysConnector) Verify(ctx context.Context, pubKey, data, sig []byte, algo *entities.Algorithm) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, pubKey, data, sig, algo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MockKeysConnectorMockRecorder) Verify(ctx, pubKey, data, sig, algo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockKeysConnector)(nil).Verify), ctx, pubKey, data, sig, algo)
}

// Encrypt mocks base method
func (m *MockKeysConnector) Encrypt(ctx context.Context, id string, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", ctx, id, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt
func (mr *MockKeysConnectorMockRecorder) Encrypt(ctx, id, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockKeysConnector)(nil).Encrypt), ctx, id, data)
}

// Decrypt mocks base method
func (m *MockKeysConnector) Decrypt(ctx context.Context, id string, data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", ctx, id, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt
func (mr *MockKeysConnectorMockRecorder) Decrypt(ctx, id, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockKeysConnector)(nil).Decrypt), ctx, id, data)
}