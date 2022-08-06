package userCase

import (
	"errors"
	"github.com/joleques/oleques-bank/account/src/application/converters"
	"github.com/joleques/oleques-bank/account/src/application/dto"
	"github.com/joleques/oleques-bank/account/src/domain"
)

type GetAccount struct {
	converter         *converters.AccountConverter
	accountRepository domain.AccountRepository
}

func (useCase GetAccount) valid() error {
	if useCase.converter == nil {
		return errors.New("converter is required")
	}
	if useCase.accountRepository == nil {
		return errors.New("repository is required")
	}
	return nil
}

func (useCase GetAccount) Get(id string) (*dto.AccountResponseDTO, error) {
	if id == "" {
		return nil, errors.New("Id empty")
	}
	account := useCase.accountRepository.Get(id)
	if account == nil {
		return nil, errors.New("Account does not exist")
	}
	toDTO := useCase.converter.ConvertToDTO(*account)

	return &toDTO, nil
}

func NewGetAccount(converter *converters.AccountConverter, accountRepository domain.AccountRepository) (*GetAccount, error) {
	useCase := GetAccount{converter: converter, accountRepository: accountRepository}
	err := useCase.valid()
	if err != nil {
		return nil, err
	}
	return &useCase, nil

}
