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

func (b *bucket) AddContact(ip string, port int, id [5]uint32) error {
	if b.content.Len() >= b.capacity {
		return errors.New("bucket is full")
	}
	for e := b.content.Front(); e != nil; e = e.Next() {
		elem := e.Value.(contact)
		if elem.ID() == id {
			return errors.New("node already in list")
		}
	}
	newContact, err := NewContact(ip, port, id)
	if err != nil {
		return err
	}
	b.content.PushBack(newContact)
	return nil
}

