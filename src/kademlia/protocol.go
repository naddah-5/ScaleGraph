package kademlia

import (
	"errors"
	"fmt"
	"log"
)

// Protocol handles the logic for sending RPC's

// Critical in order to reduce the risk of dead networks on start up.
// A dead network occurs when one or more nodes know of the network but is not known of by the network.
func (node *Node) Enter() {
	rpc := GenerateRPC(node.IP(), node.Contact)
	rpc.Enter()
	res, err := node.Send(rpc)
	if err != nil {
		log.Printf("%v - {ENTER} did not receive entry point", node.ID())
		return
	}
	if len(res.foundNodes) == 0 {
		log.Printf("[PANIC] - nil error prevented")
	}
	if res.foundNodes[0].IP() == [4]byte{0, 0, 0, 0} {
		log.Printf("%v - {ENTER} received illegal entry point", node.ID())
	}
	entryNode := res.foundNodes[0]
	branchNode := res.foundNodes[1]
	node.Ping(entryNode.IP())
	node.Ping(node.masterNode.IP())

	node.FindNode(node.Contact.ID())
	node.FindNode(branchNode.ID())
	node.FindNode(node.masterNode.ID())
}

// Logic for sending a ping RPC.
func (node *Node) Ping(address [4]byte) bool {
	rpc := GenerateRPC(address, node.Contact)
	rpc.Ping()
	res, err := node.Send(rpc)
	if err != nil {
		if node.debug {
			log.Printf("%v - [ERROR] RPC %v %s", node.ID(), rpc.id, err.Error())
		}
		return false
	} else {
		node.AddContact(res.sender)
		return true
	}
}

func (node *Node) FindNode(target [5]uint32) []Contact {
	initNodes, _ := node.FindXClosest(REPLICATION, target)
	found := node.findNodeLoop(initNodes, target)
	return found
}

func (node *Node) findNodeLoop(prevContactList []Contact, target [5]uint32) []Contact {
	contactList := make([]Contact, 0, REPLICATION)
	respChan := make(chan []Contact, 64)

	for {
		// Launch parallel queries to initial nodes.
		for _, n := range prevContactList {
			rpc := GenerateRPC(n.IP(), node.Contact)
			rpc.FindNode(target)
			go node.findNodeQuery(rpc, respChan)
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
			contactList = contactList[:REPLICATION]
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

		// If none of the new contacts are closer to the target than in the previous itteration
		// return the previous itterations contact list.
		if len(contactList) > 0 && len(prevContactList) > 0 {
			sameDist := 0
			for i := range contactList {
				if i == len(prevContactList) {
					break
				}
				closer := CloserNode(contactList[i].ID(), prevContactList[i].ID(), target)
				if !closer {
					sameDist++
				}
			}
			if len(contactList) == sameDist {
				return prevContactList
			}
		} else if len(contactList) == 0 {
			// If there are no new contacts in the new contact list, return the previous contact list.
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
func (node *Node) findNodeQuery(rpc RPC, respChan chan []Contact) {
	resp, err := node.Send(rpc)
	if err != nil {
		if node.debug {
			log.Printf("[ERROR] - %s\nin node %v with rpc:\n%s\n", err.Error(), node.ID(), rpc.Display())
		}
		respChan <- resp.foundNodes
		return
	}
	for _, n := range resp.foundNodes {
		go node.Ping(n.IP())
	}
	respChan <- resp.foundNodes
	return

}

func (node *Node) InsertAccount(accID [5]uint32) {
	inertionPoint := node.FindNode(accID)
	if len(inertionPoint) == 0 {
		log.Printf("failed to insert account %10v", accID)
	} else {
		rpc := GenerateRPC(inertionPoint[0].IP(), node.Contact)
		rpc.InsertAccount(accID)
		node.Send(rpc)
	}
}

// Searches for the closest nodes to the account and sends a store account RPC to them.
func (node *Node) StoreAccount(accID [5]uint32) {
	validators := node.FindNode(accID)
	// validators = append(validators, node.Contact)
	// SortContactsByDistance(&validators, accID)
	for _, n := range validators {
		rpc := GenerateRPC(n.IP(), node.Contact)
		rpc.StoreAccount(accID)
		node.Send(rpc)
	}
}

func (node *Node) FindAccount(accID [5]uint32) ([]Contact, error) {
	closeNodes := node.FindNode(accID)
	respChan := make(chan bool, REPLICATION)
	for _, n := range closeNodes {
		node.findAccountQuery(n.IP(), respChan, accID)
	}
	foundAccountNodes := 0
	for range len(closeNodes) {
		_, ok := <-respChan
		if !ok {
			continue
		} else {
			foundAccountNodes++
		}
	}
	if foundAccountNodes == REPLICATION {
		return closeNodes, nil
	} else {
		return closeNodes, errors.New(fmt.Sprintf("Failed to locate all nodes containing account: %v", accID))
	}
}

func (node *Node) findAccountQuery(target [4]byte, respChan chan bool, accID [5]uint32) {
	rpc := GenerateRPC(target, node.Contact)
	rpc.FindAccount(accID)
	res, err := node.Send(rpc)
	if err == nil {
		respChan <- res.findAccountSucc
	}
	respChan <- false
}

func (node *Node) DisplayAccount(accID [5]uint32) (string, error) {
	validators := node.FindNode(accID)
	log.Printf("found %d validators", len(validators))
	for _, con := range validators {
		rpc := GenerateRPC(con.IP(), node.Contact)
		rpc.DisplayAccount(accID)
		res, _ := node.Send(rpc)
		if res.displayString != "" {
			return res.displayString, nil
		}
	}
	return "", errors.New("did not find account")
}

func (node *Node) LockAccount(accID [5]uint32) ([]Contact, []chan RPC, chan RPC) {
	valGroup, _ := node.FindAccount(accID)
	valChan := make([]chan RPC, 0, REPLICATION)
	leaderChan := make(chan RPC, REPLICATION)

	for _, val := range valGroup {
		rpc := GenerateRPC(val.IP(), node.Contact)
		rpc.OverrideID(RelativeDistance(node.ID(), val.ID()))
		rpc.LockAccount(accID, leaderChan)

		go node.Send(rpc)
	}

	for range valGroup {
		resp, ok := <-leaderChan
		if !ok {
			log.Printf("node: %10v, leader chan for account %10v unexpectedly closed\n", node.ID(), accID)
		} else {
			if node.debug {
				log.Printf("node: %10v took lock for account %10v on node %10v\n", node.ID(), accID, resp.sender.ID())
			}
		}
		valChan = append(valChan, resp.lockChan)
	}

	return valGroup, valChan, leaderChan
}
