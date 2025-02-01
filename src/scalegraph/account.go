package scalegraph

import (
	"fmt"
	"sync"
)

type Account struct {
	sync.RWMutex
	id [5]uint32
	BlockChain
}

func NewAccount(id [5]uint32) *Account {
	acc := Account{
		id:         id,
		BlockChain: *NewBlockChain(),
	}
	return &acc
}

func (acc *Account) VerifyTransaction(trx *Transaction, consensusLimit int) bool {
	if !(trx.sendingAccount != acc.id || trx.receivingAccount != acc.id) {
		return false
	}
	if len(trx.validators) < consensusLimit || len(trx.confirmers) < consensusLimit {
		return false
	}
	return true
}

func (acc *Account) Display() string {
	acc.Lock()
	defer acc.Unlock()
	disp := ""
	disp += fmt.Sprintf("account id: %10v\n", acc.id)
	disp += fmt.Sprintf(acc.BlockChain.Display())
	return disp
}

