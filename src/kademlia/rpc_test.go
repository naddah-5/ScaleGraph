package kademlia

import (
	"log"
	"testing"
)

func TestRPCDisplay(t *testing.T) {
	verbose := false
	if verbose {
		sender := NewRandomContact()
		receiver := NewRandomContact()
		rpc := GenerateRPC(receiver.IP(), sender)

		log.Printf("rpc is currently:\n%s", rpc.Display())

		rpc.Ping()

		log.Print("updating RPC to a ping")
		log.Printf("rpc is currently:\n%s", rpc.Display())
	}
}
