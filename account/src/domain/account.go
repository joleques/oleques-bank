package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Account struct {
	Id         string
	Name       string
	Cpf        string
	balance    Balance
	created_at string
}

func (account Account) valid() error {
	if account.Name == "" {
		return errors.New("name is required")
	}

	if account.Cpf == "" {
		return errors.New("cpf is required")
	}
	return nil
}

func (account Account) Balance() float64 {
	return account.balance.Value
}

func (account *Account) UpdateBalance(value float64) error {
	balance, err := NewBalance(value)
	if err != nil {
		return err
	}
	account.balance = *balance
	return nil
}

type Balance struct {
	Value float64
}

func (balance Balance) valid() error {
	if balance.Value < 0 {
		return errors.New(fmt.Sprintf("Balance invalid: %.2f", balance.Value))
	}
	return nil
}

func NewAccount(name string, cpf string) (*Account, error) {
	account := Account{Id: uuid.New().String(), Name: name, Cpf: cpf, balance: Balance{}, created_at: time.Now().UTC().Format(time.RFC3339)}
	err := account.valid()
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func NewBalance(value float64) (*Balance, error) {
	newBalance := Balance{Value: value}
	err := newBalance.valid()
	if err != nil {
		return nil, err
	}
	return &newBalance, nil
}
