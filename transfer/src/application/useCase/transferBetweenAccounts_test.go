package useCase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/joleques/oleques-bank/transfer/pkg/domain/mock"
	"github.com/joleques/oleques-bank/transfer/src/application/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NaoDeveExecutarTransferenciaQuandoAccountServiceIsNil(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repository := mock.NewMockTransferRepository(controller)
	useCase, err := NewTransferBetweenAccounts(nil, repository)
	assert.Nil(t, useCase)
	assert.Equal(t, "accountService is required", err.Error())
}

func Test_NaoDeveExecutarTransferenciaQuandoRepositorioIsNil(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockAccountService(controller)
	useCase, err := NewTransferBetweenAccounts(service, nil)
	assert.Nil(t, useCase)
	assert.Equal(t, "repository is required", err.Error())
}

func Test_NaoDeveExecutarTransferenciaQuandoContaOriginNaoExiste(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	accountOrigin := "123"
	accountDestination := "456"
	msgErr := "Origin invalid"
	transferDTO := dto.TransferDTO{AccountOriginId: accountOrigin, AccountDestinationId: accountDestination, Amount: 45.58}

	service := mock.NewMockAccountService(controller)
	repository := mock.NewMockTransferRepository(controller)

	service.EXPECT().GetBalance(accountOrigin).Return(0.0, errors.New(msgErr))

	useCase, err := NewTransferBetweenAccounts(service, repository)
	assert.Nil(t, err)

	err = useCase.Transfer(transferDTO)
	assert.Equal(t, msgErr, err.Error())
}
func Test_NaoDeveExecutarTransferenciaQuandoContaDestinoNaoExiste(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	accountOrigin := "123"
	accountDestination := "456"
	msgErr := "Destination invalid"
	transferDTO := dto.TransferDTO{AccountOriginId: accountOrigin, AccountDestinationId: accountDestination, Amount: 45.58}

	service := mock.NewMockAccountService(controller)
	repository := mock.NewMockTransferRepository(controller)

	service.EXPECT().GetBalance(accountOrigin).Return(10.0, nil)
	service.EXPECT().GetBalance(accountDestination).Return(10.0, errors.New(msgErr))

	useCase, err := NewTransferBetweenAccounts(service, repository)
	assert.Nil(t, err)

	err = useCase.Transfer(transferDTO)
	assert.Equal(t, msgErr, err.Error())
}

func Test_NaoDeveExecutarTransferenciaQuandoContaSaldoInsuficiente(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	accountOrigin := "123"
	accountDestination := "456"
	transferDTO := dto.TransferDTO{AccountOriginId: accountOrigin, AccountDestinationId: accountDestination, Amount: 45.58}

	service := mock.NewMockAccountService(controller)
	repository := mock.NewMockTransferRepository(controller)

	service.EXPECT().GetBalance(accountOrigin).Return(10.0, nil)
	service.EXPECT().GetBalance(accountDestination).Return(10.0, nil)

	useCase, err := NewTransferBetweenAccounts(service, repository)
	assert.Nil(t, err)

	err = useCase.Transfer(transferDTO)
	assert.Equal(t, "withdrawal not allowed, account without balance", err.Error())
}

func Test_NaoDeveExecutarTransferenciaQuandoErroAoAtualizarContaOrigin(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	accountOrigin := "123"
	accountDestination := "456"
	msgErr := "error update"
	transferDTO := dto.TransferDTO{AccountOriginId: accountOrigin, AccountDestinationId: accountDestination, Amount: 5.58}

	service := mock.NewMockAccountService(controller)
	repository := mock.NewMockTransferRepository(controller)

	service.EXPECT().GetBalance(accountOrigin).Return(10.0, nil)
	service.EXPECT().GetBalance(accountDestination).Return(10.0, nil)
	service.EXPECT().UpdateAccount(gomock.Any()).Return(errors.New(msgErr))

	useCase, err := NewTransferBetweenAccounts(service, repository)
	assert.Nil(t, err)

	err = useCase.Transfer(transferDTO)
	assert.Equal(t, msgErr, err.Error())
}

func Test_NaoDeveExecutarTransferenciaQuandoErroAoSalvarTransferencia(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	accountOrigin := "123"
	accountDestination := "456"
	msgErr := "error save"
	transferDTO := dto.TransferDTO{AccountOriginId: accountOrigin, AccountDestinationId: accountDestination, Amount: 5.58}

	service := mock.NewMockAccountService(controller)
	repository := mock.NewMockTransferRepository(controller)

	service.EXPECT().GetBalance(accountOrigin).Return(10.0, nil)
	service.EXPECT().GetBalance(accountDestination).Return(10.0, nil)
	service.EXPECT().UpdateAccount(gomock.Any()).Return(nil)
	service.EXPECT().UpdateAccount(gomock.Any()).Return(nil)
	repository.EXPECT().Save(gomock.Any()).Return(errors.New(msgErr))

	useCase, err := NewTransferBetweenAccounts(service, repository)
	assert.Nil(t, err)

	err = useCase.Transfer(transferDTO)
	assert.Equal(t, msgErr, err.Error())
}

func Test_DeveExecutarTransferenciaComSucesso(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	accountOrigin := "123"
	accountDestination := "456"
	transferDTO := dto.TransferDTO{AccountOriginId: accountOrigin, AccountDestinationId: accountDestination, Amount: 5.58}

	service := mock.NewMockAccountService(controller)
	repository := mock.NewMockTransferRepository(controller)

	service.EXPECT().GetBalance(accountOrigin).Return(10.0, nil)
	service.EXPECT().GetBalance(accountDestination).Return(10.0, nil)
	service.EXPECT().UpdateAccount(gomock.Any()).Return(nil)
	service.EXPECT().UpdateAccount(gomock.Any()).Return(nil)
	repository.EXPECT().Save(gomock.Any()).Return(nil)

	useCase, err := NewTransferBetweenAccounts(service, repository)
	assert.Nil(t, err)

	err = useCase.Transfer(transferDTO)
	assert.Nil(t, err)
}
