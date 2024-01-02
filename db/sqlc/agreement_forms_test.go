package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAgreementForm(t *testing.T, actionType string) AgreementForm {
	accFrom := createRandomAccount(t)
	accTo := createRandomAccount(t)
	product := createRandomProduct(t)

	arg := CreateAgreementFormParams{
		FromAccount:            accFrom.ID,
		ToAccount:              accTo.ID,
		ProductIds:             []int32{product.ID},
		ActionTypeForAgreement: ActionType(actionType),
		WholesalePrice:         1.1,
		RetailPrice:            1.1,
	}

	agreementForm, err := testStore.CreateAgreementForm(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, agreementForm)

	require.Equal(t, arg.FromAccount, agreementForm.FromAccount)
	require.Equal(t, arg.ToAccount, agreementForm.ToAccount)
	require.Equal(t, arg.ProductIds, agreementForm.ProductIds)
	require.Equal(t, arg.ActionTypeForAgreement, agreementForm.ActionTypeForAgreement)
	require.Equal(t, arg.WholesalePrice, agreementForm.WholesalePrice)
	require.Equal(t, arg.RetailPrice, agreementForm.RetailPrice)

	require.NotZero(t, agreementForm.ID)

	return agreementForm
}

func TestCretateAgreementForm(t *testing.T) {
	agreement1 := createRandomAgreementForm(t, string(ActionTypeBuy))
	require.NotEmpty(t, agreement1)
	agreement2 := createRandomAgreementForm(t, string(ActionTypeSell))
	require.NotEmpty(t, agreement2)
	agreement3 := createRandomAgreementForm(t, string(ActionTypeDefect))
	require.NotEmpty(t, agreement3)
	agreement4 := createRandomAgreementForm(t, string(ActionTypeRefund))
	require.NotEmpty(t, agreement4)
}

func TestGetAgreementForm(t *testing.T) {
	agreementForm1 := createRandomAgreementForm(t, string(ActionTypeBuy))
	agreementForm2, err := testStore.GetAgreementForm(context.Background(), agreementForm1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, agreementForm2)

	require.Equal(t, agreementForm1.ID, agreementForm2.ID)
	require.Equal(t, agreementForm1.FromAccount, agreementForm2.FromAccount)
	require.Equal(t, agreementForm1.ToAccount, agreementForm2.ToAccount)
	require.Equal(t, agreementForm1.ProductIds, agreementForm2.ProductIds)
	require.Equal(t, agreementForm1.ActionTypeForAgreement, agreementForm2.ActionTypeForAgreement)
	require.Equal(t, agreementForm1.WholesalePrice, agreementForm2.WholesalePrice)
	require.Equal(t, agreementForm1.RetailPrice, agreementForm2.RetailPrice)
}

func TestListAgreementForms(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomAgreementForm(t, string(ActionTypeSell))
		createRandomAgreementForm(t, string(ActionTypeBuy))
	}

	arg := ListAgreementFormsParams{
		FromAccount: account1.ID,
		ToAccount:   account2.ID,
		Limit:       5,
		Offset:      5,
	}

	agreementForms, err := testStore.ListAgreementForms(context.Background(), arg)
	require.NoError(t, err)
	// require.Len(t, agreementForms, 5)

	for _, agreementForm := range agreementForms {
		require.NotEmpty(t, agreementForm)
		require.True(t, agreementForm.FromAccount == account1.ID ||
			agreementForm.ToAccount == account1.ID)
	}
}

// func TestUpdateAgreementForm(t *testing.T) {
// 	account1 := createRandomAccount(t)
// 	product := createRandomProduct(t)

// 	arg := UpdateAgreementFormParams{
// 		ProductIds:             []int32{product.ID},
// 		ActionTypeForAgreement: ActionTypeNone,
// 		WholesalePrice:         fake.Float64(2, 10, 1000),
// 		RetailPrice:            fake.Float64(2, 10, 1000),
// 		ID:                     account1.ID,
// 	}

// 	account2, err := testStore.UpdateAgreementForm(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, account2)

// 	// require.Equal(t,account1., account2.)
// 	// require.Equal(t,account1., account2.)
// 	// require.Equal(t,account1.ToAccount, account2.ToAccount)
// 	// require.Equal(t,account1.ProductIds, account2.ProductIds)
// 	// require.Equal(t,account1.ActionTypeForAgreement, account2.ActionTypeForAgreement)
// 	// require.Equal(t,account1.WholesalePrice, account2.WholesalePrice)
// 	// require.Equal(t,account1.RetailPrice, account2.RetailPrice)

// }

func TestDeleteAgreementForm(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testStore.DeleteAgreementForm(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testStore.GetAgreementForm(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, account2)
}
