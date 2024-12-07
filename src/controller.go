package src

import "main/src/kademlia"

func (node *Node) InputLoop() {
	for {
		rpc := <-node.controller
		node.handler(rpc)
	}
}

func (node *Node) handler(rpc kademlia.RPC) {
	switch rpc.CMD {
	case kademlia.PING:
		node.ping(rpc)
	case kademlia.STORE:
		node.store(rpc)
	}
}

func (node *Node) ping(rpc kademlia.RPC) {
	node.AddContact(rpc.Sender)
	resp := kademlia.GenerateResponse(rpc.ID, node.Contact)
	

}

func (node *Node) store(rpc kademlia.RPC) {

}
