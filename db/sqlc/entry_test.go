package db

import (
	"context"
	"testing"
	"time"

	"github.com/Placebo900/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	acc := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := queries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)
	entry2, err := queries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry, entry2)
}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}
	entry2, err := queries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.NotEqual(t, entry.Amount, entry2.Amount)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.ID, entry2.ID)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := queries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry2, err := queries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.Empty(t, entry2)
	require.NotEqual(t, entry, entry2)
}

func TestListEntry(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	arg := ListEntryParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := queries.ListEntry(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for i := range entries {
		require.NotEmpty(t, entries[i])
	}
}
