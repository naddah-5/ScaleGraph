package scalegraph

import (
	"log"
	"testing"
)

func TestMergeValidationSender(t *testing.T) {
	testName := "TestMergeValidationSender"
	verbose := false
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
	senderCons.MergeValdation(receiverCons)
	if verbose {
		log.Printf("%+v", senderCons)
	}
	if senderCons.receiverValidation == nil {
		t.Fail()
		log.Printf("[%s] - failed to merge receiver validation into sender consensus", testName)
	}
}

func TestMergeValidationReceiver(t *testing.T) {
	testName := "TestMergeValidationReceiver"
	verbose := true
	receiverCons := NewConsensus()
	receiverCons.fillReceiver(0, make([]byte, 100))
	if verbose {
		log.Printf("consensus from receiver: %+v", receiverCons)
	}
	senderCons := NewConsensus()
	senderCons.fillSender(0, make([]byte, 100))
	if verbose {
		log.Printf("consensus from sender: %+v", senderCons)
	}
	receiverCons.MergeValdation(senderCons)
	if verbose {
		log.Printf("%+v", receiverCons)
	}
	if receiverCons.receiverValidation == nil {
		t.Fail()
		log.Printf("[%s] - failed to merge sender validation into sender consensus", testName)
	}
}
func TestMergeSignatures(t *testing.T) {

}
