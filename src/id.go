package main

import (
	"math/rand"
)

// create 5 uint32 (160 bits) which are which represents the Kademlia ID
func GenerateID() ([5]uint32, error) {
	var bitArray [5]uint32
	for i := 0; i < 5; i++ {
		var section uint32 = rand.Uint32()
		bitArray[i] = section
	}
	return bitArray, nil
}

func RelativeDistance(nodeA *[5]uint32, nodeB *[5]uint32) {
	var relDist int
	for i := 0; i < len(*nodeA); i++ {
		
	}
}
