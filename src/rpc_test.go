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
		sender: NewRandomContact(),
		kNodes: kNodes,
	}
	if verbose {
		log.Printf("%+v", testRPC)
		log.Printf("%+v", testRPC.sender)
		log.Printf("%+v", testRPC.kNodes)
	}
}
