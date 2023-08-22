package sqlc

import (
	"br/simple-service/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccount(t *testing.T) {
	mockdb := new(MockDBTX)
	queries := New(mockdb)

	argAccount := CreateAccountParams{
		Username:     utils.RandomString(10),
		Email:        utils.RandomString(6) + "@example.com",
		PasswordHash: utils.RandomString(32),
	}

	// Use mockdb.On to specify the expected method call and return value
	mockdb.On("QueryRowContext", mock.Anything, createAccount, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(argAccount.Username, argAccount.Email, argAccount.PasswordHash)

	account, err := queries.CreateAccount(context.Background(), argAccount)
	assert.NoError(t, err)

	// Check the fields of the returned account
	assert.Equal(t, argAccount.Username, account.Username)
	assert.Equal(t, argAccount.Email, account.Email)
	assert.Equal(t, argAccount.PasswordHash, account.PasswordHash)

	assert.NotNil(t, account.AccountID)
	assert.NotNil(t, account.CreatedAt)

	mockdb.AssertExpectations(t)
}

// Implement similar test functions for other query methods.
// Remember to replace the expected values with values relevant to your test cases.
