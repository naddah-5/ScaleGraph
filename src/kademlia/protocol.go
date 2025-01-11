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
	found := node.findNodeLoop(initNodes, target)
	return found
}

func (node *Node) findNodeLoop(prevContactList []Contact, target [5]uint32) []Contact {
	contactList := make([]Contact, 0, REPLICATION)
	respChan := make(chan []Contact, 64)

	for {
		// Launch parallel queries to initial nodes.
		for _, n := range prevContactList {
			rpc := GenerateRPC(node.Contact)
			rpc.FindNode(n.IP(), target)
			go node.nodeQuery(rpc, respChan)
		}

		// Extract results from parallel query.
		for range prevContactList {
			resp, ok := <-respChan
			if ok {
				contactList = append(contactList, resp...)
			}
		}

		// Process the found contacts
		SortContactsByDistance(&contactList, target)
		RemoveDuplicateContacts(&contactList)
		if len(contactList) > CONCURRENCY {
			contactList = contactList[:CONCURRENCY]
		}

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

		closer := CloserNode(contactList[0].ID(), prevContactList[0].ID(), target)
		if !closer {
			return contactList
		} else if len(contactList) == 0 {
			return prevContactList
		}
		prevContactList = nil
		prevContactList = contactList
		contactList = make([]Contact, 0, REPLICATION)
	}
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
