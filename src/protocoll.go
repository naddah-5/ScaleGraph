package scalegraph

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

func (node *Node) Ping(ip [4]byte) error {
	ping := GenerateRPC(PING, node.contact, ip)
	resp, err := node.network.Send(ping)
	if err != nil {
		return errors.New(fmt.Sprintf("no ping response from IP: %+v", ip))
	}
	node.controlPong(resp)
	return nil
}

func (node *Node) Heartbeat(target contact) {

}

// High level find node RPC.
func (node *Node) FindNode(target [5]uint32) ([]contact, error) {
	closeIP, err := node.routingTable.FindXClosest(REPLICATION, target)
	if err != nil {
		return make([]contact, 0), err
	}
	var wg sync.WaitGroup
	res := list.New()
	var respChan chan []contact = make(chan []contact, REPLICATION)
	for n := closeIP.Front(); n != nil; n = n.Next() {
		closeNode := n.Value.(contact)
		rpc := GenerateRPC(FIND_NODE, node.contact, closeNode.ip)
		rpc.FindNode(closeNode.ID())
		wg.Add(1)
		go func(rpc RPC, wg *sync.WaitGroup, respChan chan []contact, node *Node) {
			resp, err := node.network.Send(rpc)
			if err != nil {
				respChan <- make([]contact, 0)
				wg.Done()
				return
			}
			respChan <- resp.KNodes
			wg.Done()
			go func(rpc RPC) {
				for _, con := range rpc.KNodes {
					node.Ping(con.IP())
				}
			}(resp)
			return
		}(rpc, &wg, respChan, node)
	}
	wg.Wait()
	for i := 0; i < REPLICATION; i++ {
		foundNodes := <-respChan
		for _, val := range foundNodes {
			res.PushBack(val)
		}
	}
	// TODO: Replace list sort with a slice sort
	SortByDistance(res, target)
	var resSlice []contact = make([]contact, 0, REPLICATION)
	i := 0
	for c := res.Front(); c != nil; c = c.Next() {
		if i >= REPLICATION {
			break
		}
		resSlice = append(resSlice, c.Value.(contact))
		i++
	}
	finalRes := node.deepSearch(resSlice, target)
	return finalRes, nil
}

// Helper function to the find node RPC, handles the reccursion.
func (node *Node) deepSearch(prevContactList []contact, target [5]uint32) []contact {
	var wg sync.WaitGroup
	var respChan chan []contact = make(chan []contact, REPLICATION)
	for i := 0; i < len(prevContactList); i++ {
		rpc := GenerateRPC(FIND_NODE, node.contact, prevContactList[i].IP())
		rpc.FindNode(target)
		wg.Add(1)
		go func(rpc RPC, wg *sync.WaitGroup, respChan chan []contact) {
			resp, err := node.Send(rpc)
			if err != nil {
				respChan <- make([]contact, 0)
				wg.Done()
				return
			}
			respChan <- resp.KNodes
			wg.Done()
			go func(rpc RPC) {
				for _, con := range rpc.KNodes {
					node.Ping(con.IP())
				}
			}(resp)

			return
		}(rpc, &wg, respChan)
	}
	res := list.New()
	wg.Wait()
	for i := 0; i < REPLICATION; i++ {
		foundNodes := <-respChan
		for _, val := range foundNodes {
			res.PushBack(val)
		}
	}
	SortByDistance(res, target)
	contactList := make([]contact, 0, REPLICATION)
	i := 0
	for c := res.Front(); c != nil; c = c.Next() {
		if i >= REPLICATION {
			break
		}
		contactList = append(contactList, c.Value.(contact))
		i++
	}
	if CompareContactSlice(prevContactList, contactList) {
		return contactList
	}
	return node.deepSearch(contactList, target)
}
