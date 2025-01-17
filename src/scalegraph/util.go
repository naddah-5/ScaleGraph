package scalegraph

import (
	"errors"
	"math/rand"
)

// returns a pseudo-random uint32 in the range (min, max]
func RandU32(min uint32, max uint32) (uint32, error) {
	if min >= max {
		return 0, errors.New("invalid range")
	}
	x := rand.Uint32()
	x %= (max - min)
	x += min
	return x, nil
}

// Returns a randomly generated id.
func RandomID() [5]uint32 {
	var res [5]uint32
	for res == [5]uint32{0, 0, 0, 0, 0} {
		for i := 0; i < 5; i++ {
			res[i] = rand.Uint32()
		}
	}
	return res
}
