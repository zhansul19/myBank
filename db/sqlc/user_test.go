package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zhansul19/myBank/util"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
func createRandomUser(t *testing.T) Users {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestGetUser(t *testing.T) {
	randUser := createRandomUser(t)

	user, err := testQueries.GetUser(context.Background(), randUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, randUser.Username, user.Username)
	require.Equal(t, randUser.FullName, user.FullName)
	require.Equal(t, randUser.Email, user.Email)
	require.Equal(t, randUser.HashedPassword, user.HashedPassword)
	
	require.WithinDuration(t, randUser.CreatedAt, user.CreatedAt, time.Second)
}
