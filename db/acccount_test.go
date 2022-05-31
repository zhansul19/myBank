package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	user := CreateAccountParams{
		Owner:    "Tom",
		Balance:  100,
		Currency: "USD",
	}
	account, err := testQueries.CreateAccount(context.Background(), user)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t,user.Owner,account.Owner)
	require.Equal(t,user.Balance,account.Balance)
	require.Equal(t,user.Currency,account.Currency)

	require.NotZero(t,account.ID)
	require.NotZero(t,account.CreatedAt)

}
