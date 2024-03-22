package scalegraph

import (
	"crypto/sha256"
	"strconv"
)

type transaction struct {
	sender   [5]uint32
	receiver [5]uint32
	amount   int
	signature
}

func NewTransaction(sender [5]uint32, receiver [5]uint32, amount int, sign signature) transaction {
	trx := transaction{
		sender:    sender,
		receiver:  receiver,
		amount:    amount,
		signature: sign,
		// signature is supposed to be PPK encryption of sender, receiver, and amount
		// for now it can be a hash
	}
	return trx
}

// returns a hash of the transaction struct in order
func (trx *transaction) Hash() []byte {
	var hash []byte
	hasher := sha256.New()
	for i := range trx.sender {
		tmp := strconv.FormatUint(uint64(trx.sender[i]), 10)
		hasher.Write([]byte(tmp))
	}
	for i := range trx.receiver {
		tmp := strconv.FormatUint(uint64(trx.receiver[i]), 10)
		hasher.Write([]byte(tmp))
	}
	strAmount := strconv.FormatInt(int64(trx.amount), 10)
	hasher.Write([]byte(strAmount))
	hasher.Write(trx.signature.hash)

	hash = hasher.Sum(nil)
	return hash
}

func (trx *transaction) delta(walletID [5]uint32) int {
	if trx.sender == walletID {
		return -trx.amount
	}
	return trx.amount
}
