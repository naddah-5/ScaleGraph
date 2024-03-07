package scalegraph

import (
	"testing"
)

func TestHandler(t *testing.T) {
	in := make(chan RPC)
	out := make(chan RPC)
	net := NewNetwork(in, out, [4]byte{0, 0, 0, 0})
	input := make([]RPC, 100)
	for range input {
		id := GenerateID()
		net.handler.Add(id)
	}
}
