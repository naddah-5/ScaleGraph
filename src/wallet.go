package scalegraph

import "sync"

type wallet struct {
	lock     sync.RWMutex
	walletID [5]uint32
	pubKey   []byte
	*blockchain
}

func NewWallet(id [5]uint32) *wallet {
	chain := NewBlockchain(id)
	newWallet := wallet{
		lock:       sync.RWMutex{},
		walletID:   id,
		blockchain: chain,
	}
	return &newWallet
}
