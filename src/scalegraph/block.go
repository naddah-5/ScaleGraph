package scalegraph

type Block struct {
	Transaction
	id   [5]uint32
	hash []byte
}

func FirstBlock() *Block {
	block := Block{
		id: RandomID(),
	}

	return &block
}

func (block *Block) NewBlock() {

}
