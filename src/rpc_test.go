package scalegraph

import (
	"log"
	"testing"
)

func TestAssertion(t *testing.T) {
	verbose := false
	var kNodes []contact
	for i := 0; i < 3; i++ {
		newContact := NewRandomContact()
		kNodes = append(kNodes, newContact)
	}
	testRPC := RPC{
		CMD:    FIND_NODE_RESPONSE,
		Sender: NewRandomContact(),
		KNodes: kNodes,
	}
	if verbose {
		log.Printf("%+v", testRPC)
		log.Printf("%+v", testRPC.Sender)
		log.Printf("%+v", testRPC.KNodes)
	}
}
