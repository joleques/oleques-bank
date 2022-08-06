package userCase

import (
	"github.com/golang/mock/gomock"
	"github.com/joleques/oleques-bank/account/pkg/domain/mock"
	"github.com/joleques/oleques-bank/account/src/application/converters"
	"github.com/joleques/oleques-bank/account/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeveRetornarErroQuandoNaoFoiPassadoConverter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repository := mock.NewMockAccountRepository(controller)

	useCase, err := NewListAccount(nil, repository)
	assert.Nil(t, useCase)
	assert.Equal(t, "converter is required", err.Error())
}

func Test_DeveRetornarErroQuandoNaoFoiPassadoRepositorio(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	converter := &converters.AccountConverter{}

	useCase, err := NewListAccount(converter, nil)
	assert.Nil(t, useCase)
	assert.Equal(t, "repository is required", err.Error())
}

func Test_DeveRetornarListaDTOs(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	joao := "Joao,"
	cpfJoao := "94785698574"
	paulo := "Paulo,"
	cpfPaulo := "12365478985"
	account1, _ := domain.NewAccount(joao, cpfJoao)
	account2, _ := domain.NewAccount(paulo, cpfPaulo)

	var accounts []*domain.Account
	accounts = append(accounts, account1)
	accounts = append(accounts, account2)
	converter := &converters.AccountConverter{}
	repository := mock.NewMockAccountRepository(controller)

	repository.EXPECT().List().Return(accounts)

	useCase, err := NewListAccount(converter, repository)
	assert.Nil(t, err)

	responseDTOs := useCase.List()

	assert.Equal(t, 2, len(responseDTOs))

	dto1 := responseDTOs[0]
	assert.NotEmpty(t, dto1.ID)
	assert.Equal(t, joao, dto1.Name)
	assert.Equal(t, cpfJoao, dto1.CPF)

	dto2 := responseDTOs[1]
	assert.NotEmpty(t, dto2.ID)
	assert.Equal(t, paulo, dto2.Name)
	assert.Equal(t, cpfPaulo, dto2.CPF)
}
