package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Constants to define behavior of the simulation.
const (
	MaxInitialBalance = 1000 // Maximum starting balance.
	MinInitialBalance = 100  // Minimum starting balance.
	NumTransactions   = 10   // Number of transactions to process.
)

func main() {
	// Initialize an account with a random starting balance.
	acc := &Account{balance: rand.Intn(MaxInitialBalance-MinInitialBalance+1) + MinInitialBalance}
	handler := NewTransactionHandler(acc) // Transaction handler for the account.
	var wg sync.WaitGroup

	fmt.Printf("Start balance:\t$%d\n", acc.balance)

	for i := 0; i < NumTransactions; i++ {
		tType := DebitType
		if rand.Intn(2) == 0 {
			tType = CreditType
		}
		txn := Transaction{
			Type:   tType,
			Amount: rand.Intn(acc.balance) + 1, // Random transaction amount.
		}

		wg.Add(1)
		go func(txn Transaction) { // Start a goroutine for each transaction.
			defer wg.Done()

			newBalance, err := handler.Process(txn)
			if err != nil {
				fmt.Printf("\n> %v:\t$%d\n> ERROR:\t%s\n", txn.Type, txn.Amount, err)
				return
			}
			fmt.Printf("\n> %v:\t$%d\n> Balance:\t$%d\n", txn.Type, txn.Amount, newBalance)
		}(txn)
	}

	wg.Wait() // Wait for all goroutines to complete.

	fmt.Printf("\nEnd balance:\t$%d\n", acc.balance) // Print the final balance.
}
