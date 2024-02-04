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
		fmt.Println(err)
		return res, errors.New(fmt.Sprintf("failed to find %d closest contacts, error: %s", x, err.Error()))
	}
	res, err = rt.router[bucketIndex].FindXClosest(x, target)
	if err.Error() == "incomplete" {
		var count int = x - res.Len()
		addRes, err := rt.findSlider(bucketIndex, count, target)
		if err != nil {
			fmt.Println(err)
			return res, err
		}
		for e := addRes.Front(); e != nil; e = e.Next() {
			elem, ok := e.Value.(contact)
			if !ok {
				return res, errors.New(fmt.Sprintf("return from findSlider is corrupted, expected a contact found: %+v", e.Value))
			} else {
				res.PushBack(elem)
			}
		}
	} else if err != nil {
		return res, err
	}

	return res, nil
}

func (rt *routingTable) findSlider(startIndex int, count int, target [5]uint32) (*list.List, error) {
	var res *list.List = list.New()

		for i := max(startIndex, 0); i >= 0; i-- {
		newCont, _ := rt.router[i].FindXClosest(count, target)
		res, _ = MergeByDistance(res, newCont, target)
	}

	for i := min(startIndex, KEYSPACE); i < KEYSPACE; i++ {
		newCont, _ := rt.router[i].FindXClosest(count, target)
		fmt.Println("appending to res:")
		for e := newCont.Front(); e != nil; e = e.Next() {
			elem, ok := e.Value.(contact)
			if !ok {
				fmt.Printf("found illegal element: %+v", e.Value)
			} else {
				fmt.Printf("elem: %+v", elem)
			}
		}
		res, _ = MergeByDistance(res, newCont, target)
	}

	return res, nil
}
