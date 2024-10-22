package kademlia

import (
	"errors"
	"slices"
	"sync"
)

type Bucket struct {
	content  []Contact
	capacity int
	sync.RWMutex
}

func NewBucket(maxCapacity int) *Bucket {
	bucket := Bucket{
		content:  make([]Contact, 0, maxCapacity),
		capacity: maxCapacity,
	}
	return &bucket
}

// Adds given contact if there is empty capacity in bucket.
// Otherwise returns an error.
func (bucket *Bucket) AddContact(contact Contact) error {
	bucket.Lock()
	defer bucket.Unlock()

	if len(bucket.content) == bucket.capacity {
		return errors.New("full bucket")
	}
	for _, v := range bucket.content {
		if v.ID() == contact.ID() {
			return nil
		}
	}
	bucket.content = append(bucket.content, contact)
	return nil
}

// Removes contact from bucket if it is present.
func (bucket *Bucket) RemoveContact(contact Contact) {
	bucket.Lock()
	defer bucket.Unlock()

	for i, v := range bucket.content {
		if v.ID() == contact.ID() {
			bucket.content = slices.Delete(bucket.content, i, i+1)
		}
	}
}

// Returns up to x contacts from the bucket.
func (bucket *Bucket) FindXClosest(x int, target [5]uint32) []Contact {
	bucket.Lock()
	defer bucket.Unlock()
	res := make([]Contact, 0, x)
	res = append(res, bucket.content...)
	SortContactsByDistance(&res, target)
	res = res[:min(x, len(bucket.content))]
	return res
}


// Returns a contact with matching ID to target if present.
// Otherwise returns an error.
func (bucket *Bucket) FindContact(target [5]uint32) (Contact, error) {
	bucket.Lock()
	defer bucket.Unlock()
	for _, v := range bucket.content {
		if v.ID() == target {
			return v, nil
		}
	}
	return Contact{}, errors.New("contact not found")
}
