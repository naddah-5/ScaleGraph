package kademlia

import (
	"fmt"
	"log"
)

// Protocol handles the logic for sending RPC's

// Logic for sending a ping RPC.
func (node *Node) Ping(address [4]byte) {
	rpc := GenerateRPC(node.Contact)
	rpc.Ping(address)
	res, err := node.Send(rpc)
	if err != nil {
		log.Printf("[ERROR] - %s", err.Error())
	}
	node.AddContact(res.sender)
}

func (node *Node) FindNode(target [5]uint32) []Contact {
	initNodes, _ := node.FindXClosest(CONCURRENCY, target)
	return node.findNodeRec(initNodes, target)
}

func (node *Node) findNodeRec(prevContactList []Contact, target [5]uint32) []Contact {
	contactList := make([]Contact, 0)
	respChan := make(chan []Contact, 64)

	// Launch parallel queries to initial nodes.
	for _, tar := range prevContactList {
		rpc := GenerateRPC(node.Contact)
		rpc.FindNode(tar.IP(), target)
		go node.nodeQuery(rpc, respChan)
	}

	// Extract results from parallel query.
	for i := range prevContactList {
		if node.debug {
			log.Printf("[DEBUG] - here %d", i)
		}
		resp, ok := <-respChan
		if ok {
			contactList = append(contactList, resp...)
		}
	}

	SortContactsByDistance(&contactList, target)
	if node.debug {
		pRes := fmt.Sprintf("found nodes:\n")
		for _, n := range contactList {
			pRes += fmt.Sprintf("%s\n", n.Display())
		}
		pRes += fmt.Sprintf("input nodes [DEBUG]:\n")
		for _, n := range prevContactList {
			pRes += fmt.Sprintf("%s\n", n.Display())
		}
		log.Printf(pRes)
	}

	if node.debug {
		log.Println("[DEBUG] - here")
	}
	// makes sure there is something in the lists to compare, prevents nil pointer.
	safetyGuard := len(prevContactList) > 0 && len(contactList) > 0
	if !safetyGuard {
		return prevContactList
	}

	precision := SliceContains(prevContactList[0].ID(), &contactList)
	if precision {
		return contactList
	} else if len(contactList) == 0 {
		return prevContactList
	}
	if node.debug {
		log.Printf("[DEBUG] - entering deeper find node reccursion\n")
	}
	return node.findNodeRec(contactList, target)
}

// Sends the given RPC and returns the reponse to the provided channel.
// If the RPC times out or returns an error, returns an empty contact.
// NOTE that you must assert the type of the result from respChan.
func (node *Node) nodeQuery(rpc RPC, respChan chan []Contact) {
	resp, err := node.Send(rpc)
	if err != nil {
		log.Printf("[ERROR] - %s\nin node %v with rpc:\n%s\n", err.Error(), node.ID(), rpc.Display())
		respChan <- resp.foundNodes
		return
	}
	for _, n := range resp.foundNodes {
		go node.AddContact(n)
	}
	respChan <- resp.foundNodes
	return

}
