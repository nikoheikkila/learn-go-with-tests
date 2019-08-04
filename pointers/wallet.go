package pointers

import (
	"errors"
	"fmt"
)

var insufficientFundsError = errors.New("Insufficient funds in wallet.")

type Bitcoin int

type Stringer interface {
	String() string
}

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return insufficientFundsError
	}

	wallet.balance -= amount

	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
