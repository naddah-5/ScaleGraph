package main

import (
	"fmt"
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

func prefixLength(idA [5]uint32, idB [5]uint32) int {
	var length int = 0
	var mask uint32 = 1 << 31
	fmt.Println(mask)
	prefixBranch:
	for i := 0; i < len(idA); i++ {
		for j := 0; j < 32; j++ {
			if true {
				length++
			} else {
				break prefixBranch
			}
		}
	}
	return length
}

func sigMatch(a uint32, b uint32) bool {
	var mask uint32 = 1 << 31
	var sigA uint32 = a & mask
	var sigB uint32 = b & mask
	if sigA == sigB {
		return true
	}
	return false
}
