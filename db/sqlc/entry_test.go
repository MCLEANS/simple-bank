package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/MCLEANS/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {

	created_account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: created_account.ID,
		Amount:    util.RandomMoney(),
	}

	created_entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, created_entry)

	require.Equal(t, created_entry.AccountID, args.AccountID)
	require.Equal(t, created_entry.Amount, args.Amount)

	require.NotZero(t, created_entry.CreatedAt)
	require.NotZero(t, created_entry.ID)

	return created_entry

}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {

	created_entry := createRandomEntry(t)

	received_entry, err := testQueries.GetEntry(context.Background(), created_entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, received_entry)

	require.Equal(t, received_entry.ID, created_entry.ID)
	require.Equal(t, received_entry.AccountID, created_entry.AccountID)
	require.Equal(t, received_entry.Amount, created_entry.Amount)
	require.WithinDuration(t, received_entry.CreatedAt, created_entry.CreatedAt, time.Second)

}

func TestUpdateEntry(t *testing.T) {

	created_entry := createRandomEntry(t)

	args := UpdateEntryParams{
		ID:     created_entry.ID,
		Amount: util.RandomMoney(),
	}

	updated_entry, err := testQueries.UpdateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updated_entry)

	require.Equal(t, created_entry.ID, updated_entry.ID)
	require.Equal(t, created_entry.AccountID, updated_entry.AccountID)
	require.Equal(t, updated_entry.Amount, args.Amount)
	require.WithinDuration(t, created_entry.CreatedAt, updated_entry.CreatedAt, time.Second)

}

func TestDeleteEntry(t *testing.T) {

	created_entry := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), created_entry.ID)

	require.NoError(t, err)

	received_entry, _err := testQueries.GetEntry(context.Background(), created_entry.ID)

	require.Error(t, _err)
	require.EqualError(t, _err, sql.ErrNoRows.Error())
	require.Empty(t, received_entry)

}
