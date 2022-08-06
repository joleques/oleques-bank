package dto

import "net/http"

type AccountDTO struct {
	Secret string `json:"secret"`
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
}

func (r2 AccountDTO) Bind(r *http.Request) error {
	return nil
}

type AccountResponseDTO struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	CPF     string  `json:"cpf"`
	Balance float64 `json:"balance"`
}
