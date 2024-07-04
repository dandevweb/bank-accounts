package main

import (
	"bank/accounts"
	"fmt"
)

func main() {
	firstAccount := accounts.CurrentAccount{}
	firstAccount.Holder = "Danilo Augusto"
	firstAccount.Balance = 300

	secondAccount := accounts.CurrentAccount{Holder: "Solange", Balance: 100}

	statusTransfer := firstAccount.Transfer(200, &secondAccount)

	fmt.Println(statusTransfer)

	fmt.Println("First Account:", firstAccount.Balance)
	fmt.Println("Second Account:", secondAccount)

	// fmt.Println("First Account:", firstAccount.Withdrawal(200))

	// fmt.Println("First Account:", firstAccount.Balance)

	// status, Balance := firstAccount.Deposit(1200)

	// fmt.Println(status, Balance)

}
