package scalegraph

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
)

type transaction struct {
	// note that the id is not unique, but has low collision risk
	// espescially when combined with the hash
	trxID  [5]uint32
	sender [5]uint32
	signature
	receiver [5]uint32
	amount   int
}

func NewTransaction(sender [5]uint32, receiver [5]uint32, amount int) transaction {
	sign := *NewSign(sender)
	trx := transaction{
		trxID:     GenerateID(),
		sender:    sender,
		signature: sign,
		receiver:  receiver,
		amount:    amount,
	}
	return trx
}

// returns a hash of the transaction struct in order
func (trx *transaction) Hash() []byte {
	var hash []byte
	hasher := sha256.New()
	for i := range trx.trxID {
		tmp := strconv.FormatUint(uint64(trx.trxID[i]), 10)
		hasher.Write([]byte(tmp))
	}
	for i := range trx.sender {
		tmp := strconv.FormatUint(uint64(trx.sender[i]), 10)
		hasher.Write([]byte(tmp))
	}
	hasher.Write(trx.signature.hash)
	for i := range trx.receiver {
		tmp := strconv.FormatUint(uint64(trx.receiver[i]), 10)
		hasher.Write([]byte(tmp))
	}
	strAmount := strconv.FormatInt(int64(trx.amount), 10)
	hasher.Write([]byte(strAmount))

	hash = hasher.Sum(nil)
	return hash
}

// Returns the balance change for the specified wallet id.
// Returns an error if the given wallet id is not involved in the transaction.
func (trx *transaction) delta(walletID [5]uint32) (int, error) {
	if trx.sender == walletID {
		return -trx.amount, nil
	} else if trx.receiver == walletID {
		return trx.amount, nil
	}
	return 0, errors.New(fmt.Sprintf("error: %+v is not involved in transaction", walletID))
}
