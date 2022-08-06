package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeveRetornarErroQuandoIdVazio(t *testing.T) {
	account, err := NewAccount("", 10.0)
	assert.Nil(t, account)
	assert.Equal(t, "account: id invalid", err.Error())
}

func Test_DeveRetornarErroQuandoBalanceMenorQueZero(t *testing.T) {
	account, err := NewAccount("123", -10.0)
	assert.Nil(t, account)
	assert.Equal(t, "account: BalanceAccount invalid", err.Error())
}

func Test_DeveCriarContaComSucesso(t *testing.T) {
	account, err := NewAccount("123", 10.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, account.Balance())
	assert.Equal(t, "123", account.Id)
}

func Test_RealizarDepositoComSucesso(t *testing.T) {
	account, err := NewAccount("123", 10.0)
	assert.Nil(t, err)
	err = account.Deposit(&Amount{value: 5.55})
	assert.Nil(t, err)
	assert.Equal(t, 15.55, account.Balance())
}

func Test_RealizarSaqueComSucesso(t *testing.T) {
	account, err := NewAccount("123", 10.0)
	assert.Nil(t, err)
	err = account.Withdraw(&Amount{value: 5.55})
	assert.Nil(t, err)
	assert.Equal(t, 4.45, account.Balance())
}

func Test_NaoDeveRealizarSaqueQuandoValorMaiorQueSaldo(t *testing.T) {
	account, err := NewAccount("123", 10.0)
	assert.Nil(t, err)
	err = account.Withdraw(&Amount{value: 15.55})
	assert.Equal(t, "withdrawal not allowed, account without BalanceAccount", err.Error())
}
