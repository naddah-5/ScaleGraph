package scalegraph

import (
	"log"
	"testing"
)

func TestNewChain(t *testing.T) {
	bc := NewBlockchain(NewRandomContact().id)
	log.Println(bc.display())
}
