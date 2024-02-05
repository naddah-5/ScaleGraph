package main

import (
	"fmt"
	"testing"
)

func TestRoutingTableAddContact(t *testing.T) {
	var testName string = "TestRoutingTableAddContact"
	var verbose bool = false
	var homeID [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var testRT routingTable = NewRoutingTable(homeID)

	for i := 0; i < 1000; i++ {
		newContact, _ := NewRandomContact()
		testRT.AddContact(newContact)
	}

	if verbose {
		fmt.Printf("routing table after insertion:\n")
		for b := 0; b < KEYSPACE; b++ {
			bucket := testRT.router[b].content
			fmt.Printf("bucket: %d\n", b)
			for e := bucket.Front(); e != nil; e = e.Next() {
				elem, ok := e.Value.(contact)
				if !ok {
					fmt.Println("element not a contact")
				} else {
					fmt.Printf("elem: %+v\n", elem)
				}
			}
		}
	}

	if testRT.router[0].content.Len() < testRT.router[0].capacity {
		bucket := testRT.router[0]
		fmt.Printf("[%s] WARNING - highly unlikely to have less than %d contacts in bucket 0\n", testName, testRT.router[0].capacity)
		fmt.Printf("bucket 0 contains:\n")
		var counter int = 0
		for e := bucket.content.Front(); e != nil; e = e.Next() {
			fmt.Printf("elem %d: %+v\n", counter, e.Value)
			counter++
		}
	}
}

func TestRoutingTableFindContact(t *testing.T) {
	var testName string = "TestRoutingTableFindContact"
	var verbose bool = false
	var homeID [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var testRT routingTable = NewRoutingTable(homeID)

	for i := 0; i < 1000; i++ {
		newContact, _ := NewRandomContact()
		testRT.AddContact(newContact)
	}

	targetContact, err := BuildContact("127.0.0.1", 80, [5]uint32{0, 0, 0, 0, 1})
	if err != nil {
		fmt.Printf("[%s] - unexpected error: %+v", testName, err.Error())
	}
	testRT.AddContact(targetContact)

	foundContacts, err := testRT.FindXClosest(3, [5]uint32{0, 0, 0, 0, 1})
	if err != nil {
		fmt.Printf("[%s] - %s\n", testName, err.Error())
	}

	if verbose {
		fmt.Printf("searching for: %+v\nfound:\n", targetContact.ID())
		for e := foundContacts.Front(); e != nil; e = e.Next() {
			elem, ok := e.Value.(contact)
			var relDist int = RelativeDistance(elem.ID(), targetContact.ID())
			if !ok {
				fmt.Printf("element is not a contact: %+v\n", e.Value)
				break
			}
			fmt.Printf("contact: ip : %-15v port : %-5v ID : %-10v relative distance : %-5d\n", elem.IP(), elem.Port(), elem.ID(), relDist)
		}
	}
	closestFound, ok := foundContacts.Front().Value.(contact)
	if !ok {
		fmt.Printf("[%s] - list has been corrupted: expected contact, found %+v", testName, foundContacts.Front().Value)
		t.Fail()
	}
	if RelativeDistance(closestFound.ID(), targetContact.ID()) != 0 {
		fmt.Printf("[%s] - failed to find correct contact, expected to find %+v with relative distance of 0", testName, targetContact)
		t.Fail()
	}
}

func TestRoutingTableRemoveContact(t *testing.T) {
	var testName string = "TestRoutingTableRemoveContact"
	var verbose bool = false
	var homeID [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var testRT routingTable = NewRoutingTable(homeID)

	for i := 0; i < 1000; i++ {
		newContact, _ := NewRandomContact()
		testRT.AddContact(newContact)
	}

	targetContact, err := BuildContact("127.0.0.1", 80, [5]uint32{0, 0, 0, 0, 1})
	if err != nil {
		fmt.Printf("[%s] - unexpected error: %+v", testName, err.Error())
	}
	testRT.AddContact(targetContact)

	foundContacts, err := testRT.FindXClosest(3, targetContact.ID())
	if err != nil {
		fmt.Printf("[%s] - %s\n", testName, err.Error())
	}


	checkAdd, ok := foundContacts.Front().Value.(contact)
	if !ok {
		fmt.Printf("[%s] - contact was not added properly, found: %+v", testName, foundContacts.Front().Value)
		t.Fail()
	}
	if checkAdd.ID() != targetContact.ID() {
		fmt.Printf("[%s] - contact was not properly added, expected %+v, found %+v", testName, targetContact, checkAdd)
	}

	if verbose {
		fmt.Printf("target, %+v, has been added to the routing table\n", targetContact)
		fmt.Printf("removing %+v from test routing table\n", targetContact)
	}

	err = testRT.RemoveContact(targetContact)
	if err != nil {
		fmt.Printf("[%s] - failed to remove contact %+v: error - %s", testName, targetContact, err.Error())
		t.Fail()
	}

	updatedContacts, err := testRT.FindXClosest(3, targetContact.ID())
	if err != nil {
		fmt.Printf("[%s] - failed to search routing table after removal", testName)
		t.Fail()
	}
	

	checkRemoved, ok := updatedContacts.Front().Value.(contact)
	if !ok {
		fmt.Printf("[%s] - test list was corrupted: expected to find a contact, found %+v", testName, updatedContacts.Front().Value)
		t.Fail()
	}
	if checkRemoved.ID() == targetContact.ID() {
		fmt.Printf("[%s] - failed to remove target contact, %+v", testName, targetContact)
		t.Fail()
	}

	if verbose {
		fmt.Printf("search for target after removal returned, %+v\n", checkRemoved)
	}
}

