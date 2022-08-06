package login

import (
	"fmt"
	"github.com/joleques/oleques-bank/account/src/domain"
)

type LoginService struct {
}

func (service LoginService) Create(account domain.Account, secret string) error {
	fmt.Println("gerar login")
	return nil
}
