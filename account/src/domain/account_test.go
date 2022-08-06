package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeveCriarContaQuandoNaoPassarNome(t *testing.T) {
	cpf := "850.444.710-38"
	account, err := NewAccount("", cpf)
	assert.Nil(t, account)
	assert.Equal(t, "name is required", err.Error())
}

func Test_DeveCriarContaQuandoNaoPassarCpf(t *testing.T) {
	joao := "Jão Ubaldo"
	account, err := NewAccount(joao, "")
	assert.Nil(t, account)
	assert.Equal(t, "cpf is required", err.Error())
}

func Test_DeveAtualizarSaldoComSucesso(t *testing.T) {
	cpf := "850.444.710-38"
	joao := "Jão Ubaldo"
	account, err := NewAccount(joao, cpf)
	assert.Nil(t, err)

	err = account.UpdateBalance(45.60)
	assert.Nil(t, err)
	assert.Equal(t, 45.60, account.Balance())

	err = account.UpdateBalance(35.60)
	assert.Nil(t, err)
	assert.Equal(t, 35.60, account.Balance())
}

func Test_NaoDeveAtualizarSaldoQuandoSaldoNovoForNegativo(t *testing.T) {
	cpf := "850.444.710-38"
	joao := "Jão Ubaldo"
	account, err := NewAccount(joao, cpf)
	assert.Nil(t, err)
	err = account.UpdateBalance(10.60)
	assert.Nil(t, err)
	assert.Equal(t, 10.60, account.Balance())

	err = account.UpdateBalance(-20.90)
	assert.Equal(t, "Balance invalid: -20.90", err.Error())
}
