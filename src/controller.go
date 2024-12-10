package src

import "main/src/kademlia"

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
		node.Ping(rpc)
	case kademlia.STORE_WALLET:
		node.StoreWallet(rpc)
	}
}

// Response logic for an incoming ping RPC.
func (node *Node) Ping(rpc kademlia.RPC) {
	node.AddContact(rpc.Sender)
	resp := kademlia.GenerateResponse(rpc.ID, node.Contact)
	resp.Pong(rpc.Sender.IP())
	node.Send(resp)
}

// Response logic for an incoming store RPC.
func (node *Node) StoreWallet(rpc kademlia.RPC) {
	
}
