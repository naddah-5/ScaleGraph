package scalegraph

import (
	"errors"
	"fmt"
	"sync"
)

type blockchain struct {
	lock sync.RWMutex
	chain []block
}

func NewBlockchain(walletID [5]uint32) *blockchain {
	blockchain := blockchain{
		lock: sync.RWMutex{},
		chain: make([]block, 0, 100),
	}
	blockchain.chain = append(blockchain.chain, BaseBlock(walletID))
	return &blockchain
}

func (b *blockchain) LastHash() []byte {
	return b.chain[len(b.chain)-1].hash
}

func (b *blockchain) LastHeight() int {
	return len(b.chain)
}

// Currently uses a bad checking method, but it is simple.
func (b *blockchain) Grow(block block) error {
	if block.senderBlockHeight != b.LastHeight()+1 {
		if block.receiverBlockHeight != b.LastHeight()+1 {
			return errors.New(fmt.Sprintf("invalid block: current block height %d, sender height %d, receiver height %d", b.LastHeight(), block.senderBlockHeight, block.receiverBlockHeight))
		}
	}
	ok := CompareHash(b.LastHash(), block.senderHashLastBlock)
	if !ok {
		ok = CompareHash(b.LastHash(), block.receiverHashLastBlock)
		if !ok {
			return errors.New("invalid block: does not match last block hash")
		}
	}

	// add check for consensus here

	b.chain = append(b.chain, block)
	return nil
}

func (b *blockchain) NewBlock(trx transaction, cons consensus) block {
	height := len(b.chain)
	return BuildBlock(height, trx, cons)
}
