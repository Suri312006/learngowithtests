package pointersnerrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	// if something is lowercase, it is private outside of the package its defined in
	balance Bitcoin
}

// this takes in a pointer to the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// takes in a struct pointer, that you can use freely inside since its dereferenced
func (w *Wallet) Deposit(amt Bitcoin) {
	//you can also do this to deference the pointer "Like in rust"
	//(*w).balance
	w.balance += amt
	// fmt.Printf("address of balance in test is %p \n", &w.balance)
}


var ErrInsufficientFunds = errors.New("Cannot withdraw, insufficient funds")

// should throw an error if youre trying to take too much
func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
        return ErrInsufficientFunds
	}

	w.balance -= amt
	return nil
}
