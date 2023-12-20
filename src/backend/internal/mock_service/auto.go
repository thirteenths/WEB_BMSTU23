package mock_service

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockIAuthorization is a mock of IAuthorization interface.
type MockIAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthorizationMockRecorder
}

// MockIAuthorizationMockRecorder is the mock recorder for MockIAuthorization.
type MockIAuthorizationMockRecorder struct {
	mock *MockIAuthorization
}

// NewMockIAuthorization creates a new mock instance.
func NewMockIAuthorization(ctrl *gomock.Controller) *MockIAuthorization {
	mock := &MockIAuthorization{ctrl: ctrl}
	mock.recorder = &MockIAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthorization) EXPECT() *MockIAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIAuthorization) CreateUser(arg0 modelDB.Person) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIAuthorizationMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIAuthorization)(nil).CreateUser), arg0)
}

// GenerationToken mocks base method.
func (m *MockIAuthorization) GenerationToken(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerationToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerationToken indicates an expected call of GenerationToken.
func (mr *MockIAuthorizationMockRecorder) GenerationToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerationToken", reflect.TypeOf((*MockIAuthorization)(nil).GenerationToken), arg0, arg1)
}

// ParseToken mocks base method.
func (m *MockIAuthorization) ParseToken(arg0 string) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockIAuthorizationMockRecorder) ParseToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockIAuthorization)(nil).ParseToken), arg0)
}
