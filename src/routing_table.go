package scalegraph

import (
	"container/list"
	"errors"
	"fmt"
	"log"
)

type routingTable struct {
	homeNode [5]uint32 // depricated
	router   [KEYSPACE]bucket
}

func NewRoutingTable(homeNode [5]uint32) routingTable {
	var newRT routingTable = routingTable{}
	newRT.homeNode = homeNode
	for i := 0; i < len(newRT.router); i++ {
		newRT.router[i] = NewBucket()
	}
	return newRT
}

// Returns the bucket index of the given id.
// Returns an error if the bucket index corresponds to the 'home bucket'.
func (rt *routingTable) BucketIndex(nodeID [5]uint32) (int, error) {
	var bucketIndex int = DistPrefixLength(rt.HomeNodeID(), nodeID)
	if bucketIndex == KEYSPACE {
		return -1, errors.New("home node does not exist in a bucket")
	}
	return bucketIndex, nil
}

func (rt *routingTable) HomeNodeID() [5]uint32 {
	return rt.homeNode
}

func (rt *routingTable) AddContact(target contact) error {
	bucketIndex, err := rt.BucketIndex(target.ID())
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find corresponding bucket index for contact, %+v, error: %s", target, err.Error()))
	}
	err = rt.router[bucketIndex].AddContact(target)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to add contact %+v, error: %s", target, err.Error()))
	}
	return nil
}

func (rt *routingTable) RemoveContact(target contact) error {
	bucketIndex, err := rt.BucketIndex(target.ID())
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find corresponding bucket index for contact, %+v, error %s", target, err.Error()))
	}
	err = rt.router[bucketIndex].RemoveContact(target)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to remove contact %+v, error: %s", target, err.Error()))
	}
	return nil
}

// Returns a list containing the up to 'x' closest known contacts.
// Returns an error if a non-contact element is encountered in a bucket.
func (rt *routingTable) FindXClosest(x int, target [5]uint32) (*list.List, error) {
	var res *list.List = list.New()
	bucketIndex, err := rt.BucketIndex(target)
	if err != nil {
		log.Println(err)
		return res, errors.New(fmt.Sprintf("failed to find %d closest contacts, error: %s", x, err.Error()))
	}
	res, err = rt.router[bucketIndex].FindXClosestBucket(x, target)
	if err.Error() == "incomplete" {
		var count int = x - res.Len()
		addRes, err := rt.findSlider(bucketIndex, count, target)
		if err != nil {
			log.Println(err)
			return res, err
		}
		res, err = MergeByDistance(res, addRes, target)
		if err != nil {
			return res, err
		}
	} else if err != nil {
		return res, err
	}

	for i := res.Len() - x; i > 0; i-- {
		res.Remove(res.Back())
	}

	return res, nil
}

// Returns a list of 'count' contacts closest to the target id, sliding outwards to from the starting bucket index.
// Returns an error if a non-contact element is found.
func (rt *routingTable) findSlider(startIndex int, count int, target [5]uint32) (*list.List, error) {
	var res *list.List = list.New()

	for i := max(startIndex-1, -1); i >= 0; i-- {
		newContent, _ := rt.router[i].FindXClosestBucket(count, target)
		var err error
		res, err = MergeByDistance(res, newContent, target)
		if err != nil {
			return res, err
		}
	}

	for i := min(startIndex+1, KEYSPACE); i < KEYSPACE; i++ {
		newContent, _ := rt.router[i].FindXClosestBucket(count, target)
		var err error
		res, err = MergeByDistance(res, newContent, target)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}
