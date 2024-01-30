package main

import (
	"testing"
	"container/list"
	"fmt"
)

func TestRelativeDistance(t *testing.T) {
	var testName string = "TestRelativeDistance"
	var idA [5]uint32 = [5]uint32{1, 2, 3, 4, 5} // 1 + 1 + 2 + 1 + 2
	var idB [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var expectedDist = 7
	res := RelativeDistance(idA, idB)
	if res != expectedDist {
		fmt.Printf("[%s] - found relative distnace %v, expected relative distance %v\n", testName, res, expectedDist)
		t.FailNow()
	}
}

func TestRelativeDistancePointing(t *testing.T) {
	var testName string = "TestRelativeDistancePointing"
	var idA [5]uint32 = [5]uint32{1, 2, 3, 4, 5}
	var idB [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	RelativeDistance(idA, idB)
	if idA[0] != 1 {
		fmt.Printf("[%s] - expected a relative distance of 1, found %d", testName, idA[0])
		t.FailNow()
	}
}

func TestHammingDistance(t *testing.T) {
	var testName string = "TestHammingDistance"
	var pointA uint32 = uint32(0b10101010101010101)
	var pointB uint32 = uint32(0b01010101010101010)
	var expectedHamDist int = 17
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		fmt.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.FailNow()
	}
}

func TestHammingDistance1(t *testing.T) {
	var testName string = "TestHammingdistance1"
	var pointA uint32 = uint32(0b10101010010101)
	var pointB uint32 = uint32(0b01010101010101010)
	var expectedHamDist int = 7
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		fmt.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.FailNow()
	}
}

func TestHammingDistance2(t *testing.T) {
	var testName string = "TestHammingDistance2"
	var pointA uint32 = uint32(0b1)
	var pointB uint32 = uint32(0b0)
	var expectedHamDist int = 1
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		fmt.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.FailNow()
	}
}

func TestHammingDistance3(t *testing.T) {
	var testName string = "TestHammingDistance3"
	var pointA uint32 = 213
	var pointB uint32 = 42
	var expectedHamDist int = 8
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		fmt.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.FailNow()
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
		fmt.Printf("[%s] - found hamming distance %v, expected hamming distance %v\n", testName, res, expectedHamDist)
		t.FailNow()
	}
}

func TestPrefixMatch(t *testing.T) {
	var testName string = "TestPrefixMatch"
	var idA [5]uint32 = [5]uint32{0, 1, 2, 3, 4}
	var idB [5]uint32 = [5]uint32{1, 2, 3, 4, 5}
	const expected int = 31
	prefixMatch := DistPrefixLength(idA, idB)
	if prefixMatch != expected {
		fmt.Printf("[%s] - expected prefix length of %d, received %d\n", testName, expected, prefixMatch)
		t.FailNow()
	}

}

func TestPrefixMatch1(t *testing.T) {
	var testName string = "TestPrefixMatch1"
	var idA [5]uint32 = [5]uint32{0, 1, 2, 3, 4}
	var idB [5]uint32 = [5]uint32{0, 2, 3, 4, 5}
	const expected int = 62
	prefixMatch := DistPrefixLength(idA, idB)
	if prefixMatch != expected {
		fmt.Printf("[%s] - expected prefix length of %d, received %d", testName, expected, prefixMatch)
		t.FailNow()
	}
}

func TestPrefixMatch2(t *testing.T) {
	var testName string = "TestPrefixMatch2"
	var idA [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var idB [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	prefixMatch := DistPrefixLength(idA, idB)
	const expected int = 160
	if prefixMatch != expected {
		fmt.Printf("[%s] - expected prefix length of %d, received %d", testName, expected, prefixMatch)
		t.FailNow()
	}
}

func TestCloserNode(t *testing.T) {
	var idA [5]uint32 = [5]uint32{0, 1, 0, 0, 0}
	var idB [5]uint32 = [5]uint32{5, 11, 12, 13, 9}
	var target [5]uint32 = [5]uint32{10, 11, 12, 13, 14}

	res := CloserNode(idA, idB, target)
	if res {
		fmt.Printf("expected B - %v, to be closer to target - %v, than A - %v\n", idB, target, idA)
		t.FailNow()
	}
}

func TestCloserNodeEqualDistance(t *testing.T) {
	var node [5]uint32 = [5]uint32{123, 4124, 213, 2312, 231}
	var target [5]uint32 = [5]uint32{0, 0, 0, 0, 0}

	res := CloserNode(node, node, target)
	if !res {
		fmt.Printf("expected res to be true since node A == node B\n")
		t.FailNow()
	}
}

func TestSortContactList(t *testing.T) {
	var testName string = "TestSortContactList"
	var testList *list.List = list.New()
	var targetNode [5]uint32 = [5]uint32{11, 12, 13, 14, 15}

	contact1, err := BuildContact("127.0.0.1", 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact2, err := BuildContact("127.0.0.2", 80, [5]uint32{6, 7, 8, 9, 10})
	if err != nil {
		fmt.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact3, err := BuildContact("127.0.0.3", 80, [5]uint32{11, 12, 13, 14, 15})
	if err != nil {
		fmt.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact4, err := BuildContact("127.0.0.4", 80, [5]uint32{16, 17, 18, 19, 20})
	if err != nil {
		fmt.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact5, err := BuildContact("127.0.0.5", 80, [5]uint32{21, 22, 23, 24, 25})
	if err != nil {
		fmt.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	contact6, err := BuildContact("127.0.0.6", 80, [5]uint32{26, 27, 28, 29, 30})
	if err != nil {
		fmt.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}

	testList.PushFront(contact1)
	testList.PushFront(contact4)
	testList.PushFront(contact3)
	testList.PushFront(contact5)
	testList.PushFront(contact2)
	testList.PushFront(contact6)

	SortByDistance(testList, targetNode)

	var prevDist int = 0
	for e := testList.Front(); e != nil; e = e.Next() {
		elem := e.Value.(contact)
		relDist := RelativeDistance(elem.ID(), targetNode)
		if relDist < prevDist {
			t.FailNow()
		}
		prevDist = relDist
	}
}
