package scalegraph

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	cons := NewConsensus()
	cons.fillSender(0, make([]byte, 100))
	log.Printf("consensus: %+v", cons)
}
