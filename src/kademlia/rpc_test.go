package kademlia

import (
	"log"
	"testing"
)

func TestRPCDisplay(t *testing.T) {
	verbose := true
	if verbose {
		sender := NewRandomContact()
		rpc := GenerateRPC(sender)

		log.Printf("rpc is currently:\n%s", rpc.Display())

		receiver := NewRandomContact()
		rpc.Ping(receiver.IP())

		log.Print("updating RPC to a ping")
		log.Printf("rpc is currently:\n%s", rpc.Display())
	}
}
