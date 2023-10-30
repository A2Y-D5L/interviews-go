# Account Balances Question

Please identify issues in the code snippet below and suggest solutions.

```go
type Account struct {
    balance int
}

func (a *Account) Deposit(amount int) {
    a.balance += amount
}

func (a *Account) Withdraw(amount int) {
    a.balance -= amount
}

func main() {
    acc := &Account{balance: 1000}

    go func() {
        for i := 0; i < 100; i++ {
            acc.Deposit(10)
        }
    }()

    go func() {
        for i := 0; i < 100; i++ {
            acc.Withdraw(5)
        }
    }()

    time.Sleep(2 * time.Second)
    fmt.Println(acc.balance)
}