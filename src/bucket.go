package main

import (
	"container/list"
	"errors"
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
		elem := e.Value.(contact)
		if elem.ID() == newContact.ID() && e != nil {
			b.content.MoveToBack(e)
			return nil
		}
	}
	b.content.PushBack(newContact)
	return nil
}

