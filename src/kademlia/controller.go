package kademlia

import (
	"errors"
	"fmt"
	"log"
)

// Controller handles the logic for receiving RPC's

func (node *Node) Handler(rpc RPC) {
	if node.debug {
		log.Printf("[DEBUG]\nNode %v - handling rpc:\n%s", node.ID(), rpc.Display())
	}
	go node.AddContact(rpc.sender)
	switch rpc.cmd {
	case PING:
		node.handlePing(rpc)
	case STORE_ACCOUNT:
		node.handleStoreAccount(rpc)
	case FIND_NODE:
		node.handleFindNode(rpc)
	}
}

// Response logic for an incoming ping RPC.
// Simply respond with a ping marked as a response.
func (node *Node) handlePing(rpc RPC) {
	resp := GenerateResponse(rpc.id, rpc.sender.IP(), node.Contact)
	resp.Pong()
	node.Send(resp)
}

func (node *Node) handleFindNode(rpc RPC) {
	res, err := node.FindXClosest(REPLICATION, rpc.findNodeTarget)
	if err != nil {
		log.Printf("Node %v - Handle Find Node Error\n%s", node.ID(), err.Error())
	}
	resp := GenerateResponse(rpc.id, rpc.sender.IP(), node.Contact)
	resp.FoundNodes(rpc.findNodeTarget, res)
	if node.debug {
		log.Printf("[DEBUG]\nresponding to find node with RPC:\n%s", rpc.Display())
	}
	go node.Send(resp)
}

// Response logic for an incoming store RPC.
func (node *Node) handleStoreAccount(rpc RPC) {
	err := node.scalegraph.AddAccount(rpc.accountID)
	resp := GenerateResponse(rpc.id, rpc.sender.IP(), node.Contact)
	resp.StoredAccount(rpc.accountID, err == nil)
	go node.Send(resp)
}

// Optional check to verify the node does not know it's not part of the validator group.
func (node *Node) storeAccountCheck(accID [5]uint32) error {
	validators, _ := node.FindXClosest(REPLICATION, accID)
	validator := CloserNode(node.ID(), validators[len(validators)-1].ID(), accID)
	if !validator {
		log.Printf("node %v: received incorrect store account RPC", node.ID())
		return errors.New(fmt.Sprintf("node %v is not a validator for account: %v", node.ID(), accID))
	}
	return nil
}

func (node *Node) handleFindAccount(rpc RPC) {
	present, _ := node.scalegraph.FindAccount(rpc.accountID)
	resp := GenerateResponse(rpc.id, rpc.sender.IP(), node.Contact)
	resp.FoundAccount(rpc.accountID, present)
	go node.Send(resp)
}
