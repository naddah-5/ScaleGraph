package scalegraph

type wallet struct {
	walletID   [5]uint32
	pk         []byte
	blockchain blockchain
}

func NewWallet(id [5]uint32) wallet {
	chain := NewBlockchain(id)
	newWallet := wallet{
		walletID:   id,
		blockchain: chain,
	}
	return newWallet
}
