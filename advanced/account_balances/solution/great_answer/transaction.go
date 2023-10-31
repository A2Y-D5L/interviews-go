package main

import "fmt"

// TransactionHandler processes transactions for an account.
type TransactionHandler struct {
	account *Account // Reference to the associated account.
}

// NewTransactionHandler initializes and returns a new TransactionHandler.
func NewTransactionHandler(account *Account) *TransactionHandler {
	return &TransactionHandler{account: account}
}

// Process handles a given transaction, updating the account balance accordingly.
// It returns the new balance and any potential error encountered.
func (th *TransactionHandler) Process(txn Transaction) (int, error) {
	switch txn.Type {
	case CreditType:
		th.account.Credit(txn.Amount) // Process credit transaction.
	case DebitType:
		err := th.account.Debit(txn.Amount) // Process debit transaction.
		if err != nil {
			return th.account.balance, err
		}
	default:
		// Handle unknown transaction types.
		return th.account.balance, fmt.Errorf("unknown transaction type")
	}
	return th.account.balance, nil
}
