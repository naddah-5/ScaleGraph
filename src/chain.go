package scalegraph

import (
	"errors"
	"fmt"
	"sync"
)

type blockchain struct {
	lock  sync.RWMutex
	chain []block
}

func NewBlockchain(walletID [5]uint32, balance int) *blockchain {
	blockchain := blockchain{
		lock:  sync.RWMutex{},
		chain: make([]block, 0, 100),
	}
	blockchain.chain = append(blockchain.chain, BaseBlock(walletID, balance))
	return &blockchain
}

func (blockchain *blockchain) LastHash() []byte {
	return blockchain.chain[len(blockchain.chain)-1].hash
}

func (blockchain *blockchain) LastHeight() int {
	return len(blockchain.chain)
}

func (blockchain *blockchain) LastBlock() block {
	return blockchain.chain[len(blockchain.chain)-1]
}

// Currently uses a bad checking method, but it is simple.
func (blockchain *blockchain) Grow(newBlock block, walletID [5]uint32) error {
	blockErr := blockchain.validateBlockData(newBlock, walletID)
	if blockErr != nil {
		return blockErr
	}
	// add check for consensus here
	consErr := blockchain.validateConsensus(newBlock, walletID)
	if consErr != nil {
		return consErr
	}


	blockchain.chain = append(blockchain.chain, newBlock)
	return nil
}

// Check the new block to verify if the concensus details regarding this chain is valid.
// Returns an error if the last block on the chain does not match either the sender or receivers last block.
func (blockchain *blockchain) validateBlockData(newBlock block, walletID [5]uint32) error {
	invalidBlock := errors.New(fmt.Sprintf("error: %+v is not a valid block for blockchain, last block %+v\n", newBlock, blockchain.LastBlock()))
	if walletID == newBlock.receiver {
		// check valid for chain
		if newBlock.receiverBlockHeight != blockchain.LastHeight()+1 {
			return invalidBlock
		} else if CompareHash(newBlock.receiverLastBlockHash, blockchain.LastBlock().hash) {
			return invalidBlock
		}
	} else if walletID == newBlock.sender {
		// check valid for chain
		if newBlock.senderBlockHeight != blockchain.LastHeight()+1 {
			return invalidBlock
		} else if CompareHash(newBlock.senderHashLastBlock, blockchain.LastBlock().hash) {
			return invalidBlock
		}
	} else {
		return invalidBlock
	}
	return nil
}

// Go through all signatures in the consensus and validate them by hashing.
// Returns an errror at the first invalid signature.
func (blockchain *blockchain) validateConsensus(newBlock block, walletID [5]uint32) error {
	for _, sign := range newBlock.validation {
		err := sign.Validate()
		if err != nil{
			return err
		}
	}
	return nil
}

func (blockchain *blockchain) NewBlock(trx transaction, cons consensus) block {
	height := len(blockchain.chain)
	return BuildBlock(height, trx, cons)
}
