package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NaoDeveCriarTransferenciaQuandoContaOriginNula(t *testing.T) {
	accountDestination, err := NewAccount("12", 25)
	amount := 10.0
	assert.Nil(t, err)
	transfer, err := NewTransfer(nil, accountDestination, amount)
	assert.Nil(t, transfer)
	assert.Equal(t, "accountOrigin invalid", err.Error())
}

func Test_NaoDeveCriarTransferenciaQuandoContaDestinoNula(t *testing.T) {
	accountOrigin, err := NewAccount("12", 25)
	amount := 10.0
	assert.Nil(t, err)
	transfer, err := NewTransfer(accountOrigin, nil, amount)
	assert.Nil(t, transfer)
	assert.Equal(t, "accountDestination invalid", err.Error())
}

func Test_NaoDeveCriarTransferenciaQuandoContaValorNegativo(t *testing.T) {
	accountOrigin, err := NewAccount("12", 25)
	amount := -10.0
	assert.Nil(t, err)
	transfer, err := NewTransfer(accountOrigin, nil, amount)
	assert.Nil(t, transfer)
	assert.Equal(t, "amount invalid", err.Error())
}

func Test_DeveCriartransferenciaComSucesso(t *testing.T) {
	accountOrigin, _ := NewAccount("12", 25.58)
	accountDestination, _ := NewAccount("12", 75.89)
	amount := 7.56
	transfer, err := NewTransfer(accountOrigin, accountDestination, amount)
	assert.Nil(t, err)
	assert.NotNil(t, transfer)
}

func Test_NaoDeveCriartransferenciaQuandoSaldoInsuficiente(t *testing.T) {
	accountOrigin, _ := NewAccount("12", 25.58)
	accountDestination, _ := NewAccount("12", 75.89)
	amount := 107.56
	transfer, err := NewTransfer(accountOrigin, accountDestination, amount)
	assert.Nil(t, transfer)
	assert.Equal(t, "withdrawal not allowed, account without BalanceAccount", err.Error())
}
