// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/khafizullokh02/warehouse_management/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateAgreementForm mocks base method.
func (m *MockStore) CreateAgreementForm(arg0 context.Context, arg1 db.CreateAgreementFormParams) (db.AgreementForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAgreementForm", arg0, arg1)
	ret0, _ := ret[0].(db.AgreementForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAgreementForm indicates an expected call of CreateAgreementForm.
func (mr *MockStoreMockRecorder) CreateAgreementForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAgreementForm", reflect.TypeOf((*MockStore)(nil).CreateAgreementForm), arg0, arg1)
}

// CreateEntryGroup mocks base method.
func (m *MockStore) CreateEntryGroup(arg0 context.Context, arg1 db.CreateEntryGroupParams) (db.EntryGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntryGroup", arg0, arg1)
	ret0, _ := ret[0].(db.EntryGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntryGroup indicates an expected call of CreateEntryGroup.
func (mr *MockStoreMockRecorder) CreateEntryGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntryGroup", reflect.TypeOf((*MockStore)(nil).CreateEntryGroup), arg0, arg1)
}

// CreateEntryItem mocks base method.
func (m *MockStore) CreateEntryItem(arg0 context.Context, arg1 db.CreateEntryItemParams) (db.EntryItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntryItem", arg0, arg1)
	ret0, _ := ret[0].(db.EntryItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntryItem indicates an expected call of CreateEntryItem.
func (mr *MockStoreMockRecorder) CreateEntryItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntryItem", reflect.TypeOf((*MockStore)(nil).CreateEntryItem), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// DeleteAgreementForm mocks base method.
func (m *MockStore) DeleteAgreementForm(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAgreementForm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgreementForm indicates an expected call of DeleteAgreementForm.
func (mr *MockStoreMockRecorder) DeleteAgreementForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgreementForm", reflect.TypeOf((*MockStore)(nil).DeleteAgreementForm), arg0, arg1)
}

// DeleteEntryGroup mocks base method.
func (m *MockStore) DeleteEntryGroup(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntryGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntryGroup indicates an expected call of DeleteEntryGroup.
func (mr *MockStoreMockRecorder) DeleteEntryGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntryGroup", reflect.TypeOf((*MockStore)(nil).DeleteEntryGroup), arg0, arg1)
}

// DeleteEntryItem mocks base method.
func (m *MockStore) DeleteEntryItem(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntryItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntryItem indicates an expected call of DeleteEntryItem.
func (mr *MockStoreMockRecorder) DeleteEntryItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntryItem", reflect.TypeOf((*MockStore)(nil).DeleteEntryItem), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockStore) DeleteProduct(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockStoreMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStore)(nil).DeleteProduct), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(arg0 context.Context, arg1 int32) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), arg0, arg1)
}

// GetAgreementForm mocks base method.
func (m *MockStore) GetAgreementForm(arg0 context.Context, arg1 int32) (db.AgreementForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgreementForm", arg0, arg1)
	ret0, _ := ret[0].(db.AgreementForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgreementForm indicates an expected call of GetAgreementForm.
func (mr *MockStoreMockRecorder) GetAgreementForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgreementForm", reflect.TypeOf((*MockStore)(nil).GetAgreementForm), arg0, arg1)
}

// GetEntryGroup mocks base method.
func (m *MockStore) GetEntryGroup(arg0 context.Context, arg1 int32) (db.EntryGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntryGroup", arg0, arg1)
	ret0, _ := ret[0].(db.EntryGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntryGroup indicates an expected call of GetEntryGroup.
func (mr *MockStoreMockRecorder) GetEntryGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntryGroup", reflect.TypeOf((*MockStore)(nil).GetEntryGroup), arg0, arg1)
}

// GetEntryItem mocks base method.
func (m *MockStore) GetEntryItem(arg0 context.Context, arg1 int32) (db.EntryItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntryItem", arg0, arg1)
	ret0, _ := ret[0].(db.EntryItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntryItem indicates an expected call of GetEntryItem.
func (mr *MockStoreMockRecorder) GetEntryItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntryItem", reflect.TypeOf((*MockStore)(nil).GetEntryItem), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockStore) GetProduct(arg0 context.Context, arg1 int32) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStore)(nil).GetProduct), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 int32) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int32) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByParams mocks base method.
func (m *MockStore) GetUserByParams(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByParams", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByParams indicates an expected call of GetUserByParams.
func (mr *MockStoreMockRecorder) GetUserByParams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByParams", reflect.TypeOf((*MockStore)(nil).GetUserByParams), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// ListAgreementForms mocks base method.
func (m *MockStore) ListAgreementForms(arg0 context.Context, arg1 db.ListAgreementFormsParams) ([]db.AgreementForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAgreementForms", arg0, arg1)
	ret0, _ := ret[0].([]db.AgreementForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAgreementForms indicates an expected call of ListAgreementForms.
func (mr *MockStoreMockRecorder) ListAgreementForms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAgreementForms", reflect.TypeOf((*MockStore)(nil).ListAgreementForms), arg0, arg1)
}

// ListEntryGroups mocks base method.
func (m *MockStore) ListEntryGroups(arg0 context.Context, arg1 db.ListEntryGroupsParams) ([]db.EntryGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntryGroups", arg0, arg1)
	ret0, _ := ret[0].([]db.EntryGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntryGroups indicates an expected call of ListEntryGroups.
func (mr *MockStoreMockRecorder) ListEntryGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntryGroups", reflect.TypeOf((*MockStore)(nil).ListEntryGroups), arg0, arg1)
}

// ListEntryItems mocks base method.
func (m *MockStore) ListEntryItems(arg0 context.Context, arg1 db.ListEntryItemsParams) ([]db.EntryItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntryItems", arg0, arg1)
	ret0, _ := ret[0].([]db.EntryItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntryItems indicates an expected call of ListEntryItems.
func (mr *MockStoreMockRecorder) ListEntryItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntryItems", reflect.TypeOf((*MockStore)(nil).ListEntryItems), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockStore) ListProducts(arg0 context.Context, arg1 db.ListProductsParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreMockRecorder) ListProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStore)(nil).ListProducts), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(arg0 context.Context, arg1 db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdateAgreementForm mocks base method.
func (m *MockStore) UpdateAgreementForm(arg0 context.Context, arg1 db.UpdateAgreementFormParams) (db.AgreementForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAgreementForm", arg0, arg1)
	ret0, _ := ret[0].(db.AgreementForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAgreementForm indicates an expected call of UpdateAgreementForm.
func (mr *MockStoreMockRecorder) UpdateAgreementForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgreementForm", reflect.TypeOf((*MockStore)(nil).UpdateAgreementForm), arg0, arg1)
}

// UpdateEntryGroup mocks base method.
func (m *MockStore) UpdateEntryGroup(arg0 context.Context, arg1 db.UpdateEntryGroupParams) (db.EntryGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntryGroup", arg0, arg1)
	ret0, _ := ret[0].(db.EntryGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntryGroup indicates an expected call of UpdateEntryGroup.
func (mr *MockStoreMockRecorder) UpdateEntryGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntryGroup", reflect.TypeOf((*MockStore)(nil).UpdateEntryGroup), arg0, arg1)
}

// UpdateEntryItem mocks base method.
func (m *MockStore) UpdateEntryItem(arg0 context.Context, arg1 db.UpdateEntryItemParams) (db.EntryItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntryItem", arg0, arg1)
	ret0, _ := ret[0].(db.EntryItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntryItem indicates an expected call of UpdateEntryItem.
func (mr *MockStoreMockRecorder) UpdateEntryItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntryItem", reflect.TypeOf((*MockStore)(nil).UpdateEntryItem), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockStore) UpdateProduct(arg0 context.Context, arg1 db.UpdateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockStoreMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStore)(nil).UpdateProduct), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}