package kademlia

import (
	"errors"
	"fmt"
	"slices"
	"sync"
)

type Bucket struct {
	homeNode Contact
	content  []Contact
	capacity int
	sync.RWMutex
}

func NewBucket(maxCapacity int, homeNode Contact) *Bucket {
	bucket := Bucket{
		homeNode: homeNode,
		content:  make([]Contact, 0, maxCapacity),
		capacity: maxCapacity,
	}
	return &bucket
}

// Adds given contact if there is empty capacity in bucket or it is closer to the home node than another node.
// Otherwise returns an error.
func (bucket *Bucket) AddContact(contact Contact) error {
	bucket.Lock()
	defer bucket.Unlock()

	for _, v := range bucket.content {
		if v.ID() == contact.ID() {
			return errors.New("cannot add two instances of a contact to a single bucket")
		}
	}
	bucket.content = append(bucket.content, contact)
	SortContactsByDistance(&bucket.content, bucket.homeNode.ID())
	var err error
	if len(bucket.content) > bucket.capacity {
		if bucket.content[len(bucket.content)-1].ID() == contact.ID() {
			err = errors.New("contact not added")
		}
	} else {
		err = nil
	}
	bucket.content = bucket.content[:min(bucket.capacity, len(bucket.content))]
	return err
}

// Searches the bucket for any node with a matching IP address and returns it if found.
// Otherwise returns a nil contact and an error.
func (bucket *Bucket) FindByIP(ip [4]byte) (Contact, error) {
	bucket.Lock()
	defer bucket.Unlock()
	for _, c := range bucket.content {
		if c.IP() == ip {
			return c, nil
		}
	}
	return Contact{}, errors.New("no match")
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

func (bucket *Bucket) Display() string {
	bucket.Lock()
	defer bucket.Unlock()
	res := ""
	for _, val := range bucket.content {
		res += fmt.Sprintf("%s\n", val.Display())
	}
	return res
}
