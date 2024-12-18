package scalegraph

import "sync"

type wallet struct {
	walletLock sync.RWMutex
	walletID   [5]uint32
	pubKey     []byte
	balance    int
	*blockchain
}

func NewWallet(id [5]uint32, balance int) *wallet {
	chain := NewBlockchain(id)
	newWallet := wallet{
		walletLock: sync.RWMutex{},
		walletID:   id,
		blockchain: chain,
	}
	return &newWallet
}

func (wallet *wallet) Balance() int {
	wallet.walletLock.RLock()
	defer wallet.walletLock.RUnlock()
	return wallet.balance
}

func (wallet *wallet) BuildBlock(trx *transaction) *block {
	newBlock := wallet.NewBlock(trx)
	cons := NewConsensus()
	switch wallet.walletID {
	case trx.sender:
		cons.fillSender(wallet.lastHeight(), wallet.lastHash())
	case trx.receiver:
		cons.fillReceiver(wallet.lastHeight(), wallet.lastHash())
	}
	newBlock.consensus = cons

	return newBlock
}
