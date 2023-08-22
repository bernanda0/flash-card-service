package sqlc

import (
	"br/simple-service/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	argAccount := CreateAccountParams{
		Username:     utils.RandomString(10),
		Email:        utils.RandomString(7) + "@example.com",
		PasswordHash: utils.RandomString(20),
	}

	account, err := q_test.CreateAccount(context.Background(), argAccount)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, argAccount.Username, account.Username)
	require.Equal(t, argAccount.Email, account.Email)
	require.Equal(t, argAccount.PasswordHash, account.PasswordHash)
	require.NotZero(t, account.AccountID)
	require.NotZero(t, account.CreatedAt)

}

func TestDeleteAccount(t *testing.T) {
	// Create a test account first
	argAccount := CreateAccountParams{
		Username:     utils.RandomString(10),
		Email:        utils.RandomString(7) + "@example.com",
		PasswordHash: utils.RandomString(20),
	}

	createdAccount, err := q_test.CreateAccount(context.Background(), argAccount)
	require.NoError(t, err)

	// Delete the created account
	deletedAccount, err := q_test.DeleteAccount(context.Background(), createdAccount.AccountID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedAccount)

	require.Equal(t, createdAccount.AccountID, deletedAccount.AccountID)
	require.Equal(t, createdAccount.Username, deletedAccount.Username)
	require.Equal(t, createdAccount.Email, deletedAccount.Email)
	require.Equal(t, createdAccount.PasswordHash, deletedAccount.PasswordHash)
	require.Equal(t, createdAccount.CreatedAt, deletedAccount.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	// Create a test account first
	argAccount := CreateAccountParams{
		Username:     utils.RandomString(10),
		Email:        utils.RandomString(7) + "@example.com",
		PasswordHash: utils.RandomString(20),
	}

	createdAccount, err := q_test.CreateAccount(context.Background(), argAccount)
	require.NoError(t, err)

	// Get the created account
	retrievedAccount, err := q_test.GetAccount(context.Background(), createdAccount.AccountID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	require.Equal(t, createdAccount.AccountID, retrievedAccount.AccountID)
	require.Equal(t, createdAccount.Username, retrievedAccount.Username)
	require.Equal(t, createdAccount.Email, retrievedAccount.Email)
	require.Equal(t, createdAccount.PasswordHash, retrievedAccount.PasswordHash)
	require.Equal(t, createdAccount.CreatedAt, retrievedAccount.CreatedAt)
}

func TestListAccounts(t *testing.T) {
	// Assuming you have test data in your database, you can directly call the ListAccounts function
	accounts, err := q_test.ListAccounts(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	// Perform assertions on the list of accounts
	for _, account := range accounts {
		// You can add specific assertions for each account in the list
		require.NotZero(t, account.AccountID)
		require.NotEmpty(t, account.Username)
		require.NotEmpty(t, account.Email)
		require.NotEmpty(t, account.PasswordHash)
		require.NotZero(t, account.CreatedAt)
	}

	// You can also perform more complex assertions based on your specific use case
}
