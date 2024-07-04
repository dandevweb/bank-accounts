package accounts

type CurrentAccount struct {
	Holder  string
	Agency  int
	Account int
	Balance float64
}

func (account *CurrentAccount) Withdrawal(value float64) string {
	canWithdrawal := value <= account.Balance && value > 0

	if canWithdrawal {
		account.Balance -= value

		return "Saque realizado com sucesso."
	}

	return "Saldo insuficiente ou valor do saque inválido."

}

func (account *CurrentAccount) Deposit(value float64) (string, float64) {
	canDeposit := value > 0

	if canDeposit {
		account.Balance += value

		return "Depósito realizado com sucesso.", account.Balance
	}

	return "Erro! O valor do depósito precisa ser positivo", account.Balance
}

func (account *CurrentAccount) Transfer(value float64, destinationAccount *CurrentAccount) bool {
	canTransfer := account.Balance > value && value > 0

	if canTransfer {
		account.Balance -= value
		destinationAccount.Deposit(value)

		return true
	}

	return false
}
