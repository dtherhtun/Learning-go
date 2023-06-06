package main

import (
	"fmt"
	"log"
)

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
	log.Printf("Insufficient Funds: Your account balance is not sufficient to complete this transaction.")
}

func (account *BankAccount) Credit(amount int) {
	if amount >= 0 {
		account.balance += amount
	}
	log.Printf("Invalid Amount: The amount you entered is invalid. Please enter a valid amount.")
}

func (account *BankAccount) String() string {
	return fmt.Sprintf("Account Owner => %s, balance => %d", account.owner, account.balance)
}

func main() {
	account1 := NewBankAccount("dther")
	account1.Credit(0)
	account1.Debit(6000)
	fmt.Println(account1)
}
