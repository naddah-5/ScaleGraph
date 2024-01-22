package main

import (
	"math/bits"
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

func RelativeDistance(nodeA *[5]uint32, nodeB *[5]uint32) int {
	var relDist int = 0
	for i := 0; i < len(*nodeA); i++ {
		relDist += hammingDistance(nodeA[i], nodeB[i])
	}
	return relDist
}

func hammingDistance(a uint32, b uint32) int {
	var hamDist int = 0
	var diffID uint32 = a ^ b
	for diffID > 0 {
		hamDist++
		rshift := bits.TrailingZeros32(diffID) + 1
		diffID = diffID >> uint32(rshift)
	}
	return hamDist
}
