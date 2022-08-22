package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

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

func TestGetAccount(t *testing.T) {
	randAccount := createRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), randAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, randAccount.Owner, account.Owner)
	require.Equal(t, randAccount.Balance, account.Balance)
	require.Equal(t, randAccount.Currency, account.Currency)

	require.Equal(t, randAccount.ID, account.ID)
	require.WithinDuration(t, randAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	randAccount := createRandomAccount(t)
	args := UpdateAccountParams{
		ID:      randAccount.ID,
		Balance: util.RandomMoney(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, randAccount.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, randAccount.Currency, account.Currency)

	require.Equal(t, randAccount.ID, account.ID)
	require.WithinDuration(t, randAccount.CreatedAt, account.CreatedAt, time.Second)
}
func TestDeleteAccount(t *testing.T) {
	randAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), randAccount.ID)
	require.NoError(t, err)

	accountDel, err := testQueries.GetAccount(context.Background(), randAccount.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountDel)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	if err != nil {
		for _, v := range accounts {
			require.Empty(t, v)
		}
	}
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, v := range accounts {
		require.NotEmpty(t, v)
	}
}
