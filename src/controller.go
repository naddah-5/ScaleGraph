package src

import "main/src/kademlia"

func (node *Node) InputLoop() {
	for {
		rpc := <-node.controller
		
	}
}

func (node *Node) ping(rpc kademlia.RPC) {

}
