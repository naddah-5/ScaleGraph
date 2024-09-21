package scalegraph

import (
	"log"
	"testing"
)

func TestNewChain(t *testing.T) {
	bc := NewBlockchain(NewRandomContact().id, 0)
	log.Println(bc.display())
}
