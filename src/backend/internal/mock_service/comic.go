package mock_service

import (
	"github.com/golang/mock/gomock"
	"github.com/thirteenths/WEB_BMSTU23/backend/internal/modelUI"
	"reflect"
)

// MockIComic is a mock of IComic interface.
type MockIComic struct {
	ctrl     *gomock.Controller
	recorder *MockIComicMockRecorder
}

// MockIComicMockRecorder is the mock recorder for MockIComic.
type MockIComicMockRecorder struct {
	mock *MockIComic
}

// NewMockIComic creates a new mock instance.
func NewMockIComic(ctrl *gomock.Controller) *MockIComic {
	mock := &MockIComic{ctrl: ctrl}
	mock.recorder = &MockIComicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIComic) EXPECT() *MockIComicMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIComic) Add(arg0 modelUI.Comic) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockIComicMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIComic)(nil).Add), arg0)
}

// Delete mocks base method.
func (m *MockIComic) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIComicMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIComic)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockIComic) Get(arg0 int) (modelUI.Comic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(modelUI.Comic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIComicMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIComic)(nil).Get), arg0)
}

// GetAll mocks base method.
func (m *MockIComic) GetAll() ([]modelUI.Comic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]modelUI.Comic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIComicMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIComic)(nil).GetAll))
}

// Update mocks base method.
func (m *MockIComic) Update(arg0 int, arg1 modelUI.Comic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIComicMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIComic)(nil).Update), arg0, arg1)
}
