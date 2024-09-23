package scalegraph

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type signature struct {
	id   [5]uint32
	hash []byte
}

func (sig *signature) display() string {
	disp := fmt.Sprintf("id: %v\n", sig.id)
	disp += fmt.Sprintf("hash: %s\n", fmt.Sprint(sig.hash))
	return disp
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
	senderValidation   *blockValidation
	receiverValidation *blockValidation
	signatureList      []*signature
}

func (cons *consensus) display() string {
	disp := "signatures: \n"
	for i := 0; i < len(cons.signatureList); i++ {
		disp += cons.signatureList[i].display()
	}
	disp += "sender block:\n"
	if cons.senderValidation != nil {
		disp += fmt.Sprint(cons.senderValidation.display() + "\n")
	}
	disp += "receiver block:\n"
	if cons.receiverValidation != nil {
		disp += fmt.Sprint(cons.receiverValidation.display() + "\n")
	}
	return disp
}

type blockValidation struct {
	blockHeight   *int // must be a pointer to not get a default value
	hashLastBlock []byte
}

func (blockVal *blockValidation) display() string {
	disp := "block validation:\n"
	disp += "block height "
	if blockVal.blockHeight != nil {
		disp += fmt.Sprint(blockVal.blockHeight) + "\n"
	}
	disp += "hash "
	if blockVal.hashLastBlock != nil {
		disp += fmt.Sprint(blockVal.hashLastBlock) + "\n"
	}
	return disp
}

func NewConsensus() *consensus {
	return &consensus{
		signatureList: make([]*signature, 0, 2*REPLICATION),
	}
}

func (cons *consensus) fillSender(height int, hash []byte) {
	sender := &blockValidation{
		blockHeight:   &height,
		hashLastBlock: hash,
	}
	cons.senderValidation = sender
}

func (cons *consensus) fillReceiver(height int, hash []byte) {
	sender := &blockValidation{
		blockHeight:   &height,
		hashLastBlock: hash,
	}
	cons.receiverValidation = sender
}

func (cons *consensus) signConsensus(sign *signature) error {
	err := sign.Validate()
	if err != nil {
		log.Printf("received invalid signature for consensus signature:\n%s", sign.display())
		return err
	}
	cons.signatureList = append(cons.signatureList, sign)
	return nil
}

func (cons *consensus) Merge(secondCons *consensus) {
	if cons.receiverValidation == nil {
		cons.receiverValidation = secondCons.receiverValidation
	}
	if cons.senderValidation == nil {
		cons.senderValidation = secondCons.senderValidation
	} 
	cons.signatureList = append(cons.signatureList, secondCons.signatureList...)
}
