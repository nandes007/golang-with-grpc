package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/nandes007/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, s Account, r Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: s.ID,
		ToAccountID:   r.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	s := createRandomAccount(t)
	r := createRandomAccount(t)

	createRandomTransfer(t, s, r)
}

func TestGetTransfer(t *testing.T) {
	s := createRandomAccount(t)
	r := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, s, r)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	s := createRandomAccount(t)
	r := createRandomAccount(t)

	transfer1 := createRandomTransfer(t, s, r)

	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	s := createRandomAccount(t)
	r := createRandomAccount(t)
	transfer := createRandomTransfer(t, s, r)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Empty(t, transfer2)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
}

func TestGetListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := createRandomAccount(t)
		r := createRandomAccount(t)
		createRandomTransfer(t, s, r)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, int(arg.Offset))

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
