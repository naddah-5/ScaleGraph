package scalegraph

import (
	"log"
	"testing"
)

func TestHash(t *testing.T) {
	verbose := false
	trx := transaction{
		Sender:   GenerateID(),
		Receiver: GenerateID(),
		Amount:   10,
	}
	if verbose {
		log.Printf("transaction is: %+v\n", trx)
	}
	hash := trx.Hash()
	if verbose {
		log.Printf("hash number one is: %+v\n", hash)
	}
	hash1 := trx.Hash()
	if verbose {
		log.Printf("hash number two is: %+v\n", hash1)
	}
	for i := range hash {
		if hash[i] != hash1[i] {
			t.Fail()
			log.Printf("%+v at index %d in hash does not match %+v at index %d in hash1", hash[i], i, hash1[i], i)
		}
	}
}