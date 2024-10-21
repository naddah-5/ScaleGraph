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

func TestFindXClosest(t *testing.T) {
	testName := "TestFindXClosest"
	verbose := false
	testBucketSize := 10
	bucket := NewBucket(testBucketSize)
	nodeA := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0})
	nodeB := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 0, 1})
	nodeC := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 5, 0})
	nodeD := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 5, 0, 0})
	nodeE := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 1, 0, 0})
	nodeF := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 1, 0, 0, 0})
	nodeG := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{1, 0, 0, 0, 0})
	nodeH := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{10, 92, 23, 233, 0})
	nodeI := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 99, 32, 0, 0})
	nodeJ := NewContact([4]byte{0, 0, 0, 0}, [5]uint32{0, 0, 0, 10, 1})

	bucket.AddContact(nodeA)
	bucket.AddContact(nodeB)
	bucket.AddContact(nodeC)
	bucket.AddContact(nodeD)
	bucket.AddContact(nodeE)
	bucket.AddContact(nodeF)
	bucket.AddContact(nodeG)
	bucket.AddContact(nodeH)
	bucket.AddContact(nodeI)
	bucket.AddContact(nodeJ)

	target := [5]uint32{10, 92, 0, 0, 0}
	res := bucket.FindXClosest(3, target)
	if verbose {
		log.Printf("[%s]\n", testName)
		log.Println("input slice")
		for _, v := range bucket.content {
			log.Printf("contact: %v\n", v)
		}
		log.Printf("target ID: %v", target)
		log.Println("returned contacts")
		for _, v := range res {
			log.Printf("contact: %2v\tdistance: %2v", v, RelativeDistance(v.ID(), target))
		}
	}
	if res[0].ID() != nodeH.ID() {
		log.Printf("[%s] - incorrect closest contact returned")
		t.Fail()
	}
}
