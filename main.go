package main

import (
	"bank/accounts"
	"fmt"
)

func PayBillet(account verifyAccount, value float64) {
	account.Withdrawal(value)
}

type verifyAccount interface {
	Withdrawal(value float64) string
}

func main() {
	// firstHolder := customers.Holder{"Bruno", "123.456.789-10", "Desenvolvedor"}
	firstAccount := accounts.SavingsAccount{}
	firstAccount.Deposit(100)
	PayBillet(&firstAccount, 60)
	fmt.Println(firstAccount)

	current := accounts.CurrentAccount{}
	current.Deposit(1000)

	PayBillet(&current, 400)
	fmt.Println(current)

}
