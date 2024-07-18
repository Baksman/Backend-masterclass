package db

import (
	"context"
	// "database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/baksman/backend_masterclass/util"
)

func createRandomUser(t *testing.T) User {

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.PasswordChangedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}

// func TestUpdateUser(t *testing.T) {
// 	account1 := createRandomUser(t)

// 	arg := UpdateAccountParams{
// 		ID:      account1.ID,
// 		Balance: util.RandomMoney(),
// 	}

// 	account2, err := testQueries.UpdateAccount(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.Equal(t, account1.ID, account2.ID)
// 	require.Equal(t, account1.Owner, account2.Owner)
// 	require.Equal(t, arg.Balance, account2.Balance)
// 	require.Equal(t, account1.Currency, account2.Currency)
// 	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
// }

// func TestDeleteAccount(t *testing.T) {
// 	account1 := createRandomUser(t)

// 	err := testQueries.DeleteAccount(context.Background(), account1.ID)
// 	require.NoError(t, err)
// 	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, account2)

// }

// func TestListAccount(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomUser(t)
// 	}

// 	args := ListAccountParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}

// 	accounts, err := testQueries.ListAccount(context.Background(), args)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, accounts)
// 	require.Len(t, accounts, 5)
// 	for _, account := range accounts {
// 		require.NotEmpty(t, account)
// 	}
// }
