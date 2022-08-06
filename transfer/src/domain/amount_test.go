package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeveRetornaErroQuandoValorMenorQueZero(t *testing.T) {
	amount, err := newAmount(-10.0)
	assert.Nil(t, amount)
	assert.Equal(t, "amount invalid", err.Error())
}

func Test_DeveCriarComSucesso(t *testing.T) {
	amount, err := newAmount(10.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, amount.value)
}
