package scalegraph

import (
	"errors"
	"fmt"
)

type vault struct {
	wallets map[[5]uint32]*wallet
}

func NewVault() vault {
	return vault{
		wallets: make(map[[5]uint32]*wallet),
	}
}

func (v *vault) Add(ID [5]uint32) error {
	_, exists := v.wallets[ID]
	if exists {
		return errors.New(fmt.Sprintf("wallet with id %+v already exists", ID))
	}
	newWallet := NewWallet(ID)
	v.wallets[ID] = &newWallet
	return nil
}
