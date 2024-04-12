package scalegraph

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	verbose := true
	senderCons := NewConsensus()
	senderCons.fillSender(0, make([]byte, 100))
	if verbose {
		log.Printf("consensus from sender: %+v", senderCons)
	}
	receiverCons := NewConsensus()
	receiverCons.fillReceiver(0, make([]byte, 100))
	if verbose {
		log.Printf("consensus from receiver: %+v", receiverCons)
	}

}
