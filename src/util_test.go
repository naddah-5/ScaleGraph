package scalegraph

import (
	"container/list"
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestRelativeDistance(t *testing.T) {
	var testName string = "TestRelativeDistance"
	var idA [5]uint32 = [5]uint32{1, 2, 3, 4, 5} // 1 + 1 + 2 + 1 + 2
	var idB [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var expectedDist = 7
	res := RelativeDistance(idA, idB)
	if res != expectedDist {
		log.Printf("[%s] - found relative distnace %v, expected relative distance %v\n", testName, res, expectedDist)
		t.Fail()
	}
}

func TestRelativeDistancePointing(t *testing.T) {
	var testName string = "TestRelativeDistancePointing"
	var idA [5]uint32 = [5]uint32{1, 2, 3, 4, 5}
	var idB [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	RelativeDistance(idA, idB)
	if idA[0] != 1 {
		log.Printf("[%s] - expected a relative distance of 1, found %d", testName, idA[0])
		t.Fail()
	}
}

func TestHammingDistance(t *testing.T) {
	var testName string = "TestHammingDistance"
	var pointA uint32 = uint32(0b10101010101010101)
	var pointB uint32 = uint32(0b01010101010101010)
	var expectedHamDist int = 17
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		log.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.Fail()
	}
}

func TestHammingDistance1(t *testing.T) {
	var testName string = "TestHammingdistance1"
	var pointA uint32 = uint32(0b10101010010101)
	var pointB uint32 = uint32(0b01010101010101010)
	var expectedHamDist int = 7
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		log.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.Fail()
	}
}

func TestHammingDistance2(t *testing.T) {
	var testName string = "TestHammingDistance2"
	var pointA uint32 = uint32(0b1)
	var pointB uint32 = uint32(0b0)
	var expectedHamDist int = 1
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		log.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.Fail()
	}
}

func TestHammingDistance3(t *testing.T) {
	var testName string = "TestHammingDistance3"
	var pointA uint32 = 213
	var pointB uint32 = 42
	var expectedHamDist int = 8
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		log.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.Fail()
	}
}

func TestHammingDistance4(t *testing.T) {
	var testName string = "TestHammingDistance4"
	var pointA uint32 = 0
	var tmp int = -1
	var pointB uint32 = uint32(tmp)
	var expectedHamDist int = 32
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		log.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.Fail()
	}
}

func TestPrefixMatch(t *testing.T) {
	var testName string = "TestPrefixMatch"
	var idA [5]uint32 = [5]uint32{0, 1, 2, 3, 4}
	var idB [5]uint32 = [5]uint32{1, 2, 3, 4, 5}
	const expected int = 31
	prefixMatch := DistPrefixLength(idA, idB)
	if prefixMatch != expected {
		log.Printf("[%s] - expected prefix length of %d, received %d\n", testName, expected, prefixMatch)
		t.Fail()
	}

}

func TestPrefixMatch1(t *testing.T) {
	var testName string = "TestPrefixMatch1"
	var idA [5]uint32 = [5]uint32{0, 1, 2, 3, 4}
	var idB [5]uint32 = [5]uint32{0, 2, 3, 4, 5}
	const expected int = 62
	prefixMatch := DistPrefixLength(idA, idB)
	if prefixMatch != expected {
		log.Printf("[%s] - expected prefix length of %d, received %d", testName, expected, prefixMatch)
		t.Fail()
	}
}

func TestPrefixMatch2(t *testing.T) {
	var testName string = "TestPrefixMatch2"
	var idA [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var idB [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	prefixMatch := DistPrefixLength(idA, idB)
	const expected int = 160
	if prefixMatch != expected {
		log.Printf("[%s] - expected prefix length of %d, received %d", testName, expected, prefixMatch)
		t.Fail()
	}
}

func TestCloserNode(t *testing.T) {
	var idA [5]uint32 = [5]uint32{0, 1, 0, 0, 0}
	var idB [5]uint32 = [5]uint32{5, 11, 12, 13, 9}
	var target [5]uint32 = [5]uint32{10, 11, 12, 13, 14}

	res := CloserNode(idA, idB, target)
	if res {
		log.Printf("expected B - %v, to be closer to target - %v, than A - %v\n", idB, target, idA)
		t.Fail()
	}
}

func TestCloserNodeEqualDistance(t *testing.T) {
	var node [5]uint32 = [5]uint32{123, 4124, 213, 2312, 231}
	var target [5]uint32 = [5]uint32{0, 0, 0, 0, 0}

	res := CloserNode(node, node, target)
	if !res {
		log.Printf("expected res to be true since node A == node B\n")
		t.Fail()
	}
}

func TestSortContactList(t *testing.T) {
	var _ string = "TestSortContactList"
	var verbose bool = false
	var testList *list.List = list.New()
	var targetNode [5]uint32 = [5]uint32{0, 0, 0, 0, 0}

	contact1 := BuildContact([4]byte{127, 0, 0, 1}, [5]uint32{1, 2, 3, 4, 5})
	contact2 := BuildContact([4]byte{127, 0, 0, 2}, [5]uint32{6, 7, 8, 9, 10})
	contact3 := BuildContact([4]byte{127, 0, 0, 3}, [5]uint32{11, 12, 13, 14, 15})
	contact4 := BuildContact([4]byte{127, 0, 0, 4}, [5]uint32{16, 17, 18, 19, 20})
	contact5 := BuildContact([4]byte{127, 0, 0, 5}, [5]uint32{21, 22, 23, 24, 25})
	contact6 := BuildContact([4]byte{127, 0, 0, 6}, [5]uint32{26, 27, 28, 29, 30})

	testList.PushFront(contact1)
	testList.PushFront(contact4)
	testList.PushFront(contact3)
	testList.PushFront(contact5)
	testList.PushFront(contact2)
	testList.PushFront(contact6)

	if verbose {
		log.Printf("list before sorting:\n")
		for e := testList.Front(); e != nil; e = e.Next() {
			elem := e.Value.(contact)
			var relDist int = RelativeDistance(elem.ID(), targetNode)
			log.Printf("elem: %+v, relDist: %d\n", elem, relDist)
		}
	}

	SortListByDistance(testList, targetNode)

	if verbose {
		log.Printf("list after sorting:\n")
		for e := testList.Front(); e != nil; e = e.Next() {
			elem := e.Value.(contact)
			var relDist int = RelativeDistance(elem.ID(), targetNode)
			log.Printf("elem: %+v, relDist: %d\n", elem, relDist)
		}
	}

	var prevDist int = 0
	for e := testList.Front(); e != nil; e = e.Next() {
		elem := e.Value.(contact)
		relDist := RelativeDistance(elem.ID(), targetNode)
		if relDist < prevDist {
			t.Fail()
		}
		prevDist = relDist
	}
}

func TestSortEmptyList(t *testing.T) {
	var testName string = "TestSortEmptyList"
	var verbose bool = false
	var testList *list.List = list.New()
	var targetNode [5]uint32 = [5]uint32{0, 0, 0, 0, 0}

	if verbose {
		log.Printf("list before sorting:\n")
		for e := testList.Front(); e != nil; e = e.Next() {
			elem := e.Value
			log.Printf("elem: %+v\n", elem)
		}
	}

	err := SortListByDistance(testList, targetNode)
	if err != nil {
		log.Printf("[%s] - failed to properly sort empty list", testName)
		t.Fail()
	}

	if verbose {
		log.Printf("list after sorting:\n")
		for e := testList.Front(); e != nil; e = e.Next() {
			elem := e.Value
			log.Printf("elem: %+v", elem)
		}
	}
}

func TestMergeByDistance(t *testing.T) {
	var testName string = "TestMergeByDistance"
	var target [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var testListA *list.List = list.New()
	var testListB *list.List = list.New()

	contact1 := BuildContact([4]byte{127, 0, 0, 1}, [5]uint32{1, 2, 3, 4, 5})
	contact2 := BuildContact([4]byte{127, 0, 0, 2}, [5]uint32{6, 7, 8, 9, 10})
	contact3 := BuildContact([4]byte{127, 0, 0, 3}, [5]uint32{11, 12, 13, 14, 15})
	contact4 := BuildContact([4]byte{127, 0, 0, 4}, [5]uint32{16, 17, 18, 19, 20})
	contact5 := BuildContact([4]byte{127, 0, 0, 5}, [5]uint32{21, 22, 23, 24, 25})
	contact6 := BuildContact([4]byte{127, 0, 0, 6}, [5]uint32{26, 27, 28, 29, 30})
	contact7 := BuildContact([4]byte{127, 0, 0, 7}, [5]uint32{31, 32, 33, 34, 35})
	contact8 := BuildContact([4]byte{127, 0, 0, 8}, [5]uint32{36, 37, 38, 39, 40})
	contact9 := BuildContact([4]byte{127, 0, 0, 9}, [5]uint32{41, 42, 43, 44, 45})
	contact10 := BuildContact([4]byte{127, 0, 0, 10}, [5]uint32{46, 47, 48, 49, 50})
	contact11 := BuildContact([4]byte{127, 0, 0, 11}, [5]uint32{51, 52, 53, 54, 55})
	contact12 := BuildContact([4]byte{127, 0, 0, 12}, [5]uint32{56, 57, 58, 59, 60})

	testListA.PushFront(contact1)
	testListA.PushFront(contact4)
	testListA.PushFront(contact3)
	testListA.PushFront(contact5)
	testListA.PushFront(contact2)
	testListA.PushFront(contact6)
	testListB.PushFront(contact7)
	testListB.PushFront(contact8)
	testListB.PushFront(contact9)
	testListB.PushFront(contact10)
	testListB.PushFront(contact11)
	testListB.PushFront(contact12)

	SortListByDistance(testListA, target)
	SortListByDistance(testListB, target)

	res, err := MergeByDistance(testListA, testListB, target)
	if err != nil {
		log.Println(err.Error())
	}

	var prevDist int = 0
	var relDist int = 0
	for e := res.Front(); e != nil; e = e.Next() {
		elem := e.Value.(contact)
		relDist = RelativeDistance(elem.ID(), target)
		if relDist < prevDist {
			prevElem := e.Prev().Value.(contact)
			log.Printf("[%s] - resulting list is not propperly sorted: %+v with relative distance %d found after %+v with relative distance %d", testName, elem, relDist, prevElem, prevDist)
			t.Fail()
			prevDist = relDist
		}

	}
}

func TestCompareHash(t *testing.T) {
	data := GenerateID()
	hasher := sha256.New()
	for i := range data {
		tmp := strconv.FormatUint(uint64(data[i]), 10)
		hasher.Write([]byte(tmp))
	}
	hash := hasher.Sum(nil)
	res := CompareHash(hash, hash)
	if !res {
		log.Println("CompareHash incorrect")
		t.Fail()
	}
}

func TestCompareHashFail(t *testing.T) {
	data := GenerateID()
	hasher := sha256.New()
	for i := range data {
		tmp := strconv.FormatUint(uint64(data[i]), 10)
		hasher.Write([]byte(tmp))
	}
	hash := hasher.Sum(nil)
	junk := "this is junk data"
	hasher.Write([]byte(junk))
	badHash := hasher.Sum(nil)
	res := CompareHash(hash, badHash)
	if res {
		log.Println("CompareHash failed to detect hash missmatch")
		log.Printf("first hash: %+v\n", hash)
		log.Printf("second hash: %+v\n", badHash)
		t.Fail()
	}
}

func TestCompareHashLength(t *testing.T) {
	data := GenerateID()
	hasher := sha256.New()
	for i := range data {
		tmp := strconv.FormatUint(uint64(data[i]), 10)
		hasher.Write([]byte(tmp))
	}
	hash := hasher.Sum(nil)
	res := CompareHash(hash, hash[:len(hash)/2])
	if res {
		log.Println("CompareHash incorrect with length missmatch")
		t.Fail()
	}

}

func TestCompareSlice(t *testing.T) {
	testName := "TestCompareSlice"
	testList := make([]contact, 100)
	for i := range testList {
		testList[i] = NewRandomContact()
	}
	shouldBeTrue := CompareContactSlice(testList, testList)
	if !shouldBeTrue {
		log.Printf("[%+v] - matching list failed", testName)
		t.Fail()
	}

	testList1 := make([]contact, 100)
	for i := range testList {
		testList1[i] = NewRandomContact()
	}
	shouldBeFalse := CompareContactSlice(testList, testList1)
	if shouldBeFalse {
		log.Printf("[%+v] - matching mismatched list should fail", testName)
		t.Fail()
	}

	shouldBeFalse = CompareContactSlice(testList, testList[:len(testList)/2])
	if shouldBeFalse {
		log.Printf("[%+v] - matching list of different length should be false", testName)
		t.Fail()
	}

	shouldBeTrue = CompareContactSlice(make([]contact, 0), make([]contact, 0))
	if !shouldBeTrue {
		log.Printf("[%+v] - matching of empty lists should be true", testName)
		t.Fail()
	}
}

func TestLargerNode(t *testing.T) {
	testName := "TestLargerNode"
	nodeA := [5]uint32{0, 0, 1, 0, 0}
	nodeB := [5]uint32{0, 0, 0, 0, 1}
	nodeC := [5]uint32{0, 0, 0, 1, 0}
	nodeD := [5]uint32{0, 0, 1, 0, 0}
	nodeE := [5]uint32{0, 1, 0, 0, 0}
	nodeF := [5]uint32{1, 0, 0, 0, 0}

	nodeG := [5]uint32{0, 23, 153, 123, 0}
	nodeH := [5]uint32{0, 23, 73, 123, 0}

	if !LargerNode(nodeA, nodeB) {
		log.Printf("[%s] - node A should be larger than node B", testName)
		t.Fail()
	}

	if !LargerNode(nodeA, nodeC) {
		log.Printf("[%s] - node A should be larger than node C", testName)
		t.Fail()
	}
	if LargerNode(nodeA, nodeD) {
		log.Printf("[%s] - node A should not be larger than node D", testName)
		t.Fail()
	}
	if LargerNode(nodeA, nodeE) {
		log.Printf("[%s] - node A should not be larger than node E", testName)
		t.Fail()
	}
	if LargerNode(nodeA, nodeF) {
		log.Printf("[%s] - node A should not be larger than node F", testName)
		t.Fail()
	}
	if !LargerNode(nodeG, nodeH) {
		log.Printf("[%s] - node G should be larger than node H", testName)
		t.Fail()
	}
}

func TestSortSliceByDistance(t *testing.T) {
	verbose := false
	testName := "TestSortSliceByDistance"
	target := [5]uint32{0, 0, 0, 0, 0}
	input := make([]contact, 0)
	nodeA := contact{[4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0}}
	nodeB := contact{[4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 1}}
	nodeC := contact{[4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 5, 0}}
	nodeD := contact{[4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 5, 0, 0}}
	nodeE := contact{[4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0}}
	nodeF := contact{[4]byte{0, 0, 0, 0}, [5]uint32{0, 1, 0, 0, 0}}
	nodeG := contact{[4]byte{0, 0, 0, 0}, [5]uint32{1, 0, 0, 0, 0}}
	nodeH := contact{[4]byte{0, 0, 0, 0}, [5]uint32{10, 92, 23, 233, 0}}
	nodeI := contact{[4]byte{0, 0, 0, 0}, [5]uint32{0, 99, 32, 0, 0}}
	nodeJ := contact{[4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 10, 1}}
	nodeK := contact{[4]byte{0, 0, 0, 0}, [5]uint32{1, 1, 1, 1, 1}}
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

	SortSliceByDistance(&input, target)

	if verbose {
		res := fmt.Sprintf("[%s]\nafter sort:\n", testName)
		for _, v := range input {
			res += fmt.Sprintf("node - %v, relative distance %d\n", v, RelativeDistance(v.ID(), target))
		}
		log.Println(res)
	}

	for i := 0; i < len(input)-1; i++ {
		if RelativeDistance(input[i].ID(), target) > RelativeDistance(input[i+1].ID(), target) {
			t.Fail()
			log.Printf("[%s] - incorrect sorting", testName)
		}
	}
}
