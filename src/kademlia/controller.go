package kademlia

import "log"

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
		node.handleStoreWallet(rpc)
	case FIND_NODE:
		node.handleFindNode(rpc)
	}
}


// Response logic for an incoming ping RPC.
// Simply respond with a ping marked as a response.
func (node *Node) handlePing(rpc RPC) {
	resp := GenerateResponse(rpc.id, node.Contact)
	resp.Pong(rpc.sender.IP())
	node.Send(resp)
}

func (node *Node) handleFindNode(rpc RPC) {
	res, err := node.FindXClosest(REPLICATION, rpc.findNodeTarget)
	if err != nil {
		log.Printf("Node %v - Handle Find Node Error\n%s", node.ID(), err.Error())
	}
	resp := GenerateResponse(rpc.id, node.Contact)
	resp.FoundNodes(rpc.sender.IP(), rpc.findNodeTarget, res)
	if node.debug {
		log.Printf("[DEBUG]\nresponding to find node with RPC:\n%s", rpc.Display())
	}
	go node.Send(resp)
}

// Logic for sending a store wallet RPC.
func (node *Node) storeWallet() {

}

// Response logic for an incoming store RPC.
func (node *Node) handleStoreWallet(rpc RPC) {

}
