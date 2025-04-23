package main

import (
	"errors"
	"fmt"
)

// task 1
type BankAccount struct {
	amount float64
}

func (account *BankAccount) Deposit(amount float64) {
	account.amount += amount
}

func (account *BankAccount) Withdraw(amount float64) error {
	if account.amount >= amount {
		account.amount -= amount
		return nil
	} else {
		return errors.New("not enough money")
	}
}

func (account BankAccount) Balance() float64 {
	return account.amount
}

func main() {
	acc := BankAccount{amount: 0}
	fmt.Println(acc.Balance())
	acc.Deposit(32.56)
	fmt.Println(acc.Balance())
	fmt.Println(acc.Withdraw(32.06))
	fmt.Println(acc.Balance())
}
