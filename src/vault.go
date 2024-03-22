package scalegraph

import (
	"errors"
	"fmt"
	"sync"
)

type vault struct {
	lock sync.RWMutex
	slot map[[5]uint32]*wallet
}

func NewVault(id [5]uint32) vault {
	return vault{
		slot: make(map[[5]uint32]*wallet),
	}
}

func (vault *vault) Add(id [5]uint32) error {
	vault.lock.Lock()
	defer vault.lock.Unlock()
	_, exists := vault.slot[id]
	if exists {
		return errors.New(fmt.Sprintf("wallet with id %+v already exists", id))
	}
	newWallet := NewWallet(id, 0)
	vault.slot[id] = newWallet
	return nil
}

func (vault *vault) Remove(id [5]uint32) error {
	vault.lock.Lock()
	defer vault.lock.Unlock()
	_, exists := vault.slot[id]
	if exists {
		return errors.New(fmt.Sprintf("wallet with id %+v already exists", id))
	}
	return nil
}
