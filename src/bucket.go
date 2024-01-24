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
	var newBucket bucket = bucket{}
	newBucket.content = list.New()
	newBucket.capacity = KBUCKETVOLUME
	return newBucket
}

func (b *bucket) AddContact(newContact contact) error {
	if b.content.Len() >= b.capacity {
		return errors.New("bucket is full")
	}
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
	b.content.PushBack(newContact)
	return nil
}

func (b *bucket) RemoveContact(target contact) error {
	for e := b.content.Front(); e != nil; e.Next() {
		elem, ok := e.Value.(contact)
		if !ok {
			return errors.New(fmt.Sprintf("bucket has been corrupted: expected a contact, found: %+v", e.Value))
		}
		if elem.ID() == target.ID() {
			fmt.Println("found match")
			b.content.Remove(e)
			fmt.Println("removed match")
			return nil
		}
	}
	return nil
}
