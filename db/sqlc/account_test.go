package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4/zero"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		UserID: user.ID,
		Name:   fake.Food().Vegetable(),
	}

	account, err := testStore.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, arg.UserID, account.UserID)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testStore.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1, account2)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:   account1.ID,
		Name: zero.StringFrom(fake.Food().Vegetable()),
	}

	account2, err := testStore.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	assert.NotEmpty(t, account2)
	assert.Equal(t, arg.Name.String, account2.Name)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testStore.DeleteAccount(context.Background(), account1.ID)
	assert.NoError(t, err)

	account2, err := testStore.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	assert.EqualError(t, err, ErrRecordNotFound.Error())
	assert.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	var createdAccounts []Account
	for i := 0; i < 10; i++ {
		createdAccounts = append(createdAccounts, createRandomAccount(t))
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)

	for idx, account := range accounts {
		assert.NotEmpty(t, account)
		assert.Equal(t, createdAccounts[len(createdAccounts)-idx-1].ID, account.ID)
		assert.Equal(t, createdAccounts[len(createdAccounts)-idx-1].Name, account.Name)
	}
}
