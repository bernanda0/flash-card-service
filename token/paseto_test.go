package token

import (
	"br/simple-service/utils"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	// params in generate token
	accountID := utils.RandomID()
	username := utils.RandomString(10)
	duration := time.Minute * 5

	token, payload, err := maker.GenerateToken(uint(accountID), username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	fmt.Println(token)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	fmt.Println(payload)

	require.Equal(t, payload.AccountID, uint(accountID))
	require.Equal(t, payload.Username, username)
}

func TestPasetoExpired(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	// params in generate token
	account_id := utils.RandomID()
	username := utils.RandomString(10)
	// minus indicate that the token is expired
	duration := -time.Minute

	token, payload, err := maker.GenerateToken(uint(account_id), username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	fmt.Println(token)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
	fmt.Println(err)

}
