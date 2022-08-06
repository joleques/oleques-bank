package domain

import "errors"

type Amount struct {
	value float64
}

func (amount Amount) valid() error {
	if amount.value < 0 {
		return errors.New("amount invalid")
	}
	return nil
}

func newAmount(value float64) (*Amount, error) {
	amount := Amount{value: value}
	err := amount.valid()
	if err != nil {
		return nil, err
	}
	return &amount, nil
}
