package banking

type Account struct {
	Owner   string
	Balance int
}

func getNewAccount() *Account {
	account := Account{Owner: "주인장", Balance: 30000}
	return &account
}
