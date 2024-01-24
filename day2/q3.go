package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	mu      sync.Mutex // for own mutex of each instance
}

// create new account
func createAccount(initialBalance int) *Account {
	return &Account{balance: initialBalance}
}

// Deposit
func (a *Account) Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	a.balance += amount
	fmt.Printf("Deposit!! new balance: Rs.%d\n", a.balance)
}

// Withdraw
func (a *Account) Withdraw(amount int) {
	if a.balance-amount > 0 {
		mu.Lock()
		defer mu.Unlock()
		a.balance -= amount
		fmt.Printf("Withdraw!! new balance is %d\n", a.balance)

	} else {
		fmt.Println("Insufficient Balance")
	}
}

func main() {
	amount := createAccount(500)
	go func() {
		for true {
			time.Sleep(time.Second)
			amount.Deposit(100)
		}
	}()
	go func() {
		for true {
			time.Sleep(time.Second)
			amount.Withdraw(50)
		}
	}()
	time.Sleep(20 * time.Second) // to start
}
