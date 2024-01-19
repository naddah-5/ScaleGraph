package main

import (
	"fmt"
	"testing"
)

func TestBucketVolume(t *testing.T) {
	var testBucket bucket = NewBucket()
	var expectedNrNodes int = BUCKETVOLUME * KBUCKETS
	var foundNrNodes int = 0

	for i := 0; i < len(testBucket.buckets); i++ {
		for j := 0; j < len(testBucket.buckets[i]); j++ {
			foundNrNodes++
		}
	}
	if expectedNrNodes != foundNrNodes {
		fmt.Println("[TestBucketVolume] - expected to find : ", expectedNrNodes, " nodes, found : ", foundNrNodes, " nodes")
		t.FailNow()
	}
}
