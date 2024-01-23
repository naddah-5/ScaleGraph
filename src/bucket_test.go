package main

import (
	"testing"
)

func TestNewBucket(t *testing.T) {
	var testBucket *Bucket = NewBucket()
	testBucket.AddContact("127.0.0.1", 80, [5]uint32{0, 0, 0, 0, 0})
	testBucket.sip()
}
