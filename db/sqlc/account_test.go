package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/MCLEANS/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {

	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {

	createRandomAccount(t)

}

func TestGetAccount(t *testing.T) {

	created_account := createRandomAccount(t)

	received_account, err := testQueries.GetAccount(context.Background(), created_account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, received_account)

	require.Equal(t, created_account.ID, received_account.ID)
	require.Equal(t, created_account.Owner, received_account.Owner)
	require.Equal(t, created_account.Balance, received_account.Balance)
	require.Equal(t, created_account.Currency, received_account.Currency)
	require.WithinDuration(t, created_account.CreatedAt, received_account.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {

	/** Create an account */
	created_account := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      created_account.ID,
		Balance: util.RandomMoney(),
	}

	updated_account, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updated_account)

	require.Equal(t, updated_account.ID, created_account.ID)
	require.Equal(t, updated_account.Owner, created_account.Owner)
	require.Equal(t, updated_account.Balance, args.Balance)
	require.Equal(t, updated_account.Currency, created_account.Currency)
	require.WithinDuration(t, updated_account.CreatedAt, created_account.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {

	created_acount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), created_acount.ID)

	require.NoError(t, err)

	/** Validate that the account has been deleted */
	received_account, _err := testQueries.GetAccount(context.Background(), created_acount.ID)

	require.Error(t, _err)
	require.EqualError(t, _err, sql.ErrNoRows.Error())
	require.Empty(t, received_account)

}

func TestListAccounts(t *testing.T) {

	/** Create 10 accounts */
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	listed_accounts, err := testQueries.ListAccounts(context.Background(), args)

	require.NoError(t, err)

	for _, account := range listed_accounts {
		require.NotEmpty(t, account)
	}

}
