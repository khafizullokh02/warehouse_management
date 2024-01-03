package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntryItem(t *testing.T, product Product) EntryItem {
	entryGroup := createRandomEntryGroup(t)

	arg := CreateEntryItemParams{
		ProductID:    product.ID,
		EntryGroupID: entryGroup.ID,
		SupCode:      fake.RandomStringWithLength(10),
	}

	entryItem, err := testStore.CreateEntryItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entryItem)

	require.Equal(t, arg.ProductID, entryItem.ProductID)
	require.Equal(t, arg.EntryGroupID, entryItem.EntryGroupID)
	require.Equal(t, arg.SupCode, entryItem.SupCode)

	require.NotZero(t, entryItem.ID)
	require.NotZero(t, entryItem.CreatedAt)

	return entryItem
}

func TestCreateEntryItem(t *testing.T) {
	product := createRandomProduct(t)
	createRandomEntryItem(t, product)
}

func TestGetEntryItem(t *testing.T) {
	product := createRandomProduct(t)

	entryItem1 := createRandomEntryItem(t, product)
	entryItem2, err := testStore.GetEntryItem(context.Background(), entryItem1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entryItem2)

	require.Equal(t, entryItem1.ID, entryItem2.ID)
	require.Equal(t, entryItem1.ProductID, entryItem2.ProductID)
	require.Equal(t, entryItem1.EntryGroupID, entryItem2.EntryGroupID)
	require.Equal(t, entryItem1.SupCode, entryItem2.SupCode)
	require.WithinDuration(t, entryItem1.CreatedAt.Time, entryItem2.CreatedAt.Time, time.Second)
}

func TestListEntryItem(t *testing.T) {
	product := createRandomProduct(t)
	for i := 0; i < 10; i++ {
		createRandomEntryItem(t, product)
	}

	arg := ListEntryItemsParams{
		Limit:  5,
		Offset: 5,
	}

	entryItems, err := testStore.ListEntryItems(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entryItems, 5)
}

func TestUpdateEntryItem(t *testing.T) {
	product := createRandomProduct(t)
	entryGroup := createRandomEntryGroup(t)
	entryItem1 := createRandomEntryItem(t, product)

	arg := UpdateEntryItemParams{
		ProductID:    product.ID,
		EntryGroupID: entryGroup.ID,
		SupCode:      product.SupCode,
		ID:           entryItem1.ID,
	}

	entryItem2, err := testStore.UpdateEntryItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entryItem2)

	require.Equal(t, entryItem1.ID, entryItem2.ID)
	require.Equal(t, entryItem1.EntryGroupID, entryItem2.EntryGroupID)
	require.Equal(t, entryItem1.SupCode, entryItem2.SupCode)
	require.WithinDuration(t, entryItem1.CreatedAt.Time, entryItem2.CreatedAt.Time, time.Second)
}

func TestDeleteEntryItem(t *testing.T) {
	product := createRandomProduct(t)
	entryItem1 := createRandomEntryItem(t, product)
	err := testStore.DeleteEntryItem(context.Background(), entryItem1.ID)
	require.NoError(t, err)

	entryItem2, err := testStore.GetEntryItem(context.Background(), entryItem1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, entryItem2)
}
