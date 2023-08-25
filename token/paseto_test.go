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

	token, err := maker.GenerateToken(uint(accountID), username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	fmt.Println(token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	fmt.Println(payload)

	require.Equal(t, payload.AccountID, uint(accountID))
	require.Equal(t, payload.Username, username)
}
