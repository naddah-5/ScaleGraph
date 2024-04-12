package scalegraph

import (
	"testing"
	"log"
)

func TestSanity(t *testing.T) {
	testName := "TestSanity"
	verbose := false

	testWallet := NewWallet(GenerateID(), 0)
	if verbose {
		log.Println(testName)
		log.Printf("current block: %+v\n", testWallet.lastBlock())
	}
}
