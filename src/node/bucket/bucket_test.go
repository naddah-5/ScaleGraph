package bucket

import (
	"fmt"
	"testing"
)

func TestNewBucket(t *testing.T) {
	var testBucket Bucket = NewBucket(2)
	fmt.Println(testBucket)
}
