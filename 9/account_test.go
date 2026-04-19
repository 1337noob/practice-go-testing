package account_test

import (
	"account"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithdraw_InsufficientFunds(t *testing.T) {
	acc := account.Account{Balance: 10}
	err := acc.Withdraw(20)

	assert.Error(t, err)
	assert.ErrorIs(t, err, account.ErrInsufficientFunds)
}

func TestWithdraw_Success(t *testing.T) {
	acc := account.Account{Balance: 30}
	err := acc.Withdraw(10)

	assert.NoError(t, err)
	assert.Equal(t, 20, acc.Balance)
}
