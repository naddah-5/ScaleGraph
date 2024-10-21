package kademlia

import (
	"fmt"
	"log"
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
	verbose := false
	testName := "TestSortSliceByDistance"
	target := [5]uint32{0, 0, 0, 0, 0}
	input := make([]Contact, 0)
	nodeA := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0})
	nodeB := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 1})
	nodeC := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 5, 0})
	nodeD := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 5, 0, 0})
	nodeE := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0})
	nodeF := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 1, 0, 0, 0})
	nodeG := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{1, 0, 0, 0, 0})
	nodeH := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{10, 92, 23, 233, 0})
	nodeI := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 99, 32, 0, 0})
	nodeJ := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 10, 1})
	nodeK := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{1, 1, 1, 1, 1})
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
		nodeA := input[i]
		nodeB := input[i+1]
		ok := CloserNode(nodeA.ID(), nodeB.ID(), target)
		if !ok && !EquiDistantNode(nodeA.ID(), nodeB.ID(), target){
			log.Printf("[%s] - node %v should be before node %v", testName, nodeA, nodeB)
			t.Fail()

		}
	}
}

func TestGreaterNode(t *testing.T) {
	testName := "TestGreaterNode"
	nodeA := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{10, 92, 23, 233, 0})
	nodeB := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{1, 1, 1, 1, 1})
	nodeC := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{1, 0, 0, 0, 0})
	nodeD := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 99, 32, 0, 0})
	nodeE := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 1, 0, 0, 0})
	nodeF := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 5, 0, 0})
	nodeG := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0})
	nodeH := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0})
	nodeI := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 10, 1})
	nodeJ := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 5, 0})
	nodeK := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 1})
	if !LargerNode(nodeA.ID(), nodeB.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeA, nodeB)
		t.Fail()
	}
	if !LargerNode(nodeB.ID(), nodeC.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeB, nodeC)
		t.Fail()
	}
	if !LargerNode(nodeC.ID(), nodeD.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeC, nodeD)
		t.Fail()
	}
	if !LargerNode(nodeD.ID(), nodeE.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeD, nodeE)
		t.Fail()
	}
	if !LargerNode(nodeE.ID(), nodeF.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeE, nodeF)
		t.Fail()
	}
	if !LargerNode(nodeF.ID(), nodeG.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeF, nodeG)
		t.Fail()
	}
	if LargerNode(nodeG.ID(), nodeH.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeG, nodeH)
		t.Fail()
	}
	if !LargerNode(nodeH.ID(), nodeI.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeH, nodeI)
		t.Fail()
	}
	if !LargerNode(nodeI.ID(), nodeJ.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeI, nodeJ)
		t.Fail()
	}
	if !LargerNode(nodeJ.ID(), nodeK.ID()) {
		log.Printf("[%s] - incorrect size assertion node %v is larger than node %v", testName, nodeJ, nodeK)
		t.Fail()
	}
}
