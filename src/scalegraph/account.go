package scalegraph

import "sync"

type Account struct {
	sync.RWMutex
	id [5]uint32
	BlockChain
}

func NewAccount(id [5]uint32) *Account {
	acc := Account{
		id: id,
		BlockChain: *NewBlockChain(),
	}
	return &acc
}
