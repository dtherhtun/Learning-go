package main

import "fmt"

type BankAccount struct {
	balance int
	owner   string
}

func NewBankAccount(owner string) *BankAccount {
	return &BankAccount{
		owner: owner,
	}
}

func (account *BankAccount) Debit(amount int) {
	if amount >= 0 && amount <= account.balance {
		account.balance -= amount
	}
}

func (account *BankAccount) Credit(amount int) {
	if amount >= 0 {
		account.balance += amount
	}
}

func (account *BankAccount) String() string {
	return fmt.Sprintf("Account Owner => %s, balance => %d", account.owner, account.balance)
}

func main() {
	account1 := NewBankAccount("dther")
	account1.Credit(5000)
	account1.Debit(3000)
	fmt.Println(account1)
}
