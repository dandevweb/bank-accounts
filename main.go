package main

import "fmt"

type CurrentAccount struct {
	holder  string
	agency  int
	account int
	balance float64
}

func (account *CurrentAccount) Withdrawal(value float64) string {
	canWithdrawal := value <= account.balance && value > 0

	if canWithdrawal {
		account.balance -= value

		return "Saque realizado com sucesso."
	}

	return "Saldo insuficiente ou valor do saque inválido."

}

func main() {
	firstAccount := CurrentAccount{}
	firstAccount.holder = "Danilo Augusto"
	firstAccount.balance = 569.9

	fmt.Println(firstAccount.balance)

	fmt.Println(firstAccount.Withdrawal(200))

	fmt.Println(firstAccount.balance)

}
