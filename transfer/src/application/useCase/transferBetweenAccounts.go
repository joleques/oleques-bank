package useCase

import (
	"errors"
	"github.com/joleques/oleques-bank/transfer/src/application/dto"
	"github.com/joleques/oleques-bank/transfer/src/domain"
)

type TransferBetweenAccounts struct {
	accountService domain.AccountService
	repository     domain.TransferRepository
}

func (useCase TransferBetweenAccounts) valid() error {
	if useCase.accountService == nil {
		return errors.New("accountService is required")
	}

	if useCase.repository == nil {
		return errors.New("repository is required")
	}
	return nil
}

func (useCase TransferBetweenAccounts) Transfer(dto dto.TransferDTO) error {
	transfer, err := useCase.executeTransfer(dto)
	if err != nil {
		return err
	}
	err = useCase.updateAccounts(*transfer)
	if err != nil {
		return err
	}
	return useCase.repository.Save(transfer)
}

func (useCase TransferBetweenAccounts) executeTransfer(dto dto.TransferDTO) (*domain.Transfer, error) {
	accountOrigin, err := useCase.buildCount(dto.AccountOriginId)
	if err != nil {
		return nil, err
	}
	accountDestination, err := useCase.buildCount(dto.AccountDestinationId)
	if err != nil {
		return nil, err
	}
	return domain.NewTransfer(accountOrigin, accountDestination, dto.Amount)

}
func (useCase TransferBetweenAccounts) updateAccounts(transfer domain.Transfer) error {
	err := useCase.accountService.UpdateAccount(transfer.Origin())
	if err != nil {
		return err
	}
	return useCase.accountService.UpdateAccount(transfer.Destination())
}

func (useCase TransferBetweenAccounts) buildCount(id string) (*domain.Account, error) {
	balanceOrigin, err := useCase.accountService.GetBalance(id)
	if err != nil {
		return nil, err
	}
	return domain.NewAccount(id, balanceOrigin)
}

func NewTransferBetweenAccounts(accountService domain.AccountService, repository domain.TransferRepository) (*TransferBetweenAccounts, error) {
	useCase := TransferBetweenAccounts{accountService: accountService, repository: repository}
	err := useCase.valid()
	if err != nil {
		return nil, err
	}
	return &useCase, nil
}
