package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zhansul19/myBank/util"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
func createRandomAccount(t *testing.T) Accounts {
	user := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), user)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, user.Owner, account.Owner)
	require.Equal(t, user.Balance, account.Balance)
	require.Equal(t, user.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
