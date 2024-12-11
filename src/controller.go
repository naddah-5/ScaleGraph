package src

import "main/src/kademlia"

// Controller handles the logic for sending and receiving RPC's

// Similar to the network listening loop, but here its purpose is to accept 
// the incoming RPC's from lower levels in the system.
func (node *Node) InputLoop() {
	for {
		rpc := <-node.controller
		node.handler(rpc)
	}
}

func (node *Node) handler(rpc kademlia.RPC) {
	switch rpc.CMD {
	case kademlia.PING:
		node.HandlePing(rpc)
	case kademlia.STORE_WALLET:
		node.HandleStoreWallet(rpc)
	}
}

// Logic for sending a ping RPC.
func (node *Node) Ping() {

}

// Response logic for an incoming ping RPC.
func (node *Node) HandlePing(rpc kademlia.RPC) {
	node.AddContact(rpc.Sender)
	resp := kademlia.GenerateResponse(rpc.ID, node.Contact)
	resp.Pong(rpc.Sender.IP())
	node.Send(resp)
}

// Logic for sending a store wallet RPC.
func (node *Node) StoreWallet() {

}

// Response logic for an incoming store RPC.
func (node *Node) HandleStoreWallet(rpc kademlia.RPC) {
	
}
