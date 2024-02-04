package main

import (
	"container/list"
	"errors"
	"fmt"
)

type routingTable struct {
	homeNode [5]uint32
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

func (rt *routingTable) FindXClosest(x int, target [5]uint32) (*list.List, error) {
	var res *list.List = list.New()
	bucketIndex, err := rt.BucketIndex(target)
	if err != nil {
		return res, errors.New(fmt.Sprintf("failed to find %d closest contacts, error: %s", x, err.Error()))
	}
	res, err = rt.router[bucketIndex].FindXClosest(x, target)
	if err.Error() == "incomplete" {
		// call split find function in a loop
		var count int = x - res.Len()
		addRes, err := rt.findSlider(bucketIndex, count, target)
		if err != nil {
			return res, err
		}
		for e := addRes.Front(); e != nil; e.Next() {
			res.PushBack(e)
		}
	} else if err != nil {
		return res, err
	}

	return res, nil
}

func (rt *routingTable) findSlider(startIndex int, count int, target [5]uint32) (*list.List, error) {
	var res *list.List = list.New()
	var leftIndex int = startIndex
	var rightIndex int = startIndex

	for count > 0 {
		leftIndex = max(leftIndex-1, -1)
		rightIndex = min(rightIndex+1, KEYSPACE)
		var leftBucket *list.List = list.New()
		var rightBucket *list.List = list.New()

		if leftIndex > -1 {
			leftBucket = rt.takeContent(leftIndex)
			SortByDistance(leftBucket, target)
		}

		if rightIndex < KEYSPACE {
			rightBucket = rt.takeContent(rightIndex)
			SortByDistance(rightBucket, target)
		}

		addRes, err := MergeByDistance(leftBucket, rightBucket, target)
		if err != nil {
			fmt.Println(err)
		}
		res, err = MergeByDistance(res, addRes, target)
		if err != nil {
			fmt.Println(err)
		}
		count = count - addRes.Len()
		fmt.Println("left index: ", leftIndex, " right index: ", rightIndex)
		if leftIndex <= 0 && rightIndex >= KEYSPACE {
			break
		}
	}

	return res, nil
}

func (rt *routingTable) takeContent(bucket int) *list.List {
	var res *list.List = list.New()
	for e := rt.router[bucket].content.Front(); e != nil; e = e.Next() {
		res.PushBack(e.Value)
	}
	return res
}
