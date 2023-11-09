# Interview Question: Account Balances

The original code provided for the interview question contains several issues, especially regarding concurrent transactions and lack of error handling for account operations.

### Problematic Code:

#### Concurrency Issue:
```go
for i := 0; i < 10; i++ {
    if rand.Intn(2) == 0 {
        go acc.Deposit(i+1, rand.Intn(100)+1)
        continue
    }
    go acc.Withdraw(i+1, rand.Intn(100)+1)
}
```
Concurrent deposit and withdrawal operations without synchronization can lead to race conditions, causing inconsistent account balances.

#### Lack of Error Handling:
```go
func (a *Account) Withdraw(i, amt int) {
    fmt.Printf("\nTransaction %d: Withdraw %d from %s", i, amt, a)
    a.balance -= amt
    fmt.Printf("New %s", a)
}
```
There's no check for insufficient funds during withdrawal, which might lead to negative balances.

## Solution

### Handling Concurrency:

In `main.go`:
```go
var wg sync.WaitGroup

for i := 0; i < NumTransactions; i++ {
    ...
    wg.Add(1)
    go func(txn Transaction) {
        defer wg.Done()
        ...
    }(txn)
}

wg.Wait() // Wait for all goroutines to complete.
```
Using a `WaitGroup` ensures synchronization and that all goroutines complete their operations.

### Improved Error Handling:

In `types.go`:
```go
func (a *Account) Debit(amount int) error {
    if a.balance < amount {
        return fmt.Errorf("insufficient funds. (Balance: $%d)", a.balance)
    }
    a.balance -= amount
    return nil
}
```
With this approach, withdrawals are checked for sufficient funds, and an error is returned if funds are insufficient.

### Modular Approach:
Splitting the logic into different files (`types.go` for data types, `transaction.go` for transaction logic, and `main.go` for the main execution) makes the codebase more readable, maintainable, and testable.
