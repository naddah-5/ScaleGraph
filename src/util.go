package scalegraph


import (
	"container/list"
	"errors"
	"fmt"
	"log"
	"math/bits"
	"math/rand"
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

// Returns the shared prefix length between the supplied ID's
func DistPrefixLength(idA [5]uint32, idB [5]uint32) int {
	var length int = 0
	for i := 0; i < len(idA); i++ {
		var segDist int = bits.LeadingZeros32(idA[i] ^ idB[i])
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
	for i := 0; i <= contactList.Len(); i++ {
		for e := contactList.Front(); e != nil && e.Next() != nil; e = e.Next() {
			elem, ok := e.Value.(contact)
			if !ok {
				return errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact found %+v\n", e.Value))
			}
			nextElem, ok := e.Next().Value.(contact)
			if !ok {
				return errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact found %+v\n", e.Next().Value))
			}

			relDist = RelativeDistance(elem.ID(), target)
			nextRelDist = RelativeDistance(nextElem.ID(), target)
			if relDist > nextRelDist {
				e.Next().Value = elem
				e.Value = nextElem

			}
		}
	}

	return nil
}

func MergeByDistance(contactListA *list.List, contactListB *list.List, target [5]uint32) (*list.List, error) {
	var relDistA int
	var relDistB int

	var res *list.List = list.New()
	var listA *list.List
	var listB *list.List

	if contactListA.Len() <= contactListB.Len() {
		listA = contactListA
		listB = contactListB
	} else {
		listA = contactListB
		listB = contactListA
	}

	for listA.Len() > 0 && listB.Len() > 0 {
		elemA, ok := listA.Front().Value.(contact)
		if !ok {
			log.Printf("bucket has been corrupted: expected a contact found %+v\n", listA.Front())
			return res, errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact found %+v\n", listA.Front()))
		}
		elemB, ok := listB.Front().Value.(contact)
		if !ok {
			log.Printf("bucket has been corrupted: expected a contact found %+v\n", listB.Front())
			return res, errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact found %+v\n", listB.Front()))
		}

		relDistA = RelativeDistance(elemA.ID(), target)
		relDistB = RelativeDistance(elemB.ID(), target)
		if relDistA <= relDistB {
			res.PushBack(elemA)
			listA.Remove(listA.Front())
		} else {
			res.PushBack(elemB)
			listB.Remove(listB.Front())
		}
	}
	if listA.Len() > 0 {
		for e := listA.Front(); e != nil; e = e.Next() {
			elem, ok := e.Value.(contact)
			if !ok {
				return res, errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact found %+v\n", e.Value))
			} else {
				res.PushBack(elem)
			}
		}
	} else if listB.Len() > 0 {
		for e := listB.Front(); e != nil; e = e.Next() {
			elem, ok := e.Value.(contact)
			if !ok {
				return res, errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact found %+v\n", e.Value))
			} else {
				res.PushBack(elem)
			}
		}
	}

	return res, nil
}

// returns a pseudo-random uint32 in the range (min, max]
func randU32(min uint32, max uint32) (uint32, error) {
	if min >= max {
		return 0, errors.New("invalid range")
	}
	var x uint32 = rand.Uint32()
	x %= (max - min)
	x += min
	return x, nil
}

// create 5 uint32 (160 bits) which are which represents the Kademlia ID
func GenerateID() [5]uint32 {
	var bitArray [5]uint32
	for i := 0; i < 5; i++ {
		var section uint32 = rand.Uint32()
		bitArray[i] = section
	}
	return bitArray
}

func GenerateIP() [4]byte {
	var ip [4]byte
	for i := 0; i < 4; i++ {
		tmpSeg, _ := randU32(0, 256)
		var seg byte = byte(tmpSeg)
		ip[i] = seg
	}
	return ip
}
