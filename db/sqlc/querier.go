// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateAgreementForm(ctx context.Context, arg CreateAgreementFormParams) (AgreementForm, error)
	CreateEntryGroup(ctx context.Context, arg CreateEntryGroupParams) (EntryGroup, error)
	CreateEntryItem(ctx context.Context, arg CreateEntryItemParams) (EntryItem, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int32) error
	DeleteAgreementForm(ctx context.Context, id int32) error
	DeleteEntryGroup(ctx context.Context, id int32) error
	DeleteEntryItem(ctx context.Context, id int32) error
	DeleteProduct(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, id int32) (Account, error)
	GetAgreementForm(ctx context.Context, id int32) (AgreementForm, error)
	GetEntryGroup(ctx context.Context, id int32) (EntryGroup, error)
	GetEntryItem(ctx context.Context, id int32) (EntryItem, error)
	GetProduct(ctx context.Context, id int32) (Product, error)
	GetUser(ctx context.Context, id int32) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListAgreementForms(ctx context.Context, arg ListAgreementFormsParams) ([]AgreementForm, error)
	ListEntryGroups(ctx context.Context, arg ListEntryGroupsParams) ([]EntryGroup, error)
	ListEntryItems(ctx context.Context, arg ListEntryItemsParams) ([]EntryItem, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateAgreementForm(ctx context.Context, arg UpdateAgreementFormParams) (AgreementForm, error)
	UpdateEntryGroup(ctx context.Context, arg UpdateEntryGroupParams) (EntryGroup, error)
	UpdateEntryItem(ctx context.Context, arg UpdateEntryItemParams) (EntryItem, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
