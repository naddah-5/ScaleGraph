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
