package scalegraph

import (
	"crypto/sha256"
	"strconv"
)

type transaction struct {
	Sender    [5]uint32
	Receiver  [5]uint32
	Amount    int
	signature
}

func NewTransaction(sender [5]uint32, receiver [5]uint32, amount int, sign signature) transaction {
	trx := transaction{
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		signature: sign, 
		// signature is supposed to be PPK encryption of sender, receiver, and amount
		// for now it can be a hash
	}
	return trx
}

// returns a hash of the transaction struct in order
func (t *transaction) Hash() []byte {
	var hash []byte
	hasher := sha256.New()
	for i := range t.Sender {
		tmp := strconv.FormatUint(uint64(t.Sender[i]), 10)
		hasher.Write([]byte(tmp))
	}
	for i := range t.Receiver {
		tmp := strconv.FormatUint(uint64(t.Receiver[i]), 10)
		hasher.Write([]byte(tmp))
	}
	strAmount := strconv.FormatInt(int64(t.Amount), 10)
	hasher.Write([]byte(strAmount))
	hasher.Write(t.signature.hash)

	hash = hasher.Sum(nil)
	return hash
}
