package main

import (
	"container/list"
	"errors"
	"fmt"
)

type Bucket struct {
	Content  list.List
	Capacity int
}

func NewBucket() *Bucket {
	var newBucket Bucket = Bucket{
		Content:  *list.New(),
		Capacity: KBUCKETVOLUME,
	}

	return &newBucket
}

func (b *Bucket) AddContact(ip string, port int, id [5]uint32) error {
	if b.Content.Len() >= b.Capacity {
		return errors.New("bucket is full")
	}
	newContact, genErr := NewContact(ip, port, id)
	if genErr != nil {
		return genErr
	}
	b.Content.PushBack(newContact)
	return nil
}

func (b *Bucket) sip() {
	cont := b.Content.Front().Value
	fmt.Printf("cont: %v\n", cont)
}
