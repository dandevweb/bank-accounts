package accounts

import "bank/customers"

type SavingsAccount struct {
	Holder                     customers.Holder
	Agency, Account, Operation int
	balance                    float64
}

func (account *SavingsAccount) Withdrawal(value float64) string {
	canWithdrawal := value <= account.balance && value > 0

	if canWithdrawal {
		account.balance -= value

		return "Saque realizado com sucesso."
	}

	return "Saldo insuficiente ou valor do saque inválido."

}

func (account *SavingsAccount) Deposit(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		account.balance += value

		return "Depósito realizado com sucesso.", account.balance
	}

	return "Erro! O valor do depósito precisa ser positivo", account.balance
}

func (account *SavingsAccount) Transfer(value float64, destinationAccount *SavingsAccount) bool {
	canTransfer := account.balance > value && value > 0

	if canTransfer {
		account.balance -= value
		destinationAccount.Deposit(value)

		return true
	}

	return false
}

func (account *SavingsAccount) GetBalance() float64 {
	return account.balance
}
