package main

import (
	"container/list"
	"errors"
	"fmt"
)

type bucket struct {
	content  *list.List
	capacity int
}

func NewBucket() bucket {
	var newBucket bucket = bucket{
		content:  list.New(),
		capacity: KBUCKETVOLUME,
	}
	return newBucket
}

// Adds the given contact to bucket, returns an error if bucket is full.
// Updates the contacts position if it already exists in bucket.
func (b *bucket) AddContact(newContact contact) error {
	for e := b.content.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			return errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact, found: %+v", e.Value))
		}
		if elem.ID() == newContact.ID() && e != nil {
			b.content.MoveToBack(e)
			return nil
		}
	}
	if b.content.Len() >= b.capacity {
		return errors.New("bucket is full")
	}
	b.content.PushBack(newContact)
	return nil
}

func (b *bucket) RemoveContact(target contact) error {
	for e := b.content.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			return errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact, found: %+v", e.Value))
		}
		if elem.ID() == target.ID() {
			b.content.Remove(e)
			return nil
		}
	}
	return nil
}

// Searches bucket for a contact matching the given ID, returns error if no match is found. Does not update contacts position in the bucket.
func (b *bucket) FindContact(target [5]uint32) (contact, error) {
	var noMatch contact = EmptyContact()
	for e := b.content.Front(); e != nil; e = e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			return noMatch, errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact, found %+v", e.Value))
		}

		if elem.ID() == target {
			return elem, nil
		}
	}
	return noMatch, errors.New("no match")
}
