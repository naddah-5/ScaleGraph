package scalegraph

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	testName := "TestMerge"
	verbose := false
	if verbose {
		log.Printf("running test: %s", testName)
	}
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

	err := senderCons.Merge(receiverCons)
	if err != nil {
		log.Printf("%s failed: %s", testName, err.Error())
	}
	if verbose {
		log.Printf("merged consensus: \n%v", senderCons.display())
	}
	for i := 0; i < len(senderCons.receiverValidation.hashLastBlock); i++ {
		if senderCons.receiverValidation.hashLastBlock[i] != receiverCons.receiverValidation.hashLastBlock[i] {
			t.Fail()
			log.Printf("%s failed, merged hash does not match original hash", testName)
		}
	}
	if senderCons.receiverValidation.blockHeight != receiverCons.receiverValidation.blockHeight {
		t.Fail()
		log.Printf("%s failed, merged block height does not match original block height", testName)
	}
}

func TestMergeSignatures(t *testing.T) {
	testName := "TestMergeSignatures"
	verbose := false
	cons := NewConsensus()
	for i := 0; i < 10; i++ {
		sign := NewSign(NewRandomContact().id)
		cons.signConsensus(sign)
	}

	secondCons := NewConsensus()
	for i := 0; i < 5; i++ {
		sign := NewSign(NewRandomContact().id)
		secondCons.signConsensus(sign)
	}

	if verbose {
		log.Printf("[%s]:\n", testName)
		log.Print("first consensus")
		log.Print(cons.display())
		log.Print("second consensus")
		log.Print(secondCons.display())
	}

	cons.Merge(secondCons)
	if verbose {
		log.Print("merged consensus:")
		log.Print(cons.display())
	}
	for i := 10; i < 15; i++ {
		if cons.signatureList[i] != secondCons.signatureList[i-10] {
			t.Fail()
			log.Printf("[%s]: consensus merge failed", testName)
		}
	}
}
