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
				elem := e.Value.(contact)
				fmt.Printf("elem: %+v\n", elem)
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
	var verbose bool = true
	var homeID [5]uint32 = [5]uint32{0, 0, 0, 0, 0}
	var testRT routingTable = NewRoutingTable(homeID)

	if verbose {
	}

	for i := 0; i < 1000; i++ {
		newContact, _ := NewRandomContact()
		testRT.AddContact(newContact)
	}

	targetContact, err := BuildContact("127.0.0.1", 80, [5]uint32{0, 0, 0, 0, 1})
	if err != nil {
		fmt.Printf("[%s] - unexpected error: %+v", testName, err.Error())
	}
	testRT.AddContact(targetContact)

	if verbose {
		fmt.Printf("routing table after insertion:\n")
		for b := 0; b < KEYSPACE; b++ {
			bucket := testRT.router[b].content
			fmt.Printf("bucket: %d\n", b)
			for e := bucket.Front(); e != nil; e = e.Next() {
				elem := e.Value.(contact)
				fmt.Printf("elem: %+v\n", elem)
			}
		}
	}

	foundContacts, err := testRT.FindXClosest(3, [5]uint32{0, 0, 0, 0, 1})
	if err != nil {
		fmt.Printf("[%s] - %s\n", testName, err.Error())
	}
	fmt.Printf("searching for: %+v\nfound:\n", targetContact.ID())
	for e := foundContacts.Front(); e != nil; e = e.Next() {
		elem := e.Value.(contact)
		fmt.Printf("contact: %+v\n", elem)
	}

}