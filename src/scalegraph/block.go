package scalegraph

type Block struct {
	Transaction
	id     [5]uint32
	prevID [5]uint32
}

func FirstBlock() *Block {
	block := Block{
		id: RandomID(),
	}
	return &block
}

func (block *Block) NewBlock(trx *Transaction) *Block {
	b := Block{
		id:     RandomID(),
		prevID: block.id,
	}
	return &b
}
