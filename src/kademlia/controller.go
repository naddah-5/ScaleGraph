package kademlia


// Controller handles the logic for sending and receiving RPC's

// Similar to the network listening loop, but here its purpose is to accept
// the incoming RPC's from lower levels in the system.
func (node *Node) InputLoop() {
	for {
		rpc := <-node.controller
		node.handler(rpc)
	}
}

func (node *Node) handler(rpc RPC) {
	switch rpc.CMD {
	case PING:
		node.HandlePing(rpc)
	case PONG:
		node.HandlePong(rpc)
	case STORE_WALLET:
		node.HandleStoreWallet(rpc)
	}
}

// Logic for sending a ping RPC.
func (node *Node) Ping() {

}

// Response logic for an incoming ping RPC.
func (node *Node) HandlePing(rpc RPC) {
	node.AddContact(rpc.Sender)
	resp := GenerateResponse(rpc.ID, node.Contact)
	resp.Pong(rpc.Sender.IP())
	node.Send(resp)
}

func (node *Node) HandlePong(rpc RPC) {

}

// Logic for sending a store wallet RPC.
func (node *Node) StoreWallet() {

}

// Response logic for an incoming store RPC.
func (node *Node) HandleStoreWallet(rpc RPC) {
	
}

