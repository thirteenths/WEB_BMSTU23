package mock_service

import (
	"github.com/golang/mock/gomock"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"reflect"
)

// MockIPoster is a mock of IPoster interface.
type MockIPoster struct {
	ctrl     *gomock.Controller
	recorder *MockIPosterMockRecorder
}

// MockIPosterMockRecorder is the mock recorder for MockIPoster.
type MockIPosterMockRecorder struct {
	mock *MockIPoster
}

// NewMockIPoster creates a new mock instance.
func NewMockIPoster(ctrl *gomock.Controller) *MockIPoster {
	mock := &MockIPoster{ctrl: ctrl}
	mock.recorder = &MockIPosterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPoster) EXPECT() *MockIPosterMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIPoster) Add(arg0 modelUI.Poster) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockIPosterMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIPoster)(nil).Add), arg0)
}

// CheckIn mocks base method.
func (m *MockIPoster) CheckIn(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckIn indicates an expected call of CheckIn.
func (mr *MockIPosterMockRecorder) CheckIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIn", reflect.TypeOf((*MockIPoster)(nil).CheckIn), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIPoster) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIPosterMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIPoster)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockIPoster) Get(arg0 int) (modelUI.Poster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(modelUI.Poster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIPosterMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIPoster)(nil).Get), arg0)
}

// GetAll mocks base method.
func (m *MockIPoster) GetAll() ([]modelUI.Poster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]modelUI.Poster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIPosterMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIPoster)(nil).GetAll))
}

// Update mocks base method.
func (m *MockIPoster) Update(arg0 int, arg1 modelUI.Poster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIPosterMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPoster)(nil).Update), arg0, arg1)
}
