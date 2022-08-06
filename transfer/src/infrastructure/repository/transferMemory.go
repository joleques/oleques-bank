package repository

import "github.com/joleques/oleques-bank/transfer/src/domain"

type TransferMemory struct {
}

var transfers map[string][]*domain.Transfer = map[string][]*domain.Transfer{}

func (base TransferMemory) Save(transfer *domain.Transfer) error {
	element := transfers[transfer.Origin().Id]
	if element == nil {
		element = []*domain.Transfer{}
	}
	element = append(element, transfer)
	transfers[transfer.Origin().Id] = element
	return nil
}

func (base TransferMemory) List(id string) []*domain.Transfer {
	element := transfers[id]
	if element == nil {
		element = []*domain.Transfer{}
	}
	return element
}
