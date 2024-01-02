package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntryGroup(t *testing.T) EntryGroup {
	arg := CreateEntryGroupParams{
		Quantity:         int32(fake.RandomDigit()),
		ActionType:       ActionTypeBuy,
		PricingType:      PricingTypeWholesale,
		Price:            fake.RandomFloat(2, 100, 1000),
		Currency:         CurrencyCodeUsd,
		EntryGroupStatus: EntryGroupStatusAccepted,
	}

	entryGroup, err := testStore.CreateEntryGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entryGroup)

	require.Equal(t, arg.Quantity, entryGroup.Quantity)
	require.Equal(t, arg.ActionType, entryGroup.ActionType)
	require.Equal(t, arg.PricingType, entryGroup.PricingType)
	require.Equal(t, arg.Price, entryGroup.Price)
	require.Equal(t, arg.Currency, entryGroup.Currency)
	require.Equal(t, arg.EntryGroupStatus, entryGroup.EntryGroupStatus)

	require.NotZero(t, entryGroup.ID)
	require.NotZero(t, entryGroup.CreatedAt)

	return entryGroup
}

func TestCreateEntryGroup(t *testing.T) {
	createRandomEntryGroup(t)
}

func TestGetEntryGroup(t *testing.T) {
	entryGroup1 := createRandomEntryGroup(t)
	entryGroup2, err := testStore.GetEntryGroup(context.Background(), entryGroup1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entryGroup2)

	require.Equal(t, entryGroup1.ID, entryGroup2.ID)
	require.Equal(t, entryGroup1.Quantity, entryGroup2.Quantity)
	require.Equal(t, entryGroup1.ActionType, entryGroup2.ActionType)
	require.Equal(t, entryGroup1.PricingType, entryGroup2.PricingType)
	require.Equal(t, entryGroup1.Price, entryGroup2.Price)
	require.Equal(t, entryGroup1.Currency, entryGroup2.Currency)
	require.Equal(t, entryGroup1.EntryGroupStatus, entryGroup2.EntryGroupStatus)
	require.WithinDuration(t, entryGroup1.CreatedAt.Time, entryGroup2.CreatedAt.Time, time.Second)
}

func TestListEntryGroup(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntryGroup(t)
	}

	arg := ListEntryGroupsParams{
		Limit:  5,
		Offset: 5,
	}

	entryGroups, err := testStore.ListEntryGroups(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entryGroups, 5)

	for _, entryGroup := range entryGroups {
		require.NotEmpty(t, entryGroup)
	}
}
