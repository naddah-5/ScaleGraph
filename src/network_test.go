package scalegraph

import (
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	VERBOSE := false

	in := make(chan RPC)
	out := make(chan RPC)
	net := NewNetwork(in, out, [4]byte{0, 0, 0, 0}, [4]byte{1, 1, 1, 1})
	input := make([][5]uint32, 10)
	for i := range input {
		id := GenerateID()
		net.Add(id)
		input[i] = id
	}
	if VERBOSE {
		log.Printf("handler is:")
		for key, value := range net.set {
			log.Printf("key: %+v, value: %+v", key, value)
		}
	}
	
	_, err := net.Retrieve(input[0])
	if err != nil {
		t.Fail()
	}

	for i := range input {
		if i % 2 == 0 {
			net.Drop(input[i])
		}
	}

	if VERBOSE {
		log.Printf("handler is:")
		for key, value := range net.set {
			log.Printf("key: %+v, value: %+v", key, value)
		}
	}

	unexpected, err := net.Retrieve(input[0])
	if err == nil {
		t.Fail()
		log.Printf("expected handle to be removed, found:%+v", unexpected)
	}
}
