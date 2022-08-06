package domain

type AccountRepository interface {
	Save(account *Account) error
	Remove(account *Account)
	Get(id string) *Account
	List() []*Account
}

type LoginService interface {
	Create(account Account, secret string) error
}
