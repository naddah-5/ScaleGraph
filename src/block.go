package scalegraph

import (
	"crypto/sha256"
	"strconv"
)

type block struct {
	transaction
	consensus
	hash []byte
}

// Build the base block of a chain, which only contains the height and hash.
// The base hash is a sha256 hash of the walletID
func BaseBlock(walletID [5]uint32) block {
	var hash []byte
	hasher := sha256.New()
	for i := range walletID {
		tmp := strconv.FormatUint(uint64(walletID[i]), 10)
		hasher.Write([]byte(tmp))
	}
	hasher.Write(hash)
	block := block{
		hash:   hash,
	}
	return block
}

// Construct a block from the sender side, i.e. a proposal.
func BuildBlock(height int, trx transaction, cons consensus) block {
	hash := trx.Hash()
	block := block{
		transaction: trx,
		consensus:   cons,
		hash:        hash,
	}
	return block
}

