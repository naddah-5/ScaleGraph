package scalegraph

import (
	"log"
	"testing"
)

func TestFillNewBucket(t *testing.T) {
	var testName string = "TestFillBucket"
	var testBucket bucket = NewBucket()
	contact1 := BuildContact([4]byte{127, 0, 0, 1}, 80, [5]uint32{1, 2, 3, 4, 5})
	contact2 := BuildContact([4]byte{127, 0, 0, 2}, 80, [5]uint32{6, 7, 8, 9, 10})
	contact3 := BuildContact([4]byte{127, 0, 0, 3}, 80, [5]uint32{11, 12, 13, 14, 15})
	contact4 := BuildContact([4]byte{127, 0, 0, 4}, 80, [5]uint32{16, 17, 18, 19, 20})
	contact5 := BuildContact([4]byte{127, 0, 0, 5}, 80, [5]uint32{21, 22, 23, 24, 25})
	overflowContact := BuildContact([4]byte{127, 0, 0, 6}, 80, [5]uint32{26, 27, 28, 29, 30})

	err := testBucket.AddContact(contact1)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact5)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	err = testBucket.AddContact(overflowContact)
	if err == nil {
		log.Printf("[%s] - expected full Bucket error", testName)
		t.FailNow()
	}
}

func TestDoubbleAddBucket(t *testing.T) {
	var testName string = "TestDoubbleAddBucket"
	var testBucket bucket = NewBucket()
	testContact := BuildContact([4]byte{127, 0, 0, 1}, 80, [5]uint32{1, 2, 3, 4, 5})
	bufferContact := BuildContact([4]byte{127, 0, 0, 2}, 80, [5]uint32{6, 7, 8, 9, 10})

	err := testBucket.AddContact(testContact)
	if err != nil {
		log.Printf("[%s] - unexpected error when adding contact: %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(bufferContact)
	if err != nil {
	}
	err = testBucket.AddContact(testContact)
	if err != nil {
		log.Printf("[%s] - unexpected error when adding existing contact: %s", testName, err.Error())
		t.FailNow()
	}

	var lastSeen contact = testBucket.content.Back().Value.(contact)
	if lastSeen.ID() != testContact.ID() {
		log.Printf("[%s] - expected last seen contact to have id: %+v, found node ID: %+v\n", testName, testContact.ID(), lastSeen.ID())
	}
}

func TestRemoveHeadContact(t *testing.T) {
	var testName string = "TestRemoveHeadContact"
	var testBucket bucket = NewBucket()
	contact1 := BuildContact([4]byte{127, 0, 0, 1}, 80, [5]uint32{1, 2, 3, 4, 5})
	contact2 := BuildContact([4]byte{127, 0, 0, 2}, 80, [5]uint32{6, 7, 8, 9, 10})
	contact3 := BuildContact([4]byte{127, 0, 0, 3}, 80, [5]uint32{11, 12, 13, 14, 15})
	contact4 := BuildContact([4]byte{127, 0, 0, 4}, 80, [5]uint32{16, 17, 18, 19, 20})

	err := testBucket.AddContact(contact1)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	err = testBucket.RemoveContact(contact1)
	if err != nil {
		log.Printf("[%s] - failed to remove contact, error: %s", testName, err.Error())
		t.FailNow()
	}
	head := testBucket.content.Front().Value.(contact)
	if head.ID() == contact1.ID() {
		log.Printf("[%s] - failed to remove contact, %+v, from Bucket\n", testName, contact1)
	}
}

func TestRemoveCenterContact(t *testing.T) {
	var testName string = "TestRemoveCenterContact"
	var testBucket bucket = NewBucket()
	contact1 := BuildContact([4]byte{127, 0, 0, 1}, 80, [5]uint32{1, 2, 3, 4, 5})
	contact2 := BuildContact([4]byte{127, 0, 0, 2}, 80, [5]uint32{6, 7, 8, 9, 10})
	contact3 := BuildContact([4]byte{127, 0, 0, 3}, 80, [5]uint32{11, 12, 13, 14, 15})
	contact4 := BuildContact([4]byte{127, 0, 0, 4}, 80, [5]uint32{16, 17, 18, 19, 20})

	err := testBucket.AddContact(contact1)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	errOne := testBucket.RemoveContact(contact3)
	if errOne != nil {
		log.Printf("[%s] - failed to remove contact, error: %s", testName, err.Error())
	}

	for e := testBucket.content.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			log.Printf("[%s] - Bucket has been corrupted: expected contact, found: %+v\n", testName, e)
			t.FailNow()
		}
		if elem.ID() == contact3.ID() {
			log.Printf("[%s] - failed to remove contact, %+v\n", testName, contact3)
			t.FailNow()
		}
	}
}

func TestFindContact(t *testing.T) {
	var testName string = "TestFindContat"
	var testBucket bucket = NewBucket()
	contact1 := BuildContact([4]byte{127, 0, 0, 1}, 80, [5]uint32{1, 2, 3, 4, 5})
	contact2 := BuildContact([4]byte{127, 0, 0, 2}, 80, [5]uint32{6, 7, 8, 9, 10})
	contact3 := BuildContact([4]byte{127, 0, 0, 3}, 80, [5]uint32{11, 12, 13, 14, 15})
	contact4 := BuildContact([4]byte{127, 0, 0, 4}, 80, [5]uint32{16, 17, 18, 19, 20})
	contact5 := BuildContact([4]byte{127, 0, 0, 5}, 80, [5]uint32{21, 22, 23, 24, 25})

	err := testBucket.AddContact(contact1)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact5)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	foundContact, err := testBucket.FindContact(contact4.ID())
	if err != nil {
		log.Printf("[%s] - unexpected error when searching for contact: searched for %+v, error - %s\n", testName, contact4.ID(), err.Error())
	}
	if foundContact.ID() != contact4.ID() {
		log.Printf("[%s] - contact missmatch: expected %+v, found %+v\n", testName, contact4, foundContact)
		t.FailNow()
	}
	if foundContact.IP() != contact4.IP() {
		log.Printf("[%s] - contact missmatch: expected %+v, found %+v\n", testName, contact4, foundContact)
		t.FailNow()
	}
	if foundContact.Port() != contact4.Port() {
		log.Printf("[%s] - contact missmatch: expected %+v, found %+v\n", testName, contact4, foundContact)
		t.FailNow()
	}

}

func TestFindMissingContact(t *testing.T) {
	var testName string = "TestFindMissingContact"
	var testBucket bucket = NewBucket()
	contact1 := BuildContact([4]byte{127, 0, 0, 1}, 80, [5]uint32{1, 2, 3, 4, 5})
	contact2 := BuildContact([4]byte{127, 0, 0, 2}, 80, [5]uint32{6, 7, 8, 9, 10})
	contact3 := BuildContact([4]byte{127, 0, 0, 3}, 80, [5]uint32{11, 12, 13, 14, 15})
	contact4 := BuildContact([4]byte{127, 0, 0, 4}, 80, [5]uint32{16, 17, 18, 19, 20})
	contact5 := BuildContact([4]byte{127, 0, 0, 5}, 80, [5]uint32{21, 22, 23, 24, 25})
	contact6 := BuildContact([4]byte{127, 0, 0, 6}, 80, [5]uint32{26, 27, 28, 29, 30})

	err := testBucket.AddContact(contact1)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact5)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	_, err = testBucket.FindContact(contact6.ID())
	if err == nil {
		log.Printf("[%s] - uncaught error: expected no matches but error was not raised", testName)
		t.FailNow()
	}

}

func TestFindXClosest(t *testing.T) {
	var testName string = "TestFindXClosest"
	const inspectTest bool = false
	var target [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var testBucket bucket = NewBucket()

	contact1 := BuildContact([4]byte{127, 0, 0, 1}, 80, [5]uint32{1, 2, 3, 4, 5})
	contact2 := BuildContact([4]byte{127, 0, 0, 2}, 80, [5]uint32{6, 7, 8, 9, 10})
	contact3 := BuildContact([4]byte{127, 0, 0, 3}, 80, [5]uint32{11, 12, 13, 14, 15})
	contact4 := BuildContact([4]byte{127, 0, 0, 4}, 80, [5]uint32{16, 17, 18, 19, 20})
	contact5 := BuildContact([4]byte{127, 0, 0, 5}, 80, [5]uint32{21, 22, 23, 24, 25})

	err := testBucket.AddContact(contact1)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact2)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact3)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact4)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}
	err = testBucket.AddContact(contact5)
	if err != nil {
		log.Printf("[%s] - %s", testName, err.Error())
		t.FailNow()
	}

	if inspectTest {
		log.Println("Bucket before:")
		for e := testBucket.content.Front(); e != nil; e = e.Next() {
			elem, ok := e.Value.(contact)
			if !ok {
				log.Printf("corrupted Bucket: %+v\n", e.Value)
			}
			relDist := RelativeDistance(elem.ID(), target)
			log.Printf("elem: %+v, relDist: %d\n", elem, relDist)
		}
	}
	res, err := testBucket.FindXClosest(2, target)
	if err != nil {
		log.Println(err.Error())
	}

	if inspectTest {
		log.Println("selected:")
		for e := res.Front(); e != nil; e = e.Next() {
			elem := e.Value.(contact)
			relDist := RelativeDistance(elem.ID(), target)
			log.Printf("elem: %+v, relDist: %d\n", elem, relDist)
		}
		log.Println("Bucket after select")
		for e := testBucket.content.Front(); e != nil; e = e.Next() {
			elem := e.Value.(contact)
			log.Printf("elem: %+v\n", elem)
		}
	}

	resOne := res.Front().Value.(contact)
	resTwo := res.Back().Value.(contact)
	if resOne.ID() != contact1.ID() {
		log.Printf("[%s] - expected to find %+v, found %+v\n", testName, contact1, resOne)
		t.Fail()
	}
	if resTwo.ID() != contact4.ID() {
		log.Printf("[%s] - expected to find %+v, found %+v\n", testName, contact4, resTwo)
		t.Fail()
	}

}
