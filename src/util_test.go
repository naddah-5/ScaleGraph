package scalegraph


import (
	"container/list"
	"log"
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
	var testName string = "TestSortContactList"
	var verbose bool = false
	var testList *list.List = list.New()
	var targetNode [5]uint32 = [5]uint32{0, 0, 0, 0, 0}

	contact1, err := BuildContact([4]byte{127, 0, 0, 1}, 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact2, err := BuildContact([4]byte{127, 0, 0, 2}, 80, [5]uint32{6, 7, 8, 9, 10})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact3, err := BuildContact([4]byte{127, 0, 0, 3}, 80, [5]uint32{11, 12, 13, 14, 15})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact4, err := BuildContact([4]byte{127, 0, 0, 4}, 80, [5]uint32{16, 17, 18, 19, 20})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact5, err := BuildContact([4]byte{127, 0, 0, 5}, 80, [5]uint32{21, 22, 23, 24, 25})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact6, err := BuildContact([4]byte{127, 0, 0, 6}, 80, [5]uint32{26, 27, 28, 29, 30})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}

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

	SortByDistance(testList, targetNode)

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

	err := SortByDistance(testList, targetNode)
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

	contact1, err := BuildContact([4]byte{127, 0, 0, 1}, 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact2, err := BuildContact([4]byte{127, 0, 0, 2}, 80, [5]uint32{6, 7, 8, 9, 10})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact3, err := BuildContact([4]byte{127, 0, 0, 3}, 80, [5]uint32{11, 12, 13, 14, 15})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact4, err := BuildContact([4]byte{127, 0, 0, 4}, 80, [5]uint32{16, 17, 18, 19, 20})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact5, err := BuildContact([4]byte{127, 0, 0, 5}, 80, [5]uint32{21, 22, 23, 24, 25})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact6, err := BuildContact([4]byte{127, 0, 0, 6}, 80, [5]uint32{26, 27, 28, 29, 30})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact7, err := BuildContact([4]byte{127, 0, 0, 7}, 80, [5]uint32{31, 32, 33, 34, 35})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact8, err := BuildContact([4]byte{127, 0, 0, 8}, 80, [5]uint32{36, 37, 38, 39, 40})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact9, err := BuildContact([4]byte{127, 0, 0, 9}, 80, [5]uint32{41, 42, 43, 44, 45})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact10, err := BuildContact([4]byte{127, 0, 0, 10}, 80, [5]uint32{46, 47, 48, 49, 50})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact11, err := BuildContact([4]byte{127, 0, 0, 11}, 80, [5]uint32{51, 52, 53, 54, 55})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact12, err := BuildContact([4]byte{127, 0, 0, 12}, 80, [5]uint32{56, 57, 58, 59, 60})
	if err != nil {
		log.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}

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

	SortByDistance(testListA, target)
	SortByDistance(testListB, target)

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
