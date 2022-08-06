package repository

import (
	"github.com/joleques/oleques-bank/account/src/domain"
)

type AccountMemory struct {
}

var accounts map[string]*domain.Account = map[string]*domain.Account{}

func (repository AccountMemory) Save(account *domain.Account) error {
	key := account.Cpf
	accountBase := accounts[key]
	if accountBase == nil {
		accounts[key] = account
		return nil
	}
	accounts[key].Name = account.Name
	accounts[key].UpdateBalance(account.Balance())
	return nil
}

func (repository AccountMemory) Remove(account *domain.Account) {
	delete(accounts, account.Id)
}

func (repository AccountMemory) Get(id string) *domain.Account {
	for _, account := range repository.List() {
		if account.Id == id {
			return account
		}
	}
	return nil
}
func (repository AccountMemory) List() []*domain.Account {
	var accountsReturn []*domain.Account
	for _, account := range accounts {
		accountsReturn = append(accountsReturn, account)
	}
	return accountsReturn
}
