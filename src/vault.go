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

func NewVault() *vault {
	return &vault{
		slot: make(map[[5]uint32]*wallet),
	}
}

func (vault *vault) Add(wallet *wallet) error {
	vault.lock.Lock()
	defer vault.lock.Unlock()
	_, exists := vault.slot[wallet.walletID]
	if exists {
		return errors.New(fmt.Sprintf("wallet with id %+v already exists", wallet.walletID))
	}
	vault.slot[wallet.walletID] = wallet
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

func (vault *vault) FindWallet(id [5]uint32) (*wallet, error) {
	vault.lock.RLock()
	defer vault.lock.RUnlock()
	
	wallet, ok := vault.slot[id]
	if !ok {
		return nil, errors.New("wallet not found")
	}
	return wallet, nil
}
