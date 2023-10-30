package mock_service

import (
	"github.com/golang/mock/gomock"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"reflect"
)

// MockIPlace is a mock of IPlace interface.
type MockIPlace struct {
	ctrl     *gomock.Controller
	recorder *MockIPlaceMockRecorder
}

// MockIPlaceMockRecorder is the mock recorder for MockIPlace.
type MockIPlaceMockRecorder struct {
	mock *MockIPlace
}

// NewMockIPlace creates a new mock instance.
func NewMockIPlace(ctrl *gomock.Controller) *MockIPlace {
	mock := &MockIPlace{ctrl: ctrl}
	mock.recorder = &MockIPlaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPlace) EXPECT() *MockIPlaceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIPlace) Add(arg0 modelUI.Place) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockIPlaceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIPlace)(nil).Add), arg0)
}

// Delete mocks base method.
func (m *MockIPlace) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIPlaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIPlace)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockIPlace) Get(arg0 int) (modelUI.Place, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(modelUI.Place)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIPlaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIPlace)(nil).Get), arg0)
}

// GetAll mocks base method.
func (m *MockIPlace) GetAll() ([]modelUI.Place, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]modelUI.Place)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIPlaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIPlace)(nil).GetAll))
}

// Update mocks base method.
func (m *MockIPlace) Update(arg0 int, arg1 modelUI.Place) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIPlaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPlace)(nil).Update), arg0, arg1)
}
