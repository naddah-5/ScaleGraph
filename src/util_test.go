package src

import (
	"fmt"
	"log"
	"main/src/kademlia"
	"testing"
)

func TestRelativeDistance(t *testing.T) {
	testName := "TestRelativeDistance"

	pA := [5]uint32{0, 0, 0, 0, 0}
	pB := [5]uint32{0, 0, 0, 0, 1}
	pC := [5]uint32{0, 0, 1, 0, 0}
	pD := [5]uint32{0, 0, 0, 1, 0}
	pE := [5]uint32{0, 12, 0, 0, 0}
	pF := [5]uint32{0, 0, 0, 0, 10}
	pG := [5]uint32{1, 0, 0, 0, 0}

	pAB := RelativeDistance(pA, pB)
	pAC := RelativeDistance(pA, pC)
	pAD := RelativeDistance(pA, pD)
	pAE := RelativeDistance(pA, pE)
	pAF := RelativeDistance(pA, pF)
	pAG := RelativeDistance(pA, pG)

	pH := [5]uint32{1, 1, 1, 1, 1}
	pI := [5]uint32{1, 1, uint32(10), 1, 1}
	pHI := RelativeDistance(pH, pI)

	if pAB != [5]uint32{0, 0, 0, 0, 1} {
		log.Printf("[%s] - relative distance pA -> pB incorrect, received %v", testName, pAB)
		t.Fail()
	}
	if pAC != [5]uint32{0, 0, 1, 0, 0} {
		log.Printf("[%s] - relative distance pA -> pC incorrect, received %v", testName, pAC)
		t.Fail()
	}
	if pAD != [5]uint32{0, 0, 0, 1, 0} {
		log.Printf("[%s] - relative distance pA -> pD incorrect, received %v", testName, pAD)
		t.Fail()
	}
	if pAE != [5]uint32{0, 12, 0, 0, 0} {
		log.Printf("[%s] - relative distance pA -> pE incorrect, received %v", testName, pAE)
		t.Fail()
	}
	if pAF != [5]uint32{0, 0, 0, 0, 10} {
		log.Printf("[%s] - relative distance pA -> pF incorrect, received %v", testName, pAF)
		t.Fail()
	}
	if pAG != [5]uint32{1, 0, 0, 0, 0} {
		log.Printf("[%s] - relative distance pA -> pG incorrect, received %v", testName, pAG)
		t.Fail()
	}
	if pHI != [5]uint32{0, 0, 11, 0, 0} {
		log.Printf("[%s] - relative distance pH -> pI incorrect, received %v", testName, pHI)
		t.Fail()
	}
}

func TestDistPrefixLength(t *testing.T) {
	testName := "TestDistPrefixLength"
	idA := [5]uint32{0, 1, 2, 3, 4}
	idB := [5]uint32{1, 2, 3, 4, 5}
	expected := 31
	prefixMatch := DistPrefixLength(idA, idB)
	if prefixMatch != expected {
		log.Printf("[%s] - expected prefix length of %d, received %d\n", testName, expected, prefixMatch)
		t.Fail()
	}
}

func TestDistPrefixLength2(t *testing.T) {
	testName := "TestDistPrefixLength2"
	idA := [5]uint32{0, 1, 2, 3, 4}
	idB := [5]uint32{0, 2, 3, 4, 5}
	expected := 62
	prefixMatch := DistPrefixLength(idA, idB)
	if prefixMatch != expected {
		log.Printf("[%s] - expected prefix length of %d, received %d\n", testName, expected, prefixMatch)
		t.Fail()
	}
}

func TestDistPrefixLength3(t *testing.T) {
	testName := "TestDistPrefixLength3"
	idA := [5]uint32{0, 0, 0, 0, 0}
	idB := [5]uint32{0, 2, 3, 4, 5}
	expected := 62
	prefixMatch := DistPrefixLength(idA, idB)
	if prefixMatch != expected {
		log.Printf("[%s] - expected prefix length of %d, received %d\n", testName, expected, prefixMatch)
		t.Fail()
	}
}
func TestDistPrefixLength4(t *testing.T) {
	testName := "TestDistPrefixLength3"
	idA := [5]uint32{0, 0, 0, 0, 0}
	idB := [5]uint32{0, 0, 0, 0, 0}
	expected := 160
	prefixMatch := DistPrefixLength(idA, idB)
	if prefixMatch != expected {
		log.Printf("[%s] - expected prefix length of %d, received %d\n", testName, expected, prefixMatch)
		t.Fail()
	}
}

func TestSortSliceByDistance(t *testing.T) {
	verbose := true
	testName := "TestSortSliceByDistance"
	target := [5]uint32{0, 0, 0, 0, 0}
	input := make([]kademlia.Contact, 0)
	nodeA := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0})
	nodeB := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 1})
	nodeC := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 5, 0})
	nodeD := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 5, 0, 0})
	nodeE := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0})
	nodeF := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 1, 0, 0, 0})
	nodeG := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{1, 0, 0, 0, 0})
	nodeH := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{10, 92, 23, 233, 0})
	nodeI := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 99, 32, 0, 0})
	nodeJ := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 10, 1})
	nodeK := kademlia.NewContact([4]byte{0, 0, 0, 0}, [5]uint32{1, 1, 1, 1, 1})
	input = append(input, nodeA)
	input = append(input, nodeB)
	input = append(input, nodeC)
	input = append(input, nodeD)
	input = append(input, nodeE)
	input = append(input, nodeF)
	input = append(input, nodeG)
	input = append(input, nodeH)
	input = append(input, nodeI)
	input = append(input, nodeJ)
	input = append(input, nodeK)

	if verbose {
		s := fmt.Sprintf("[%s]\nbefore sort:\n", testName)
		for _, v := range input {
			s += fmt.Sprintf("node - %v, relative distance %d\n", v, RelativeDistance(v.ID(), target))
		}
		log.Println(s)
	}

	SortContactsByDistance(&input, target)

	if verbose {
		res := fmt.Sprintf("[%s]\nafter sort:\n", testName)
		for _, v := range input {
			res += fmt.Sprintf("node - %v, relative distance %d\n", v, RelativeDistance(v.ID(), target))
		}
		log.Println(res)
	}

	for i := 0; i < len(input)-1; i++ {
		ok, _ := CloserNode(input[i].ID(), input[i+1].ID(), target)
		if !ok {
			t.Fail()

		}
	}
}
