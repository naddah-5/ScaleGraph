package scalegraph

import (
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	VERBOSE := true

	in := make(chan RPC)
	out := make(chan RPC)
	net := NewNetwork(in, out, [4]byte{0, 0, 0, 0})
	input := make([]RPC, 10)
	for range input {
		id := GenerateID()
		net.handler.Add(id)
	}
	if VERBOSE {
		log.Printf("handler is:")
	}
}
