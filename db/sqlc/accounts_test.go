package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Placebo900/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	acc, err := queries.CreateAccount(context.Background(), arg)
	require.NoError(t, err, sql.ErrNoRows)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)
	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc := createRandomAccount(t)
	acc2, err := queries.GetAccount(context.Background(), acc.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc, acc2)
}

func TestUpdateAccount(t *testing.T) {
	acc := createRandomAccount(t)
	acc.Balance = util.RandomMoney()
	arg := UpdateAccountParams{ID: acc.ID, Balance: acc.Balance}
	acc2, err := queries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc, acc2)
}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)
	err := queries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)

	acc2, err := queries.GetAccount(context.Background(), acc.ID)
	require.Error(t, err)
	require.Empty(t, acc2)
	require.NotEqual(t, acc, acc2)
}

func TestListAccount(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := queries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
		require.Equal(t, lastAccount.Owner, acc.Owner)
	}
}
