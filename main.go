package main

import (
	"bank/accounts"
	"fmt"
)

func main() {
	// firstHolder := customers.Holder{"Bruno", "123.456.789-10", "Desenvolvedor"}
	firstAccount := accounts.CurrentAccount{}
	firstAccount.Deposit(100)

	fmt.Println(firstAccount.GetBalance())

}
