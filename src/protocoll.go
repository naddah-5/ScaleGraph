package scalegraph

import (
	"container/list"
	"errors"
	"fmt"
	"log"
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
	closeIP, _ := node.routingTable.FindXClosest(REPLICATION, target)
	res := list.New()
	respChan := make(chan []contact, 100*REPLICATION)
	for n := closeIP.Front(); n != nil; n = n.Next() {
		closeNode := n.Value.(contact)
		rpc := GenerateRPC(FIND_NODE, node.contact, closeNode.ip)
		rpc.FindNode(closeNode.ID())
		go node.alphaFindNode(rpc, respChan)
	}
	for i := 0; i < min(REPLICATION, closeIP.Len()); i++ {
		foundNodes := <-respChan
		if DEBUG {
			log.Printf("[info] - received find node alpha response: %+v", foundNodes)
		}
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

func (node *Node) alphaFindNode(rpc RPC, respChan chan []contact) {
	resp, err := node.network.Send(rpc)
	if err != nil {
		log.Println("[error] - find node alpha: no response")
		respChan <- make([]contact, 0)
		return
	}
	if DEBUG {
		log.Println("[info] - sending find node alpha response")
	}
	respChan <- resp.KNodes
	go func(rpc RPC) {
		for _, con := range rpc.KNodes {
			go node.Ping(con.IP())
		}
	}(resp)
	return

}

// Helper function to the find node RPC, handles the reccursion.
func (node *Node) deepSearch(prevContactList []contact, target [5]uint32) []contact {
	var respChan chan []contact = make(chan []contact, REPLICATION)
	for i := 0; i < len(prevContactList); i++ {
		rpc := GenerateRPC(FIND_NODE, node.contact, prevContactList[i].IP())
		rpc.FindNode(target)
		go node.alphaFindNode(rpc, respChan)
	}
	res := list.New()
	for i := 0; i < min(REPLICATION, len(prevContactList)); i++ {
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

func (node *Node) StoreWallet(wallet *wallet) error {
	return nil
}
