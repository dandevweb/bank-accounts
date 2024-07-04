package accounts

import "bank/customers"

type CurrentAccount struct {
	Holder  customers.Holder
	Agency  int
	Account int
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

func (account *CurrentAccount) Deposit(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		account.balance += value

		return "Depósito realizado com sucesso.", account.balance
	}

	return "Erro! O valor do depósito precisa ser positivo", account.balance
}

func (account *CurrentAccount) Transfer(value float64, destinationAccount *CurrentAccount) bool {
	canTransfer := account.balance > value && value > 0

	if canTransfer {
		account.balance -= value
		destinationAccount.Deposit(value)

		return true
	}

	return false
}

func (account *CurrentAccount) GetBalance() float64 {
	return account.balance
}
