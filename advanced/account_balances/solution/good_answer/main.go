package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	mu      sync.Mutex // Protects concurrent access to balance
	balance int
}

// Deposit adds an amount to the account.
func (acct *Account) Deposit(txnID, amt int) {
	acct.mu.Lock() // Ensure exclusive access during deposit
	defer acct.mu.Unlock()
	fmt.Printf("\nTransaction %d: Depositing $%d", txnID, amt)
	acct.balance += amt
	fmt.Printf("\nNew account balance: $%d", acct.balance)
}

// Withdraw deducts an amount from the account. Returns error if insufficient funds.
func (acct *Account) Withdraw(txnID, amt int) {
	acct.mu.Lock() // Ensure exclusive access during withdrawal
	defer acct.mu.Unlock()
	fmt.Printf("\nTransaction %d: Withdrawing $%d", txnID, amt)
	if acct.balance < amt {
		fmt.Printf("\nTransaction %d: ERROR balance of $%d  - insufficient funds to withdraw $%d", txnID, acct.balance, amt)
		return
	}
	acct.balance -= amt
	fmt.Printf("\nNew account balance: $%d", acct.balance)
}
func main() {
	acct := &Account{balance: rand.Intn(900) + 101}
	fmt.Printf("Starting account balance: $%d\n", acct.balance)
	defer fmt.Printf("\n\nFinal account balance: $%d", acct.balance)
	for i := 0; i < 10; i++ {
		txnID := i + 1
		if rand.Intn(2) == 0 {
			go acct.Deposit(txnID, rand.Intn(100)+1)
			continue
		}
		go acct.Withdraw(i+1, rand.Intn(100)+1)
	}
	time.Sleep(time.Second)
}
