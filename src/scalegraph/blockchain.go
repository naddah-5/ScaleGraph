package scalegraph

import "sync"

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
