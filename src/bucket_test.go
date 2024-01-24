package main

import (
	"fmt"
	"testing"
)

func TestFillNewBucket(t *testing.T) {
	var testBucket bucket = NewBucket()
	err := testBucket.AddContact("127.0.0.1", 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact("127.0.0.2", 80, [5]uint32{6, 7, 8, 9, 10})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact("127.0.0.3", 80, [5]uint32{11, 12, 13, 14, 15})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact("127.0.0.4", 80, [5]uint32{16, 17, 18, 19, 20})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact("127.0.0.5", 80, [5]uint32{21, 22, 23, 24, 25})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact("127.0.0.6", 80, [5]uint32{26, 27, 28, 29, 30})
	if err == nil {
		fmt.Println("[TestFillNewBucket] - expected full bucket error, bucket contains:")
		for e := testBucket.content.Front(); e != nil; e = e.Next() {
			fmt.Printf("found element %+v\n", e)
		}
		t.FailNow()
	}
}

func TestDoubbleAddBucket(t *testing.T) {
	var testBucket bucket = NewBucket()
	err := testBucket.AddContact("127.0.0.1", 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Println("[TestDoubbleAddBucket] - unexpected error on first contact addition")
	}
	err = testBucket.AddContact("127.0.0.1", 80, [5]uint32{1, 2, 3, 4, 5})
	if err == nil {
		fmt.Println("[TestDoubbleAddBucket] - uncaught doubble add of contact")
		t.FailNow()
	}

}
