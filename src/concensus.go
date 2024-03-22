package scalegraph

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
)

type signature struct {
	id   [5]uint32
	hash []byte
}

func NewSign(id [5]uint32) signature {
	pubSign := hashID(id)
	sign := signature{
		id:   id,
		hash: pubSign,
	}
	return sign
}

func hashID(id [5]uint32) []byte {
	hasher := sha256.New()
	for i := range id {
		tmp := strconv.FormatUint(uint64(id[i]), 10)
		hasher.Write([]byte(tmp))
	}
	return hasher.Sum(nil)

}

func (sign *signature) Validate() error {
	validationSign := hashID(sign.id)
	if CompareHash(sign.hash, validationSign) {
		return nil
	}

	return errors.New(fmt.Sprintf("error: invalid signature, %+v", sign.hash))
}

type consensus struct {
	senderBlockHeight     int
	senderHashLastBlock   []byte
	receiverBlockHeight   int
	receiverLastBlockHash []byte
	validation            []signature
}

// Might want to flip consensus order

func NewConsensus(height int, hash []byte) consensus {
	return consensus{
		senderBlockHeight:   height,
		senderHashLastBlock: hash,
		validation:          make([]signature, 0, 2*REPLICATION),
	}
}

func (cons *consensus) Concur(height int, hash []byte) {
	cons.receiverBlockHeight = height
	cons.receiverLastBlockHash = hash
}

func (cons *consensus) Approved(sign signature) {
	cons.validation = append(cons.validation, sign)
}
