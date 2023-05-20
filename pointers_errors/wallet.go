package pointers_errors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds, cannot withdraw")

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(a Bitcoin) {
	w.balance += a
}

func (w *Wallet) Withdraw(a Bitcoin) error {
	if a > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= a
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
