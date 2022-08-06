package userCase

import (
	"github.com/golang/mock/gomock"
	"github.com/joleques/oleques-bank/account/pkg/domain/mock"
	dto2 "github.com/joleques/oleques-bank/account/src/application/dto"
	"github.com/joleques/oleques-bank/account/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeveDarErroQuandoRepositorioNulo(t *testing.T) {
	useCase, err := NewUpdateBalance(nil)
	assert.Nil(t, useCase)
	assert.Equal(t, "repository is required", err.Error())
}

func Test_DeveRetornarErroQuandoNaoEncontraConta(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	id := "123"
	valor := 10.0
	dto := dto2.BalanceDTO{ID: id, Balance: valor}

	repository := mock.NewMockAccountRepository(controller)
	repository.EXPECT().Get(id).Return(nil)

	useCase, err := NewUpdateBalance(repository)
	assert.Nil(t, err)
	err = useCase.Update(dto)

	assert.Equal(t, "account invalid", err.Error())
}

func Test_DeveAtualizarConta(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	id := "123"
	valor := 10.0
	cpf := "850.444.710-38"
	joao := "Jão Ubaldo"
	dto := dto2.BalanceDTO{ID: id, Balance: valor}
	account, _ := domain.NewAccount(joao, cpf)

	repository := mock.NewMockAccountRepository(controller)
	repository.EXPECT().Get(id).Return(account)
	repository.EXPECT().Save(account)

	useCase, err := NewUpdateBalance(repository)
	assert.Nil(t, err)
	err = useCase.Update(dto)
	assert.Nil(t, err)
	assert.Equal(t, valor, account.Balance())
}
