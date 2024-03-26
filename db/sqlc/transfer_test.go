package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {

	created_acccount_first := createRandomAccount(t)
	created_account_second := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccount: created_acccount_first.ID,
		ToAccount:   created_account_second.ID,
		Amount:      created_acccount_first.Balance,
	}

	created_transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, created_transfer)

	require.Equal(t, created_transfer.FromAccount, args.FromAccount)
	require.Equal(t, created_transfer.ToAccount, args.ToAccount)
	require.Equal(t, created_transfer.Amount, args.Amount)

	require.NotZero(t, created_transfer.ID)
	require.NotZero(t, created_transfer.CreatedAt)

	return created_transfer

}

func TestCreateTransfer(t *testing.T) {

	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {

	created_transfer := createRandomTransfer(t)

	received_transfer, err := testQueries.GetTransfer(context.Background(), created_transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, received_transfer)

	require.Equal(t, created_transfer.ID, received_transfer.ID)
	require.Equal(t, created_transfer.ToAccount, received_transfer.ToAccount)
	require.Equal(t, created_transfer.FromAccount, received_transfer.FromAccount)
	require.WithinDuration(t, created_transfer.CreatedAt, received_transfer.CreatedAt, time.Second)

}
