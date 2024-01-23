package main

import (
	"fmt"
	"testing"
)

func TestGenerateIDRandomness(t *testing.T) {
	var baseID [5]uint32
	baseID, _ = GenerateID()
	var confidence float64 = 1000000
	var coincidence float64 = 0
	var warning float64 = 0.0001
	for i := 0.0; i < confidence; i++ {
		sample, _ := GenerateID()
		if sample == baseID {
			coincidence++
		}
	}
	if coincidence/confidence >= warning {
		fmt.Printf("[TestGenerateIDRandomness] - WARNING: low probability event: %d, collisions in %d generations", int(coincidence), int(confidence))
		t.FailNow()
	}
}

func TestRelativeDistance(t *testing.T) {
	var idA [5]uint32 = [5]uint32{1, 2, 3, 4, 5} // 1 + 1 + 2 + 1 + 2
	var idB [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var expectedDist = 7
	res := RelativeDistance(&idA, &idB)
	if res != expectedDist {
		fmt.Printf("[TestRelativeDistance] - found relative distnace %v, expected relative distance %v\n", res, expectedDist)
		t.FailNow()
	}
}

func TestRelativeDistancePointing(t *testing.T) {
	var idA [5]uint32 = [5]uint32{1, 2, 3, 4, 5}
	var idB [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	RelativeDistance(&idA, &idB)
	if idA[0] != 1 {
		t.FailNow()
	}
}

func TestHammingDistance(t *testing.T) {
	var pointA uint32 = uint32(0b10101010101010101)
	var pointB uint32 = uint32(0b01010101010101010)
	var expectedHamDist int = 17
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		fmt.Printf("[TestHammingDistance] - found hamming distance %v, expected hamming distance %v\n", res, expectedHamDist)
		t.FailNow()
	}
}

func TestHammingDistance1(t *testing.T) {
	var pointA uint32 = uint32(0b10101010010101)
	var pointB uint32 = uint32(0b01010101010101010)
	var expectedHamDist int = 7
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		fmt.Printf("[TestHammingDistance1] - found hamming distance %v, expected hamming distance %v\n", res, expectedHamDist)
		t.FailNow()
	}
}

func TestHammingDistance2(t *testing.T) {
	var pointA uint32 = uint32(0b1)
	var pointB uint32 = uint32(0b0)
	var expectedHamDist int = 1
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		fmt.Printf("[TestHammingDistance1] - found hamming distance %v, expected hamming distance %v\n", res, expectedHamDist)
		t.FailNow()
	}
}

func TestHammingDistance3(t *testing.T) {
	var pointA uint32 = 213
	var pointB uint32 = 42
	var expectedHamDist int = 8
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		fmt.Printf("[TestHammingDistance3] - found hamming distance %v, expected hamming distance %v\n", res, expectedHamDist)
		t.FailNow()
	}
}

func TestHammingDistance4(t *testing.T) {
	var pointA uint32 = 0
	var tmp int = -1
	var pointB uint32 = uint32(tmp)
	var expectedHamDist int = 32
	res := hammingDistance(pointA, pointB)
	if res != expectedHamDist {
		fmt.Printf("[TestHammingDistance4] - found hamming distance %v, expected hamming distance %v\n", res, expectedHamDist)
		t.FailNow()
	}
}

func TestPrefixMatch(t *testing.T) {
	var idA [5]uint32 = [5]uint32{0, 1, 2, 3, 4}
	var idB [5]uint32 = [5]uint32{1, 2, 3, 4, 5}
	prefixMatch := prefixLength(idA, idB)
	fmt.Printf("found prefix of length %d\n", prefixMatch)

}
