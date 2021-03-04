package db

import (
    "github.com/stretchr/testify/require"
    "testing"
    "simplebank/util"
    "context"
    "time"
    "database/sql"
)

func createRandomAccount(t *testing.T) Account{
    arg := CreateAccountParams{
        Owner:    util.RandomOwner(),
        Balance:  util.RandomMoney(),
        Currency: util.RandomCurrency(),
    }
    //t.Logf("arg=%+v", arg)
    account, err := testQueies.CreateAccount(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, account)
   // t.Logf("account=%+v", account)
    require.Equal(t, arg.Owner, account.Owner)
    require.Equal(t, arg.Balance, account.Balance)
    require.Equal(t, arg.Currency, account.Currency)
    
    require.NotZero(t, account.ID)
    require.NotZero(t, account.CreatedAt)
    
    return account
}

func TestCreateAccount(t *testing.T){
    createRandomAccount(t)
}

func TestGetAccount(t *testing.T){
    account1 := createRandomAccount(t)
    account2, err := testQueies.GetAccount(context.Background(), account1.ID)
    require.NoError(t, err)
    
    require.Equal(t, account1.ID, account2.ID)
    require.Equal(t, account1.Owner, account2.Owner)
    require.Equal(t, account1.Balance, account2.Balance)
    require.Equal(t, account1.Currency, account2.Currency)
    require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T){
    account1 := createRandomAccount(t)
    arg := UpdateAccountParams{
        ID :account1.ID,
        Balance:util.RandomMoney(),
    }
    account2, err := testQueies.UpdateAccount(context.Background(), arg)
    require.NoError(t, err)

    require.Equal(t, account1.ID, account2.ID)
    require.Equal(t, account1.Owner, account2.Owner)
    require.Equal(t, arg.Balance, account2.Balance)
    require.Equal(t, account1.Currency, account2.Currency)
    require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T)  {
    account1 := createRandomAccount(t)
    
    err := testQueies.DeleteAccount(context.Background(), account1.ID)
    require.NoError(t, err)
    
    acount2, err := testQueies.GetAccount(context.Background(), account1.ID)
    require.EqualError(t, err, sql.ErrNoRows.Error())
    require.Empty(t, acount2)
}


func TestListAccount(t *testing.T) {
    for i:=0;i<10;i++ {
        createRandomAccount(t)
    }
    
    arg := ListAccountsParams{
        Limit:  5,
        Offset: 5,
    }
    
    accounts, err := testQueies.ListAccounts(context.Background(), arg)
    require.NoError(t, err)
    
    require.Len(t, accounts, 5)
    for  _,account := range accounts {
        require.NotEmpty(t, account)
    }
}