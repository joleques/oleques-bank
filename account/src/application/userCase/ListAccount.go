package userCase

import (
	"errors"
	"github.com/joleques/oleques-bank/account/src/application/converters"
	"github.com/joleques/oleques-bank/account/src/application/dto"
	"github.com/joleques/oleques-bank/account/src/domain"
)

type ListAccount struct {
	converter         *converters.AccountConverter
	accountRepository domain.AccountRepository
}

func NewListAccount(converter *converters.AccountConverter, accountRepository domain.AccountRepository) (*ListAccount, error) {
	useCase := ListAccount{converter: converter, accountRepository: accountRepository}
	err := useCase.valid()
	if err != nil {
		return nil, err
	}
	return &useCase, nil

}

func (useCase ListAccount) List() []dto.AccountResponseDTO {
	list := useCase.accountRepository.List()
	var DTOs []dto.AccountResponseDTO
	for _, account := range list {
		toDTO := useCase.converter.ConvertToDTO(*account)
		DTOs = append(DTOs, toDTO)
	}
	return DTOs
}

func (useCase ListAccount) valid() error {
	if useCase.converter == nil {
		return errors.New("converter is required")
	}
	if useCase.accountRepository == nil {
		return errors.New("repository is required")
	}
	return nil
}
