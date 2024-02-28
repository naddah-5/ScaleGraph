package scalegraph

import (
	"crypto/sha256"
	"strconv"
)

type transaction struct {
	Sender    [5]uint32
	Recipient [5]uint32
	Ammount   int
	Signature []byte
}

// returns a hash of the transaction struct in order
func (t *transaction) Hash() []byte {
	var hash []byte
	hasher := sha256.New()
	for i := range t.Sender {
		tmp := strconv.FormatUint(uint64(t.Sender[i]), 10)
		hasher.Write([]byte(tmp))
	}
	for i := range t.Recipient {
		tmp := strconv.FormatUint(uint64(t.Recipient[i]), 10)
		hasher.Write([]byte(tmp))
	}
	strAmount := strconv.FormatInt(int64(t.Ammount), 10)
	hasher.Write([]byte(strAmount))
	hasher.Write(t.Signature)

	hash = hasher.Sum(nil)
	return hash
}
