package scalegraph

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	verbose := true
	senderCons := NewConsensus()
	randomSenderBytes := make([]byte, 100)
	randomReceiverBytes := make([]byte, 100)
	for i := 0; i < 100; i++ {
		bSender, _ := randU32(0, 256)
		bReceiver, _ := randU32(0, 256)
		randomSenderBytes[i] = byte(bSender)
		randomReceiverBytes[i] = byte(bReceiver)
	}
	senderCons.fillSender(0, randomSenderBytes)
	if verbose {
		log.Print("Senders concensus")
		log.Print(senderCons.senderValidation.display())
	}
	receiverCons := NewConsensus()
	receiverCons.fillReceiver(0, randomReceiverBytes)
	if verbose {
		log.Print("Receivers concensus")
		log.Print(receiverCons.receiverValidation.display())
	}

}
