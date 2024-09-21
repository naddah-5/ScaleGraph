package scalegraph

import (
	"errors"
	"fmt"
)

type blockchain struct {
	chain []*block
}

func NewBlockchain(walletID [5]uint32, balance int) *blockchain {
	blockchain := blockchain{
		chain: make([]*block, 0, 100),
	}
	blockchain.chain = append(blockchain.chain, BaseBlock(walletID))
	return &blockchain
}

func (blockchain *blockchain) display() string {
	disp := "chain:\n"
	for i, v := range blockchain.chain {
		disp += fmt.Sprintf("block %d\n", i)
		disp += v.display()
	}

	return disp
}

func (blockchain *blockchain) lastHash() []byte {
	return blockchain.chain[len(blockchain.chain)-1].hash
}

func (blockchain *blockchain) lastHeight() int {
	return len(blockchain.chain)
}

func (blockchain *blockchain) lastBlock() *block {
	return blockchain.chain[len(blockchain.chain)-1]
}

func (blockchain *blockchain) Grow(newBlock *block, walletID [5]uint32) error {
	blockErr := blockchain.validateBlockData(newBlock, walletID)
	if blockErr != nil {
		return blockErr
	}
	consErr := blockchain.ValidateConsensus(newBlock)
	if consErr != nil {
		return consErr
	}

	blockchain.chain = append(blockchain.chain, newBlock)
	return nil
}

// Check the new block to verify if the concensus details regarding this chain is valid.
// Returns an error if the last block on the chain does not match either the sender or receivers last block.
// Does not lock.
func (blockchain *blockchain) validateBlockData(newBlock *block, walletID [5]uint32) error {
	invalidBlock := errors.New(fmt.Sprintf("error: %+v is not a valid block for blockchain, last block %+v\n", newBlock, blockchain.lastBlock()))
	if walletID == newBlock.receiver {
		if *newBlock.consensus.receiverValidation.blockHeight != blockchain.lastHeight()+1 {
			return invalidBlock
		} else if CompareHash(newBlock.consensus.receiverValidation.hashLastBlock, blockchain.lastBlock().hash) {
			return invalidBlock
		}
	} else if walletID == newBlock.sender {
		if *newBlock.consensus.senderValidation.blockHeight != blockchain.lastHeight()+1 {
			return invalidBlock
		} else if CompareHash(newBlock.consensus.senderValidation.hashLastBlock, blockchain.lastBlock().hash) {
			return invalidBlock
		}
	} else {
		return invalidBlock
	}
	return nil
}

// Go through all signatures in the consensus and validate them by hashing.
// Returns an errror at the first invalid signature.
func (blockchain *blockchain) ValidateConsensus(newBlock *block) error {
	for _, sign := range newBlock.signatureList {
		err := sign.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (blockchain *blockchain) NewBlock(trx *transaction) *block {
	return BuildBlock(trx)
}
