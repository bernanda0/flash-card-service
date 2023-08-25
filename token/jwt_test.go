package token

import (
	"br/simple-service/utils"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJwtMaker(t *testing.T) {
	maker, err := NewJwtMaker(utils.RandomString(32))
	require.NoError(t, err)

	// params in generate token
	account_id := utils.RandomID()
	username := utils.RandomString(10)
	duration := time.Minute * 5

	// for testing
	issued_at := time.Now()
	expired_at := time.Now().Add(duration)

	token, err := maker.GenerateToken(uint(account_id), username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	fmt.Println(token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	fmt.Println(payload)

	require.Equal(t, payload.AccountID, uint(account_id))
	require.WithinDuration(t, payload.IssuedAt, issued_at, time.Second)
	require.WithinDuration(t, payload.ExpiredAt, expired_at, time.Second)

}

func TestTokenExpired(t *testing.T) {
	maker, err := NewJwtMaker(utils.RandomString(32))
	require.NoError(t, err)

	// params in generate token
	account_id := utils.RandomID()
	username := utils.RandomString(10)
	// minus indicate that the token is expired
	duration := -time.Minute

	token, err := maker.GenerateToken(uint(account_id), username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	fmt.Println(token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
	fmt.Println(err)

}
