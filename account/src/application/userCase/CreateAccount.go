package userCase

import (
	"errors"
	"github.com/joleques/oleques-bank/account/src/application/converters"
	"github.com/joleques/oleques-bank/account/src/application/dto"
	"github.com/joleques/oleques-bank/account/src/domain"
)

type CreateAccount struct {
	converter         *converters.AccountConverter
	accountRepository domain.AccountRepository
	loginService      domain.LoginService
}

func NewCreateAccount(converter *converters.AccountConverter, accountRepository domain.AccountRepository, loginService domain.LoginService) (*CreateAccount, error) {
	useCase := CreateAccount{converter: converter, accountRepository: accountRepository, loginService: loginService}
	err := useCase.valid()
	if err != nil {
		return nil, err
	}
	return &useCase, nil
}

func (useCase CreateAccount) Create(accountDTO dto.AccountDTO) error {
	account, err := useCase.converter.Convert(accountDTO)
	if err != nil {
		return err
	}
	err = useCase.accountRepository.Save(account)
	if err != nil {
		return err
	}
	err = useCase.loginService.Create(*account, accountDTO.Secret)
	if err != nil {
		useCase.accountRepository.Remove(account)
		return err
	}
	return nil
}

func (useCase CreateAccount) valid() error {
	if useCase.converter == nil {
		return errors.New("converter is required")
	}
	if useCase.accountRepository == nil {
		return errors.New("accountRepository is required")
	}
	if useCase.loginService == nil {
		return errors.New("loginService is required")
	}
	return nil
}
