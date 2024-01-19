package main

import (
	"fmt"
	"math/rand"
)

// create 20 uint8 (160 bits) which are converted into a string which represents the Kademlia ID
func GenerateID() (string, error) {
	var bitvector [5]uint32
	for i := 0; i < 5; i++ {
		var section uint32 = rand.Uint32()
		bitvector[i] = section
	}
	fmt.Println(bitvector)
	return "", nil
}
