package scalegraph

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type block struct {
	*transaction
	*consensus
	hash []byte
}

// Build the base block of a chain, which only contains the height and hash.
// The base hash is a sha256 hash of the walletID
func BaseBlock(walletID [5]uint32) *block {
	hasher := sha256.New()
	for i := range walletID {
		tmp := strconv.FormatUint(uint64(walletID[i]), 10)
		hasher.Write([]byte(tmp))
	}
	hash := hasher.Sum(nil)
	block := block{
		hash: hash,
	}
	return &block
}

func BuildBlock(trx *transaction) *block {
	hash := trx.Hash()
	block := block{
		transaction: trx,
		hash:        hash,
	}
	return &block
}

func (block *block) display() string {
	disp := "block:\n"
	if block.transaction != nil {
		disp += block.transaction.display()
	} else {
		disp += "transaction:\n"
	}
	if block.consensus != nil {
		disp += block.consensus.display()
	} else {
		disp += "consensus:\n"
	}
	disp += "hash: " + fmt.Sprint(block.hash)

	return disp
}
