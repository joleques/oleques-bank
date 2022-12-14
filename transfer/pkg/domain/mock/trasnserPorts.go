// Code generated by MockGen. DO NOT EDIT.
// Source: src/domain/ports.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/joleques/oleques-bank/transfer/src/domain"
)

// MockAccountService is a mock of AccountService interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// GetBalance mocks base method.
func (m *MockAccountService) GetBalance(id string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockAccountServiceMockRecorder) GetBalance(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockAccountService)(nil).GetBalance), id)
}

// UpdateAccount mocks base method.
func (m *MockAccountService) UpdateAccount(account domain.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountServiceMockRecorder) UpdateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountService)(nil).UpdateAccount), account)
}

// MockTransferRepository is a mock of TransferRepository interface.
type MockTransferRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransferRepositoryMockRecorder
}

// MockTransferRepositoryMockRecorder is the mock recorder for MockTransferRepository.
type MockTransferRepositoryMockRecorder struct {
	mock *MockTransferRepository
}

// NewMockTransferRepository creates a new mock instance.
func NewMockTransferRepository(ctrl *gomock.Controller) *MockTransferRepository {
	mock := &MockTransferRepository{ctrl: ctrl}
	mock.recorder = &MockTransferRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferRepository) EXPECT() *MockTransferRepositoryMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockTransferRepository) List(id string) []*domain.Transfer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", id)
	ret0, _ := ret[0].([]*domain.Transfer)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockTransferRepositoryMockRecorder) List(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTransferRepository)(nil).List), id)
}

// Save mocks base method.
func (m *MockTransferRepository) Save(transfer *domain.Transfer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", transfer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockTransferRepositoryMockRecorder) Save(transfer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTransferRepository)(nil).Save), transfer)
}
