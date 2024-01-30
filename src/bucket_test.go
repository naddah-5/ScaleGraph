package main

import (
	"fmt"
	"testing"
)

func TestFillNewBucket(t *testing.T) {
	var testName string = "TestFillBucket"
	var testBucket bucket = NewBucket()
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
	overflowContact, err := BuildContact("127.0.0.6", 80, [5]uint32{26, 27, 28, 29, 30})
	if err != nil {
		fmt.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}

	err = testBucket.AddContact(contact1)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact5)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	err = testBucket.AddContact(overflowContact)
	if err == nil {
		fmt.Printf("[%s] - expected full bucket error", testName)
		t.FailNow()
	}
}

func TestDoubbleAddBucket(t *testing.T) {
	var testName string = "TestDoubbleAddBucket"
	var testBucket bucket = NewBucket()
	testContact, err := BuildContact("127.0.0.1", 80, [5]uint32{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()
	}
	bufferContact, err := BuildContact("127.0.0.2", 80, [5]uint32{6, 7, 8, 9, 10})
	if err != nil {
		fmt.Printf("[%s] - invalid contact construction: %s", testName, err.Error())
		t.FailNow()

	}

	err = testBucket.AddContact(testContact)
	if err != nil {
		fmt.Printf("[%s] - unexpected error when adding contact: %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(bufferContact)
	if err != nil {
	}
	err = testBucket.AddContact(testContact)
	if err != nil {
		fmt.Printf("[%s] - unexpected error when adding existing contact: %s", testName, err.Error())
		t.FailNow()
	}

	var lastSeen contact = testBucket.content.Back().Value.(contact)
	if lastSeen.ID() != testContact.ID() {
		fmt.Printf("[%s] - expected last seen contact to have id: %+v, found node ID: %+v\n", testName, testContact.ID(), lastSeen.ID())
	}
}

func TestRemoveHeadContact(t *testing.T) {
	var testName string = "TestRemoveHeadContact"
	var testBucket bucket = NewBucket()
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

	err = testBucket.AddContact(contact1)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	err = testBucket.RemoveContact(contact1)
	if err != nil {
		fmt.Printf("[%s] - failed to remove contact, error: %s", testName, err.Error())
		t.FailNow()
	}
	head := testBucket.content.Front().Value.(contact)
	if head.ID() == contact1.ID() {
		fmt.Printf("[%s] - failed to remove contact, %+v, from bucket\n", testName, contact1)
	}
}

func TestRemoveCenterContact(t *testing.T) {
	var testName string = "TestRemoveCenterContact"
	var testBucket bucket = NewBucket()
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

	err = testBucket.AddContact(contact1)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	errOne := testBucket.RemoveContact(contact3)
	if errOne != nil {
		fmt.Printf("[%s] - failed to remove contact, error: %s", testName, err.Error())
	}

	for e := testBucket.content.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			fmt.Printf("[%s] - bucket has been corrupted: expected contact, found: %+v\n", testName, e)
			t.FailNow()
		}
		if elem.ID() == contact3.ID() {
			fmt.Printf("[%s] - failed to remove contact, %+v\n", testName, contact3)
			t.FailNow()
		}
	}
}

func TestFindContact(t *testing.T) {
	var testName string = "TestFindContat"
	var testBucket bucket = NewBucket()
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

	err = testBucket.AddContact(contact1)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact5)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	foundContact, err := testBucket.FindContact(contact4.ID())
	if err != nil {
		fmt.Printf("[%s] - unexpected error when searching for contact: searched for %+v, error - %s\n", testName, contact4.ID(), err.Error())
	}
	if foundContact.ID() != contact4.ID() {
		fmt.Printf("[%s] - contact missmatch: expected %+v, found %+v\n", testName, contact4, foundContact)
		t.FailNow()
	}
	if foundContact.IP() != contact4.IP() {
		fmt.Printf("[%s] - contact missmatch: expected %+v, found %+v\n", testName, contact4, foundContact)
		t.FailNow()
	}
	if foundContact.Port() != contact4.Port() {
		fmt.Printf("[%s] - contact missmatch: expected %+v, found %+v\n", testName, contact4, foundContact)
		t.FailNow()
	}

}

func TestFindMissingContact(t *testing.T) {
	var testName string = "TestFindMissingContact"
	var testBucket bucket = NewBucket()
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

	err = testBucket.AddContact(contact1)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact5)
	if err != nil {
		fmt.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	_, err = testBucket.FindContact(contact6.ID())
	if err == nil {
		fmt.Printf("[%s] - uncaught error: expected no matches but error was not raised", testName)
		t.FailNow()
	}

}


