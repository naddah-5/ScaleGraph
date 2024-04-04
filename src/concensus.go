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

func NewSign(id [5]uint32) *signature {
	pubSign := hashID(id)
	sign := signature{
		id:   id,
		hash: pubSign,
	}
	return &sign
}

func hashID(id [5]uint32) []byte {
	hasher := sha256.New()
	for i := range id {
		tmp := strconv.FormatUint(uint64(id[i]), 10)
		hasher.Write([]byte(tmp))
	}
	return hasher.Sum(nil)

}

// Validates the concensus signature and returns and error if it is invalid.
func (sign *signature) Validate() error {
	validationSign := hashID(sign.id)
	if CompareHash(sign.hash, validationSign) {
		return nil
	}

	return errors.New(fmt.Sprintf("error: invalid signature, %+v", sign.hash))
}

type consensus struct {
	*senderValidation
	*receiverValidation
	signatureList []signature
}

type senderValidation struct {
	senderBlockHeight   int
	senderHashLastBlock []byte
}

type receiverValidation struct {
	receiverBlockHeight   int
	receiverLastBlockHash []byte
}

func NewConsensus() *consensus {
	return &consensus{
		signatureList: make([]signature, 0, 2*REPLICATION),
	}
}

func (cons *consensus) fillSender(height int, hash []byte) {
	sender := &senderValidation{
		senderBlockHeight: height,
		senderHashLastBlock: hash,
	}
	cons.senderValidation = sender
}
func (cons *consensus) fillReceiver(height int, hash []byte) {
	receiver := &receiverValidation{
		receiverBlockHeight: height,
		receiverLastBlockHash: hash,
	}
	cons.receiverValidation = receiver
}

func (cons *consensus) Merge(secondCons *consensus) {
	if cons.senderValidation != nil {
		cons.receiverValidation = secondCons.receiverValidation
	} else {
		cons.senderValidation = secondCons.senderValidation
	}
	cons.signatureList = append(cons.signatureList, secondCons.signatureList...)
}
