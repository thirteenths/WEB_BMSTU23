package mock_service

import (
	"github.com/golang/mock/gomock"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"reflect"
)

// MockIEvent is a mock of IEvent interface.
type MockIEvent struct {
	ctrl     *gomock.Controller
	recorder *MockIEventMockRecorder
}

// MockIEventMockRecorder is the mock recorder for MockIEvent.
type MockIEventMockRecorder struct {
	mock *MockIEvent
}

// NewMockIEvent creates a new mock instance.
func NewMockIEvent(ctrl *gomock.Controller) *MockIEvent {
	mock := &MockIEvent{ctrl: ctrl}
	mock.recorder = &MockIEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEvent) EXPECT() *MockIEventMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIEvent) Add(arg0 modelUI.Event) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockIEventMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIEvent)(nil).Add), arg0)
}

// Delete mocks base method.
func (m *MockIEvent) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIEventMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIEvent)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockIEvent) Get(arg0 int) (modelUI.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(modelUI.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIEventMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIEvent)(nil).Get), arg0)
}

// GetAll mocks base method.
func (m *MockIEvent) GetAll() ([]modelUI.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]modelUI.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIEventMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIEvent)(nil).GetAll))
}

// Update mocks base method.
func (m *MockIEvent) Update(arg0 int, arg1 modelUI.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIEventMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIEvent)(nil).Update), arg0, arg1)
}
