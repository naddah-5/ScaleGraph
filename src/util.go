package src

import (
	"errors"
	"main/src/kademlia"
	"math/bits"
	"math/rand"
)

// returns a pseudo-random uint32 in the range (min, max]
func randU32(min uint32, max uint32) (uint32, error) {
	if min >= max {
		return 0, errors.New("invalid range")
	}
	x := rand.Uint32()
	x %= (max - min)
	x += min
	return x, nil
}

// returns the xor distance metric for between the nodes
func RelativeDistance(nodeA [5]uint32, nodeB [5]uint32) [5]uint32 {
	relDist := [5]uint32{0, 0, 0, 0, 0}
	for i := 0; i < 5; i++ {
		relDist[i] = nodeA[i] ^ nodeB[i]
	}
	return relDist
}

// Returns true if node A is closer to the target than node B, returns false if node B is closer to target than node A.
// Returns an error if node A and B are the same distance from the target.
func CloserNode(nodeA [5]uint32, nodeB [5]uint32, target [5]uint32) (bool, error) {
	distA := RelativeDistance(nodeA, target)
	distB := RelativeDistance(nodeB, target)
	for i := 0; i < 5; i++ {
		if distA[i] > distB[i] {
			return true, nil
		} else if distB[i] > distA[i] {
			return false, nil
		}
	}
	return false, errors.New("nodes have same distance to target node")
}

// Returns the shared prefix length between the supplied ID's
func DistPrefixLength(idA [5]uint32, idB [5]uint32) int {
	length := 0
	for i := 0; i < len(idA); i++ {
		segDist := bits.LeadingZeros32(idA[i] ^ idB[i])
		length += segDist
		if segDist != 32 {
			break
		}
	}
	return length
}

// returns true if node A is larger than node node
// returns false if node B is larger than or equal to node A
func LargerNode(nodeA [5]uint32, nodeB [5]uint32) bool {
	for i := 0; i < 5; i++ {
		if nodeA[i] < nodeB[i] {
			return false
		}
	}
	return true
}

// sorts contact slice based on distance to the target
func SortContactsByDistance(input *[]kademlia.Contact, target [5]uint32) {
	for i := 1; i < len(*input); i++ {
		for j := 0; j < len(*input)-1; j++ {
			nodeA := (*input)[j]
			nodeB := (*input)[j+1]
			sortCriterion, err := CloserNode(nodeA.ID(), nodeB.ID(), target)
			if sortCriterion || (err != nil) {
				(*input)[j] = nodeB
				(*input)[j+1] = nodeA
				if LargerNode(nodeA.ID(), nodeB.ID()) {
					(*input)[j] = nodeB
					(*input)[j+1] = nodeA
				}
			}
		}
	}
}
