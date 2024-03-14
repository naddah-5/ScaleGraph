package scalegraph

import (
	"errors"
	"fmt"
)

type vault struct {
	slot map[[5]uint32]wallet
}

func NewVault(id [5]uint32) vault {
	return vault{
		slot: make(map[[5]uint32]wallet),
	}
}

func (v *vault) Add(id [5]uint32) error {
	_, exists := v.slot[id]
	if exists {
		return errors.New(fmt.Sprintf("wallet with id %+v already exists", id))
	}
	newWallet := NewWallet(id)
	v.slot[id] = newWallet
	return nil
}

func (v *vault) Remove(id [5]uint32) error {
	_, exists := v.slot[id]
	if exists {
		return errors.New(fmt.Sprintf("wallet with id %+v already exists", id))
	}
	return nil
}
