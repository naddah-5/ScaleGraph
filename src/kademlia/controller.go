package kademlia

// Controller handles the logic for receiving RPC's

// Similar to the network listening loop, but here its purpose is to accept
// the incoming RPC's from lower levels in the system.
func (node *Node) InputLoop() {
	for {
		rpc := <-node.controller
		node.handler(rpc)
	}
}

func (node *Node) handler(rpc RPC) {
	node.AddContact(rpc.Sender)
	switch rpc.CMD {
	case PING:
		node.HandlePing(rpc)
	case STORE_WALLET:
		node.HandleStoreWallet(rpc)
	}
}

// Response logic for an incoming ping RPC.
// Simply respond with a ping marked as a response.
func (node *Node) HandlePing(rpc RPC) {
	resp := GenerateResponse(rpc.ID, node.Contact)
	resp.Ping(rpc.Sender.IP())
	node.Send(resp)
}

// Logic for sending a store wallet RPC.
func (node *Node) StoreWallet() {

}

// Response logic for an incoming store RPC.
func (node *Node) HandleStoreWallet(rpc RPC) {
	
}


