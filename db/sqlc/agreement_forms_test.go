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

	arg := CreateAgreementFormsParams{
		FromAccount:            accFrom.ID,
		ToAccount:              accTo.ID,
		ProductIds:             []int32{product.ID},
		ActionTypeForAgreement: ActionType(actionType),
		WholesalePrice:         1.1,
		RetailPrice:            1.1,
	}

	agreementForm, err := testStore.CreateAgreementForms(context.Background(), arg)
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

func TestAgreementFormCreate(t *testing.T) {
	agreement1 := createRandomAgreementForm(t, string(ActionTypeBuy))
	require.NotEmpty(t, agreement1)
	agreement2 := createRandomAgreementForm(t, string(ActionTypeSell))
	require.NotEmpty(t, agreement2)
}
