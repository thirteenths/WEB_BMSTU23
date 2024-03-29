// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/storage/repository.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/thirteenths/WEB_BMSTU23/backend/pkg/storage"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create))
}

// DeleteById mocks base method.
func (m *MockRepository) DeleteById(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockRepositoryMockRecorder) DeleteById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockRepository)(nil).DeleteById), id)
}

// Drop mocks base method.
func (m *MockRepository) Drop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop.
func (mr *MockRepositoryMockRecorder) Drop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockRepository)(nil).Drop))
}

// Insert mocks base method.
func (m *MockRepository) Insert(model any) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", model)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockRepositoryMockRecorder) Insert(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepository)(nil).Insert), model)
}

// SelectAll mocks base method.
func (m *MockRepository) SelectAll() ([]storage.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll")
	ret0, _ := ret[0].([]storage.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll.
func (mr *MockRepositoryMockRecorder) SelectAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockRepository)(nil).SelectAll))
}

// SelectByField mocks base method.
func (m *MockRepository) SelectByField(fieldName string, fieldValue storage.FieldValue) ([]storage.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByField", fieldName, fieldValue)
	ret0, _ := ret[0].([]storage.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByField indicates an expected call of SelectByField.
func (mr *MockRepositoryMockRecorder) SelectByField(fieldName, fieldValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByField", reflect.TypeOf((*MockRepository)(nil).SelectByField), fieldName, fieldValue)
}

// SelectById mocks base method.
func (m *MockRepository) SelectById(id int) (storage.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", id)
	ret0, _ := ret[0].(storage.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockRepositoryMockRecorder) SelectById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockRepository)(nil).SelectById), id)
}

// UpdateById mocks base method.
func (m *MockRepository) UpdateById(id int, model storage.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", id, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockRepositoryMockRecorder) UpdateById(id, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockRepository)(nil).UpdateById), id, model)
}





// MockDifQManager is a mock of DifQManager interface.
type MockDifQManager struct {
	ctrl     *gomock.Controller
	recorder *MockDifQManagerMockRecorder
}

// MockDifQManagerMockRecorder is the mock recorder for MockDifQManager.
type MockDifQManagerMockRecorder struct {
	mock *MockDifQManager
}

// NewMockDifQManager creates a new mock instance.
func NewMockDifQManager(ctrl *gomock.Controller) *MockDifQManager {
	mock := &MockDifQManager{ctrl: ctrl}
	mock.recorder = &MockDifQManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDifQManager) EXPECT() *MockDifQManagerMockRecorder {
	return m.recorder
}

// DeleteDataComic mocks base method.
func (m *MockDifQManager) DeleteDataComic(idPerson, idSchedule int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataComic", idPerson, idSchedule)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataComic indicates an expected call of DeleteDataComic.
func (mr *MockDifQManagerMockRecorder) DeleteDataComic(idPerson, idSchedule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataComic", reflect.TypeOf((*MockDifQManager)(nil).DeleteDataComic), idPerson, idSchedule)
}

// DeleteDataViewer mocks base method.
func (m *MockDifQManager) DeleteDataViewer(idPerson, idSchedule int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataViewer", idPerson, idSchedule)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDataViewer indicates an expected call of DeleteDataViewer.
func (mr *MockDifQManagerMockRecorder) DeleteDataViewer(idPerson, idSchedule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataViewer", reflect.TypeOf((*MockDifQManager)(nil).DeleteDataViewer), idPerson, idSchedule)
}

// InsertDataComic mocks base method.
func (m *MockDifQManager) InsertDataComic(idPerson, idSchedule int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertDataComic", idPerson, idSchedule)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertDataComic indicates an expected call of InsertDataComic.
func (mr *MockDifQManagerMockRecorder) InsertDataComic(idPerson, idSchedule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertDataComic", reflect.TypeOf((*MockDifQManager)(nil).InsertDataComic), idPerson, idSchedule)
}

// InsertDataViewer mocks base method.
func (m *MockDifQManager) InsertDataViewer(idPerson, idSchedule int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertDataViewer", idPerson, idSchedule)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertDataViewer indicates an expected call of InsertDataViewer.
func (mr *MockDifQManagerMockRecorder) InsertDataViewer(idPerson, idSchedule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertDataViewer", reflect.TypeOf((*MockDifQManager)(nil).InsertDataViewer), idPerson, idSchedule)
}

// MockIStorage is a mock of IStorage interface.
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage.
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance.
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// QueryManager mocks base method.
func (m *MockIStorage) QueryManager() storage.DifQManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryManager")
	ret0, _ := ret[0].(storage.DifQManager)
	return ret0
}

// QueryManager indicates an expected call of QueryManager.
func (mr *MockIStorageMockRecorder) QueryManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryManager", reflect.TypeOf((*MockIStorage)(nil).QueryManager))
}

// Repositories mocks base method.
func (m *MockIStorage) Repositories() map[storage.TableName]storage.Repository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Repositories")
	ret0, _ := ret[0].(map[storage.TableName]storage.Repository)
	return ret0
}

// Repositories indicates an expected call of Repositories.
func (mr *MockIStorageMockRecorder) Repositories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repositories", reflect.TypeOf((*MockIStorage)(nil).Repositories))
}
