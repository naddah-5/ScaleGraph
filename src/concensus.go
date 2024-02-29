package scalegraph

import (
	"crypto/sha256"
	"strconv"
)

type signature struct {
	ID   [5]uint32
	hash []byte
}

func NewSign(id [5]uint32) signature {
	//TODO
	hasher := sha256.New()
	for i := range(id) {
		tmp := strconv.FormatUint(uint64(id[i]), 10)
		hasher.Write([]byte(tmp))
	}
	pubSign := hasher.Sum(nil)
	sign := signature{
		ID:   id,
		hash: pubSign,
	}
	return sign
}

type consensus struct {
	senderBlockHeight     int
	senderHashLastBlock   []byte
	receiverBlockHeight   int
	receiverHashLastBlock []byte
	validation            []signature
}

func NewConsensus(height int, hash []byte) consensus {
	return consensus{
		senderBlockHeight:   height,
		senderHashLastBlock: hash,
		validation:          make([]signature, 0),
	}
}

func (c *consensus) Concur(height int, hash []byte) {
	c.receiverBlockHeight = height
	c.receiverHashLastBlock = hash
}

func (c *consensus) Approve(sign signature) {
	c.validation = append(c.validation, sign)
}
