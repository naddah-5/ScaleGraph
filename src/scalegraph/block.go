package scalegraph

import "fmt"

type Block struct {
	id     [5]uint32
	prevID [5]uint32
	*Transaction
}

func FirstBlock(id [5]uint32, trx *Transaction) *Block {
	block := Block{
		id:          id,
		prevID:      [5]uint32{0, 0, 0, 0, 0},
		Transaction: trx,
	}
	return &block
}

func (block *Block) NewBlock(id [5]uint32, trx *Transaction) *Block {
	b := Block{
		id:          id,
		prevID:      block.id,
		Transaction: trx,
	}
	return &b
}

func (block *Block) Display() string {
	disp := ""
	disp += fmt.Sprintf("block id: %10v\nprevious block id: %10v\n", block.id, block.prevID)
	disp += fmt.Sprintf(block.Transaction.Display())

	return disp
}
