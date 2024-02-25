package scalegraph

type wallet struct {
	ID    [5]uint32
	PK    []byte
	Chain chain
}

func NewWallet(id [5]uint32) wallet {
	newWallet := wallet{
		ID: id,
	}
	return newWallet
}
