package scalegraph

import (
	"testing"
)

func TestSanity(t *testing.T) {
	testConsensus := NewConsensus(1, []byte{0, 1, 2, 3})
	sign := NewSign([5]uint32{0, 1, 2, 3, 4})
	testConsensus.Approved(sign)
	
	if len(testConsensus.validation) == 0 {
		t.FailNow()
	}
}
