package account

import "errors"

var ErrInsufficientFunds = errors.New("insufficient funds")

type Account struct {
	Balance int
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.Balance {
		return ErrInsufficientFunds
	}

	a.Balance -= amount

	return nil
}
