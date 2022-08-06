package userCase

import (
	"errors"
	"github.com/joleques/oleques-bank/account/src/application/dto"
	"github.com/joleques/oleques-bank/account/src/domain"
)

type UpdateBalance struct {
	accountRepository domain.AccountRepository
}

func NewUpdateBalance(accountRepository domain.AccountRepository) (*UpdateBalance, error) {
	useCase := UpdateBalance{accountRepository: accountRepository}
	err := useCase.valid()
	if err != nil {
		return nil, err
	}
	return &useCase, nil
}

func (useCase UpdateBalance) valid() error {
	if useCase.accountRepository == nil {
		return errors.New("repository is required")
	}
	return nil
}

func (useCase UpdateBalance) Update(dto dto.BalanceDTO) error {
	account := useCase.accountRepository.Get(dto.ID)
	if account == nil {
		return errors.New("account invalid")
	}

	err := account.UpdateBalance(dto.Balance)
	if err != nil {
		return err
	}
	return useCase.accountRepository.Save(account)
}
