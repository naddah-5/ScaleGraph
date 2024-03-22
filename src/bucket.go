package scalegraph

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type bucket struct {
	lock     sync.RWMutex
	content  *list.List
	capacity int
}

func NewBucket() *bucket {
	var newBucket bucket = bucket{
		content:  list.New(),
		capacity: KBUCKETVOLUME,
	}
	return &newBucket
}

// Adds the given contact to Bucket, returns an error if Bucket is full.
// Updates the contacts position if it already exists in Bucket.
func (bucket *bucket) AddContact(newContact contact) error {
	bucket.lock.Lock()
	defer bucket.lock.Unlock()
	for e := bucket.content.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			return errors.New(fmt.Sprintf("Bucket has been corrupted: expected a contact, found: %+v", e.Value))
		}
		if elem.ID() == newContact.ID() && e != nil {
			bucket.content.MoveToBack(e)
			return nil
		}
	}
	if bucket.content.Len() >= bucket.capacity {
		return errors.New("Bucket is full")
	}
	bucket.content.PushBack(newContact)
	return nil
}

// Removes a contact from the bucket, returns an error if a non-contact element is found.
// Does not return an error if the contact is not present.
func (bucket *bucket) RemoveContact(target contact) error {
	bucket.lock.Lock()
	defer bucket.lock.Unlock()
	for e := bucket.content.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			return errors.New(fmt.Sprintf("Bucket has been corrupted: expected a contact, found: %+v", e.Value))
		}
		if elem.ID() == target.ID() {
			bucket.content.Remove(e)
			return nil
		}
	}
	return nil
}

// Searches Bucket for a contact matching the given ID, returns error if no match is found. Does not update contacts position in the Bucket.
func (bucket *bucket) FindContact(target [5]uint32) (contact, error) {
	bucket.lock.RLock()
	defer bucket.lock.RUnlock()
	var noMatch contact = EmptyContact()
	for e := bucket.content.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			return noMatch, errors.New(fmt.Sprintf("Bucket has been corrupted: expected a contact, found %+v", e.Value))
		}

		if elem.ID() == target {
			return elem, nil
		}
	}
	return noMatch, errors.New("no match")
}

// Returns up to x closest contacts to the given node id, if the Bucket
// contain less than x contacts all contacts are returned along with a "incomplete" error.
// The result is returned in a sorted list, from closest to target to furthest from target.
// Note that this method always performs a deep copy of the Bucket.
func (bucket *bucket) FindXClosestBucket(x int, target [5]uint32) (*list.List, error) {
	var res *list.List = list.New()
	bucket.lock.RLock()
	for e := bucket.content.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			return res, errors.New(fmt.Sprintf("Bucket has been corrupted: expected a contact, found %+v\n", e.Value))
		}
		res.PushFront(elem)
	}
	bucket.lock.RUnlock()
	err := SortByDistance(res, target)
	if err != nil {
		return res, err
	}

	if res.Len() > x {
		for i := res.Len() - x; i > 0; i-- {
			res.Remove(res.Back())
		}
	}

	if res.Len() < x {
		return res, errors.New("incomplete")
	}

	return res, nil
}
