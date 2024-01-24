package main

import (
	"fmt"
	"testing"
)

func TestFillNewBucket(t *testing.T) {
	var testBucket bucket = NewBucket()
	contact1, err := NewContact("127.0.0.1", 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - invalid contact construction: ", err.Error())
	}
	contact2, err := NewContact("127.0.0.2", 80, [5]uint32{6, 7, 8, 9, 10})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - invalid contact construction: ", err.Error())
	}
	contact3, err := NewContact("127.0.0.3", 80, [5]uint32{11, 12, 13, 14, 15})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - invalid contact construction: ", err.Error())
	}
	contact4, err := NewContact("127.0.0.4", 80, [5]uint32{16, 17, 18, 19, 20})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - invalid contact construction: ", err.Error())
	}
	contact5, err := NewContact("127.0.0.5", 80, [5]uint32{21, 22, 23, 24, 25})
	overflowContact, err := NewContact("127.0.0.6", 80, [5]uint32{26, 27, 28, 29, 30})
	if err != nil {
		fmt.Println("[TestFillNewBucket] - invalid contact construction: ", err.Error())
	}

	err = testBucket.AddContact(contact1)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact(contact5)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}

	err = testBucket.AddContact(overflowContact)
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
	testContact, err := NewContact("127.0.0.1", 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Println("[TestDoubbleAddBucket] - invalid contact construction: ", err.Error())
	}
	bufferContact, err := NewContact("127.0.0.2", 80, [5]uint32{6, 7, 8, 9, 10})
	if err != nil {
		fmt.Println("[TestDoubbleAddBucket] - invalid contact construction: ", err.Error())
	}

	err = testBucket.AddContact(testContact)
	if err != nil {
		fmt.Println("[TestDoubbleAddBucket] - unexpected error when adding contact: ", err.Error())
	}
	err = testBucket.AddContact(bufferContact)
	if err != nil {
		fmt.Println("[TestDoubbleAddBucket] - unexpected error when adding contact: ", err.Error())
	}
	err = testBucket.AddContact(testContact)
	if err != nil {
		fmt.Println("[TestDoubbleAddBucket] - unexpected error when adding existing contact: ", err.Error())
		t.FailNow()
	}

	var lastSeen contact = testBucket.content.Back().Value.(contact)
	if lastSeen.ID() != testContact.ID() {
		fmt.Printf("[TestDoubleAddBucket] - expected last seen contact to have id: %+v, found node ID: %+v\n", testContact.ID(), lastSeen.ID())
	}
}

func TestRemoveHeadContact(t *testing.T) {
	var testBucket bucket = NewBucket()
	contact1, err := NewContact("127.0.0.1", 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Println("[] - invalid contact construction: ", err.Error())
	}
	contact2, err := NewContact("127.0.0.2", 80, [5]uint32{6, 7, 8, 9, 10})
	if err != nil {
		fmt.Println("[] - invalid contact construction: ", err.Error())
	}
	contact3, err := NewContact("127.0.0.3", 80, [5]uint32{11, 12, 13, 14, 15})
	if err != nil {
		fmt.Println("[] - invalid contact construction: ", err.Error())
	}
	contact4, err := NewContact("127.0.0.4", 80, [5]uint32{16, 17, 18, 19, 20})
	if err != nil {
		fmt.Println("[] - invalid contact construction: ", err.Error())
	}

	err = testBucket.AddContact(contact1)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}

	err = testBucket.RemoveContact(contact1)
	if err != nil {
		fmt.Println("[TestRemoveHeadContact] - failed to remove contact, error: ", err.Error())
	}
	head := testBucket.content.Front().Value.(contact)
	if head.ID() == contact1.ID() {
		fmt.Printf("[TestRemoveHeadContact] - failed to remove contact, %+v, from bucket\n", contact1)
	}
}

func TestRemoveCenterContact(t *testing.T) {
	var testBucket bucket = NewBucket()
	contact1, err := NewContact("127.0.0.1", 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Println("[] - invalid contact construction: ", err.Error())
	}
	contact2, err := NewContact("127.0.0.2", 80, [5]uint32{6, 7, 8, 9, 10})
	if err != nil {
		fmt.Println("[] - invalid contact construction: ", err.Error())
	}
	contact3, err := NewContact("127.0.0.3", 80, [5]uint32{11, 12, 13, 14, 15})
	if err != nil {
		fmt.Println("[] - invalid contact construction: ", err.Error())
	}
	contact4, err := NewContact("127.0.0.4", 80, [5]uint32{16, 17, 18, 19, 20})
	if err != nil {
		fmt.Println("[] - invalid contact construction: ", err.Error())
	}

	fmt.Println("print 1")
	err = testBucket.AddContact(contact1)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	fmt.Println("print 3")
	err = testBucket.AddContact(contact2)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	fmt.Println("print 4")
	err = testBucket.AddContact(contact3)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}
	fmt.Println("print 5")
	err = testBucket.AddContact(contact4)
	if err != nil {
		fmt.Println("[TestFillNewBucket] - ", err.Error())
	}

	fmt.Println("print 6")
	err = testBucket.RemoveContact(contact3)
	fmt.Println("print 2")
	if err != nil {
		fmt.Println("print 8")
		fmt.Println("[TestRemoveCenterContact] - failed to remove contact, error: ", err.Error())
	}

	fmt.Println("print 7")
	var counter int = 0
	fmt.Println(counter)
	for e := testBucket.content.Front(); e != nil; e = e.Next() {
		counter++
		fmt.Println(counter)
		elem, ok := e.Value.(contact)
		if !ok {
			fmt.Printf("[TestRemoveCenterContact] - bucket has been corrupted: expected contact, found: %+v\n", e)
		}
		if elem.ID() == contact3.ID() {
			fmt.Printf("[TestRemoveCenterContact] - failed to remove contact, %+v\n", contact3)
			t.FailNow()
		}
	}
}
