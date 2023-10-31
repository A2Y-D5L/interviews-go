package main

import "fmt"

// Account represents a basic banking account with a balance.
type Account struct {
	balance int
}

// Credit adds the specified amount to the account balance.
func (a *Account) Credit(amount int) {
	a.balance += amount
}

// Debit removes the specified amount from the account balance.
// Returns an error if there are insufficient funds.
func (a *Account) Debit(amount int) error {
	if a.balance < amount {
		return fmt.Errorf("insufficient funds. (Balance: $%d)", a.balance)
	}
	a.balance -= amount
	return nil
}

// TransactionType represents the type of a transaction (Credit or Debit).
type TransactionType int

// String returns a readable representation of the transaction type.
func (tt TransactionType) String() string {
	switch tt {
	case CreditType:
		return "Credit"
	case DebitType:
		return "Debit"
	default:
		return "Unknown Transaction Type"
	}
}

// Constants representing the types of transactions.
const (
	CreditType TransactionType = iota
	DebitType
)

// Transaction represents a banking transaction with type, amount, and an ID.
type Transaction struct {
	Type   TransactionType
	Amount int
	TxnID  int
}