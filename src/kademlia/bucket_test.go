package kademlia

import (
	"log"
	"testing"
)

func TestAddContact(t *testing.T) {
	testName := "TestAddContact"
	verbose := false
	testBucketSize := 10
	bucket := NewBucket(testBucketSize)
	for i := 0; i < testBucketSize; i++ {
		bucket.AddContact(NewRandomContact())
	}

	if verbose {
		log.Printf("[%s] - full bucket\n", testName)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
	}

	con := NewRandomContact()
	err := bucket.AddContact(con)
	if err == nil {
		log.Printf("[%s] - should not be able to add contact %v\n", testName, con)
		t.Fail()
	}

	if verbose {
		log.Printf("[%s] - bucket after add contact call\n", testName)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
	}
}

func TestRemoveContact(t *testing.T) {
	testName := "TestRemoveContact"
	verbose := false
	testBucketSize := 10
	bucket := NewBucket(testBucketSize)
	for i := 0; i < testBucketSize-1; i++ {
		bucket.AddContact(NewRandomContact())
	}
	con := NewRandomContact()
	bucket.AddContact(con)

	if verbose {
		log.Printf("[%s] - full bucket\n", testName)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
		log.Println()
	}

	bucket.RemoveContact(con)

	if verbose {
		log.Printf("[%s] - bucket after remove contact %v call\n", testName, con)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
	}
}

func TestRemoveContact1(t *testing.T) {
	testName := "TestRemoveContact1"
	verbose := false
	testBucketSize := 10
	bucket := NewBucket(testBucketSize)
	con := NewRandomContact()
	bucket.AddContact(con)
	for i := 0; i < testBucketSize-1; i++ {
		bucket.AddContact(NewRandomContact())
	}

	if verbose {
		log.Printf("[%s] - full bucket\n", testName)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
		log.Println()
	}

	bucket.RemoveContact(con)

	if verbose {
		log.Printf("[%s] - bucket after remove contact %v call\n", testName, con)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
	}
}

func TestRemoveContact2(t *testing.T) {
	testName := "TestRemoveContact2"
	verbose := false
	testBucketSize := 10
	bucket := NewBucket(testBucketSize)
	for i := 0; i < testBucketSize/2; i++ {
		bucket.AddContact(NewRandomContact())
	}
	con := NewRandomContact()
	bucket.AddContact(con)
	for i := 0; i < (testBucketSize/2)-1; i++ {
		bucket.AddContact(NewRandomContact())
	}

	if verbose {
		log.Printf("[%s] - full bucket\n", testName)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
		log.Println()
	}

	bucket.RemoveContact(con)

	if verbose {
		log.Printf("[%s] - bucket after remove contact %v call\n", testName, con)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
	}
}

func TestRemoveContact3(t *testing.T) {
	testName := "TestRemoveContact3"
	verbose := false
	testBucketSize := 10
	bucket := NewBucket(testBucketSize)
	for i := 0; i < testBucketSize; i++ {
		bucket.AddContact(NewRandomContact())
	}
	con := NewRandomContact()

	if verbose {
		log.Printf("[%s] - full bucket\n", testName)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
		log.Println()
	}

	bucket.RemoveContact(con)

	if verbose {
		log.Printf("[%s] - bucket after remove contact %v call\n", testName, con)
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
	}
}
