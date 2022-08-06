package dto

import "net/http"

type TransferDTO struct {
	AccountOriginId      string  `json:"account_origin_id"`
	AccountDestinationId string  `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
}

func (r2 TransferDTO) Bind(r *http.Request) error {
	return nil
}
