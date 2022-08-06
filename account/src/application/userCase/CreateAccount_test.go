package userCase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/joleques/oleques-bank/account/pkg/domain/mock"
	"github.com/joleques/oleques-bank/account/src/application/converters"
	dto2 "github.com/joleques/oleques-bank/account/src/application/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeveRetornarErroQuandoCriarUseCaseSemConverter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := mock.NewMockAccountRepository(controller)
	service := mock.NewMockLoginService(controller)

	useCase, err := NewCreateAccount(nil, repo, service)
	assert.Nil(t, useCase)
	assert.Equal(t, "converter is required", err.Error())
}

func Test_DeveRetornarErroQuandoCriarUseCaseSemLoginService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := mock.NewMockAccountRepository(controller)

	useCase, err := NewCreateAccount(&converters.AccountConverter{}, repo, nil)
	assert.Nil(t, useCase)
	assert.Equal(t, "loginService is required", err.Error())
}

func Test_DeveRetornarErroQuandoCriarUseCaseSemRepositorio(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	service := mock.NewMockLoginService(controller)

	useCase, err := NewCreateAccount(&converters.AccountConverter{}, nil, service)
	assert.Nil(t, useCase)
	assert.Equal(t, "accountRepository is required", err.Error())
}

func Test_DeveRetornarErroQuandoConversaoInvalida(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	joao := "João Paulo"
	dto := dto2.AccountDTO{Name: joao}

	repo := mock.NewMockAccountRepository(controller)
	service := mock.NewMockLoginService(controller)
	converter := &converters.AccountConverter{}

	useCase, err := NewCreateAccount(converter, repo, service)
	assert.Nil(t, err)

	err = useCase.Create(dto)
	assert.Equal(t, "cpf is required", err.Error())
}

func Test_DeveRetornarErroQuandoRepoSalvarComErro(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	cpf := "850.444.710-38"
	joao := "João Paulo"
	msgErr := "error test repo"
	dto := dto2.AccountDTO{Name: joao, CPF: cpf}

	repo := mock.NewMockAccountRepository(controller)
	service := mock.NewMockLoginService(controller)
	converter := &converters.AccountConverter{}

	repo.EXPECT().Save(gomock.Any()).Return(errors.New(msgErr))

	useCase, err := NewCreateAccount(converter, repo, service)
	assert.Nil(t, err)

	err = useCase.Create(dto)
	assert.Equal(t, msgErr, err.Error())
}

func Test_DeveRetornarErroQuandoGerarLoginComErro(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	cpf := "850.444.710-38"
	joao := "João Paulo"
	msgErr := "error test login"
	secret := "secret_123"
	dto := dto2.AccountDTO{Name: joao, CPF: cpf, Secret: secret}

	repo := mock.NewMockAccountRepository(controller)
	service := mock.NewMockLoginService(controller)
	converter := &converters.AccountConverter{}

	repo.EXPECT().Save(gomock.Any()).Return(nil)
	service.EXPECT().Create(gomock.Any(), secret).Return(errors.New(msgErr))
	repo.EXPECT().Remove(gomock.Any())

	useCase, err := NewCreateAccount(converter, repo, service)
	assert.Nil(t, err)

	err = useCase.Create(dto)
	assert.Equal(t, msgErr, err.Error())
}

func Test_DeveCriarContaComSucesso(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	cpf := "850.444.710-38"
	joao := "João Paulo"
	secret := "secret_123"
	dto := dto2.AccountDTO{Name: joao, CPF: cpf, Secret: secret}

	repo := mock.NewMockAccountRepository(controller)
	service := mock.NewMockLoginService(controller)
	converter := &converters.AccountConverter{}

	repo.EXPECT().Save(gomock.Any()).Return(nil)
	service.EXPECT().Create(gomock.Any(), secret).Return(nil)

	useCase, err := NewCreateAccount(converter, repo, service)
	assert.Nil(t, err)

	err = useCase.Create(dto)
	assert.Nil(t, err)
}
