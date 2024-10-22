package kademlia

import (
	"errors"
)

type RoutingTable struct {
	homeNode [5]uint32
	table    []*Bucket
	keySpace int
}

// Creates and populates a new routing table with buckets.
func NewRoutingTable(homeNode [5]uint32, keySpace int, kBucket int) *RoutingTable {
	router := RoutingTable{
		homeNode: homeNode,
		table:    make([]*Bucket, 0),
		keySpace: keySpace,
	}
	for i := 0; i < keySpace; i++ {
		router.table = append(router.table, NewBucket(kBucket))
	}
	return &router
}

// Returns the bucket index for target ID.
// If index is out of scope, i.e. the home node, returns an error.
func (router *RoutingTable) BucketIndex(target [5]uint32) (int, error) {
	index := DistPrefixLength(target, router.homeNode)
	if index < 0 || index > 159 {
		return 0, errors.New("invalid index")
	}
	return index, nil
}

// Attempts to add the contact to the routing table at the correct bucket.
// Returns an error if adding home node or bucket is full.
func (router *RoutingTable) AddContact(contact Contact) error {
	index, err := router.BucketIndex(contact.ID())
	if err != nil {
		return errors.New("can not add home node to router")
	}
	err = router.table[index].AddContact(contact)
	if err != nil {
		return err
	}
	return nil
}

// Attempts to remove contact from the corresponding bucket.
// If contact is not found does nothing.
func (router *RoutingTable) RemoveContact(contact Contact) {
	index, err := router.BucketIndex(contact.ID())
	if err != nil {
		return
	}
	router.table[index].RemoveContact(contact)
}

func (router *RoutingTable) FindXClosest(x int, target [5]uint32) ([]Contact, error) {
	res := make([]Contact, 0, x)
	index, err := router.BucketIndex(target)
	if err != nil {
		index = router.keySpace - 1
	}
	firstBucket := router.table[index].FindXClosest(x, target)
	res = append(res, firstBucket...)
	for i := 1; len(res) < x; i++ {
		leftBucket := make([]Contact, 0, x)
		rightBucket := make([]Contact, 0, x)
		if (index - i) >= 0 {
			leftBucket = append(leftBucket, router.table[index-i].FindXClosest(x, target)...)
		}
		res = append(res, leftBucket...)
		if (index + i) < router.keySpace {
			rightBucket = append(rightBucket, router.table[index+i].FindXClosest(x, target)...)
		}
		res = append(res, rightBucket...)
		if (index-i) <= 0 && (index+i) >= router.keySpace {
			break
		}
	}
	if len(res) > x {
		res = res[:x]
	}

	return res, nil
}
