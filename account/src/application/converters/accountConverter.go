package converters

import (
	"github.com/joleques/oleques-bank/account/src/application/dto"
	"github.com/joleques/oleques-bank/account/src/domain"
)

type AccountConverter struct {
}

func (converter AccountConverter) Convert(dto dto.AccountDTO) (*domain.Account, error) {
	account, err := domain.NewAccount(dto.Name, dto.CPF)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (converter AccountConverter) ConvertToDTO(account domain.Account) dto.AccountResponseDTO {
	dto := dto.AccountResponseDTO{ID: account.Id, Name: account.Name, CPF: account.Cpf, Balance: account.Balance()}
	return dto
}
