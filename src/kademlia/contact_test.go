package kademlia

import (
	"log"
	"testing"
)

// Manual inspection for the contact display.
func TestDisplayContact(t *testing.T) {
	testName := "TestDisplayContact"
	verbose := false
	if verbose {
		log.Printf("[%s]\n", testName)
		for range 100 {
			con := NewRandomContact()
			log.Printf("%s\n", con.Display())
		}
	}
}
