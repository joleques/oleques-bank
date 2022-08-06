package userCase

import (
	"github.com/golang/mock/gomock"
	"github.com/joleques/oleques-bank/account/pkg/domain/mock"
	"github.com/joleques/oleques-bank/account/src/application/converters"
	"github.com/joleques/oleques-bank/account/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeveRetornarErroQuandoNaoTemonverter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repository := mock.NewMockAccountRepository(controller)

	useCase, err := NewGetAccount(nil, repository)
	assert.Nil(t, useCase)
	assert.Equal(t, "converter is required", err.Error())
}

func Test_DeveRetornarErroQuandoNaoTemRepositorio(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	converter := &converters.AccountConverter{}

	useCase, err := NewGetAccount(converter, nil)
	assert.Nil(t, useCase)
	assert.Equal(t, "repository is required", err.Error())
}

func Test_DeveRetornarErroQuandoAccountNaoExiste(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	id := "798-test"

	converter := &converters.AccountConverter{}
	repository := mock.NewMockAccountRepository(controller)

	repository.EXPECT().Get(id).Return(nil)

	useCase, err := NewGetAccount(converter, repository)
	assert.Nil(t, err)

	responseDTOs, err := useCase.Get(id)
	assert.Nil(t, responseDTOs)
	assert.Equal(t, "Account does not exist", err.Error())
}

func Test_DeveRetornarErroQuandoIdVazio(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	id := ""

	converter := &converters.AccountConverter{}
	repository := mock.NewMockAccountRepository(controller)

	useCase, err := NewGetAccount(converter, repository)
	assert.Nil(t, err)
	responseDTOs, err := useCase.Get(id)
	assert.Nil(t, responseDTOs)
	assert.Equal(t, "Id empty", err.Error())
}

func Test_DeveRetornarDTO(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	id := "123-test"
	joao := "Joao,"
	cpfJoao := "94785698574"
	account, _ := domain.NewAccount(joao, cpfJoao)

	converter := &converters.AccountConverter{}
	repository := mock.NewMockAccountRepository(controller)

	repository.EXPECT().Get(id).Return(account)

	useCase, err := NewGetAccount(converter, repository)
	assert.Nil(t, err)

	responseDTOs, err := useCase.Get(id)
	assert.Nil(t, err)

	assert.NotEmpty(t, responseDTOs.ID)
	assert.Equal(t, joao, responseDTOs.Name)
	assert.Equal(t, cpfJoao, responseDTOs.CPF)
}
