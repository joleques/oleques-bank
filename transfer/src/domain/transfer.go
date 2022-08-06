package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Transfer struct {
	Id                 string
	accountOrigin      *Account
	accountDestination *Account
	amount             *Amount
	createAt           time.Time
}

func (transfer Transfer) Origin() Account {
	return *transfer.accountOrigin
}

func (transfer Transfer) Destination() Account {
	return *transfer.accountDestination
}

func (transfer *Transfer) valid() error {
	if transfer.accountOrigin == nil {
		return errors.New("accountOrigin invalid")
	}
	if transfer.accountDestination == nil {
		return errors.New("accountDestination invalid")
	}
	if transfer.amount == nil {
		return errors.New("amount invalid")
	}

	err := transfer.accountOrigin.Withdraw(transfer.amount)
	if err != nil {
		return err
	}
	return transfer.accountDestination.Deposit(transfer.amount)
}

func NewTransfer(accountOrigin *Account, accountDestination *Account, amount float64) (*Transfer, error) {
	amountObj, err := newAmount(amount)
	if err != nil {
		return nil, err
	}
	transfer := Transfer{Id: uuid.New().String(), accountDestination: accountDestination, accountOrigin: accountOrigin, amount: amountObj, createAt: time.Now()}
	err = transfer.valid()
	if err != nil {
		return nil, err
	}
	return &transfer, nil
}
