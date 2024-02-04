package db

import (
	"context"
	"git/db/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)


func createRandomAccount(t *testing.T) Account {
		arg := CreateAccountParams {
			Owner : util.RandomOnwer(),
			Balance: util.RandomMoney(), 
			Currency : util.RandomeCurrency(),
		}
		
		account, err := testQueries.CreateAccount(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t,account)
		require.Equal(t, arg.Owner, account.Owner)
		require.Equal(t, arg.Balance, account.Balance)
		require.Equal(t, arg.Currency, account.Currency)
		
		require.NotZero(t, account.ID)
		require.NotZero(t,account.CreatedAt)
		
		return account 
}


func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	// create account 
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccountForUpdate(context.Background(), account1.ID)

	require.NoError(t,err)
	require.NotEmpty(t, account2)
	
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt,time.Second )


}

func TestUpdateAccount(t *testing.T) {
	account1 :=createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account1.ID,
		Balance: util.RandomMoney(),
	  
	}
	
	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t,err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt,time.Second )

}


func TestListAccount(t *testing.T) {
	for i := 0; i<10; i++ {
		createRandomAccount(t)
	}

	arg := listAccountsParams {
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.listAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts,5)

	for _, account := range accounts {
		require.NotEmpty(t,account)
	}

	
}