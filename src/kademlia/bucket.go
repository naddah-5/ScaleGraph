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
