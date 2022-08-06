package domain

type AccountService interface {
	GetBalance(id string) (float64, error)
	UpdateAccount(account Account) error
}

type TransferRepository interface {
	Save(transfer *Transfer) error
	List(id string) []*Transfer
}
