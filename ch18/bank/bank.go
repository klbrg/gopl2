// Package bank provides a toy account with an audit log.
package bank

import (
	"errors"
	"log/slog"
)

type Account struct {
	Log     *slog.Logger
	balance int
}

var ErrInsufficient = errors.New("insufficient funds")

func (a *Account) Deposit(amount int) {
	a.balance += amount
	a.Log.Info("deposit", "amount", amount, "balance", a.balance)
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.balance {
		a.Log.Warn("withdrawal refused",
			slog.Int("amount", amount), slog.Int("balance", a.balance))
		return ErrInsufficient
	}
	a.balance -= amount
	a.Log.Info("withdrawal", "amount", amount, "balance", a.balance)
	return nil
}
