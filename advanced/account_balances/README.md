# Account Transactions Challenge

This challenge assesses handling of concurrency, transaction ordering, interface adherence, and code clarity within the context of banking transactions.

## Challenge

1.	Understand the code’s functionality.
2.	Identify code issues.
3.	Suggest fixes for issues.
4.	Optionally, test solutions in the Go playground.

You can find the challenge code in [here](./challenge).

## What to Look For

1. **Concurrency:**\
	Check for race conditions due to concurrent transaction execution.Transactions must be atomic and consistent.
2.	**Transaction Order:**\
	Transactions should maintain their intended order to prevent issues like account overdrafts.
3.	**Interface Usage:**\
	Ensure the Transaction interface and its implementations are correctly used and behave as expected.
4.	**Transfers:**\
	When transferring between accounts, ensure money deducted from one account is added to another without errors.
5.	**Code Quality:**\
	Examine the clarity of code, variable names, and method structure. The code should be easy to understand and follow.
6.	**Transaction Types:**\
	Verify transactions are handled right based on their type: credit, debit, or transfer.
