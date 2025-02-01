package scalegraph

import (
	"fmt"
	"sync"
)

type BlockChain struct {
	sync.RWMutex
	chain []Block
}

func NewBlockChain() *BlockChain {
	bc := BlockChain{
		chain: make([]Block, 0, 32),
	}
	return &bc
}

// Appends a new block containing the transaction to the end of the block chain.
func (bc *BlockChain) AddBlock(trx *Transaction) {
	bc.Lock()
	defer bc.Unlock()
	var newBlock Block
	if len(bc.chain) == 0 {
		// if there are no previous blocks the chain must be started
		newBlock = *(FirstBlock(RandomID(), trx))
	} else {
		newBlock = *(bc.chain[len(bc.chain)].NewBlock(RandomID(), trx))
	}
	bc.chain = append(bc.chain, newBlock)
}


func (bc *BlockChain) Display() string {
	bc.Lock()
	defer bc.Unlock()
	disp := ""
	for i, b := range bc.chain {
		disp += fmt.Sprintf("block %3d:\n", i)
		disp += fmt.Sprintf(b.Display())
	}
	return disp
}
