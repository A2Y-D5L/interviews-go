# Account Balances Challenge

There are a number of issues in the code snippet below.

1. Read through the code thoroughly, explaining its functionality as you go.
2. Describe any problems you see in the code and suggest solutions.
3. If time permits, you can test your solutions interactively in the [Go playground](https://go.dev/play/p/qntUS3sRpqk).

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Account struct {
	balance int
}

func (a *Account) String() string {
	return fmt.Sprintln("account balance:", a.balance)
}

func (a *Account) Deposit(txnID, amt int) {
	fmt.Printf("\n%d: Deposit %d into %s", txnID, amt, a)
	a.balance += amt
	fmt.Printf("New %s", a)
}

func (a *Account) Withdraw(txnID, amt int) {
	fmt.Printf("\n%d: Withdraw %d from %s", txnID, amt, a)
	a.balance -= amt
	fmt.Printf("New %s", a)
}

func main() {
	acc := &Account{balance: rand.Intn(900) + 101}
	fmt.Printf("Starting %s", acc)
	defer fmt.Printf("\nFinal %s", acc)
	for i := 0; i < 10; i++ {
		if rand.Intn(2) == 0 {
			go acc.Deposit(i+1, rand.Intn(100)+1)
			continue
		}
		go acc.Withdraw(i+1, rand.Intn(100)+1)
	}
	time.Sleep(time.Second)
}
```