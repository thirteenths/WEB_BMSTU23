package mock_storage

import (
	"github.com/golang/mock/gomock"
	"reflect"

	"github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

// MockRequestMaker is a mock of RequestMaker interface.
type MockRequestMaker struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMakerMockRecorder
}

// MockRequestMakerMockRecorder is the mock recorder for MockRequestMaker.
type MockRequestMakerMockRecorder struct {
	mock *MockRequestMaker
}

// NewMockRequestMaker creates a new mock instance.
func NewMockRequestMaker(ctrl *gomock.Controller) *MockRequestMaker {
	mock := &MockRequestMaker{ctrl: ctrl}
	mock.recorder = &MockRequestMakerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestMaker) EXPECT() *MockRequestMakerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRequestMaker) Create() storage.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(storage.Query)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRequestMakerMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRequestMaker)(nil).Create))
}

// DeleteById mocks base method.
func (m *MockRequestMaker) DeleteById(id int) storage.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", id)
	ret0, _ := ret[0].(storage.Query)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockRequestMakerMockRecorder) DeleteById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockRequestMaker)(nil).DeleteById), id)
}

// Drop mocks base method.
func (m *MockRequestMaker) Drop() storage.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop")
	ret0, _ := ret[0].(storage.Query)
	return ret0
}

// Drop indicates an expected call of Drop.
func (mr *MockRequestMakerMockRecorder) Drop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockRequestMaker)(nil).Drop))
}

// Insert mocks base method.
func (m *MockRequestMaker) Insert(model any) storage.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", model)
	ret0, _ := ret[0].(storage.Query)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRequestMakerMockRecorder) Insert(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRequestMaker)(nil).Insert), model)
}

// SelectByField mocks base method.
func (m *MockRequestMaker) SelectByField(fieldName string, fieldValue any) storage.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByField", fieldName, fieldValue)
	ret0, _ := ret[0].(storage.Query)
	return ret0
}

// SelectByField indicates an expected call of SelectByField.
func (mr *MockRequestMakerMockRecorder) SelectByField(fieldName, fieldValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByField", reflect.TypeOf((*MockRequestMaker)(nil).SelectByField), fieldName, fieldValue)
}

// SelectById mocks base method.
func (m *MockRequestMaker) SelectById(id int) storage.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", id)
	ret0, _ := ret[0].(storage.Query)
	return ret0
}

// SelectById indicates an expected call of SelectById.
func (mr *MockRequestMakerMockRecorder) SelectById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockRequestMaker)(nil).SelectById), id)
}

// UpdateById mocks base method.
func (m *MockRequestMaker) UpdateById(id int) storage.Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", id)
	ret0, _ := ret[0].(storage.Query)
	return ret0
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockRequestMakerMockRecorder) UpdateById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockRequestMaker)(nil).UpdateById), id)
}
