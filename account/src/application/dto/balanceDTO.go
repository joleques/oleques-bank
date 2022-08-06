package dto

import "net/http"

type BalanceDTO struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

func (r2 BalanceDTO) Bind(r *http.Request) error {
	return nil
}
