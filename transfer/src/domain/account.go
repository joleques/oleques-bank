package domain

import (
	"errors"
	"github.com/shopspring/decimal"
)

type Account struct {
	Id             string
	BalanceAccount *Amount `json:"balance"`
}

func (account Account) Balance() float64 {
	return account.BalanceAccount.value
}

func (account Account) valid() error {
	if account.Id == "" {
		return errors.New("account: id invalid")
	}
	if account.BalanceAccount == nil {
		return errors.New("account: BalanceAccount invalid")
	}
	return nil
}

func (account *Account) Withdraw(amount *Amount) error {
	v1 := decimal.NewFromFloat(amount.value)
	v2 := decimal.NewFromFloat(account.BalanceAccount.value)
	newValue := v2.Sub(v1)
	return account.updateBalance(newValue)
}

func (account *Account) Deposit(amount *Amount) error {
	v1 := decimal.NewFromFloat(amount.value)
	v2 := decimal.NewFromFloat(account.BalanceAccount.value)
	newValue := v2.Add(v1)
	return account.updateBalance(newValue)

}
func (account *Account) updateBalance(newValue decimal.Decimal) error {
	floatValue, _ := newValue.Float64()
	balance, _ := newAmount(floatValue)
	account.BalanceAccount = balance
	err := account.valid()
	if err != nil {
		return errors.New("withdrawal not allowed, account without BalanceAccount")
	}
	return nil
}

func NewAccount(id string, balance float64) (*Account, error) {
	value, err := newAmount(balance)
	if err != nil {
		return nil, errors.New("account: BalanceAccount invalid")
	}
	account := Account{Id: id, BalanceAccount: value}
	err = account.valid()
	if err != nil {
		return nil, err
	}
	return &account, nil
}
