package main

import (
	"container/list"
	"errors"
	"fmt"
	"math/bits"
)

func RelativeDistance(nodeA [5]uint32, nodeB [5]uint32) int {
	var relDist int = 0
	for i := 0; i < len(nodeA); i++ {
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

func DistPrefixLength(idA [5]uint32, idB [5]uint32) int {
	var length int = 0
	for i := 0; i < len(idA); i++ {
		var segDist int = bits.LeadingZeros32(idA[i]^idB[i])
		length += segDist
		if segDist != 32 {
			break
		}
	}
	return length
}

// Returns true if node A is closer to or the same distance to target node as node B.
func CloserNode(nodeA [5]uint32, nodeB [5]uint32, target [5]uint32) bool {
	var relDistA int = RelativeDistance(nodeA, target)
	var relDistB int = RelativeDistance(nodeB, target)
	if relDistA <= relDistB {
		return true
	}
	return false
}

func SortByDistance(contactList *list.List, target [5]uint32) error {
	var relDist int
	var nextRelDist int
	for i := 0; i < contactList.Len(); i++ {
		for e := contactList.Front(); e != nil; e = e.Next() {
			if e.Next() == nil {
				// DO NOT REMOVE: The e != nil in the loop header is purely decorative
				// it does not share scope with the loop body
				break
			}
			elem, ok := e.Value.(contact)
			if !ok {
				return errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact found %+v\n", e.Value))
			}
			nextElem, ok := e.Next().Value.(contact)
			if !ok {
				return errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact found %+v\n", e.Value))
			}

			relDist = RelativeDistance(elem.ID(), target)
			nextRelDist = RelativeDistance(nextElem.ID(), target)
			if relDist > nextRelDist  {
				contactList.MoveAfter(e, e.Next())
			}
		}
	}

	return nil
}
