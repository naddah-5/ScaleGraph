package kademlia

import "log"

// Controller handles the logic for receiving RPC's

func (node *Node) Handler(rpc RPC) {
	go node.AddContact(rpc.sender)
	log.Printf("handling: %s", rpc.cmd)
	switch rpc.cmd {
	case PING:
		node.HandlePing(rpc)
	case STORE_WALLET:
		node.HandleStoreWallet(rpc)
	}
}

// Response logic for an incoming ping RPC.
// Simply respond with a ping marked as a response.
func (node *Node) HandlePing(rpc RPC) {
	resp := GenerateResponse(rpc.id, node.Contact)
	resp.Pong(rpc.sender.IP())
	res, err := node.Send(resp)
	if err != nil {
		log.Printf("[ERROR] - %v\n%s", node.ID(), err.Error())
	}
	log.Printf(res.Display() + "\n")
}

func (node *Node) HandleFindNode(rpc RPC) {

}

// Logic for sending a store wallet RPC.
func (node *Node) StoreWallet() {

}

// Response logic for an incoming store RPC.
func (node *Node) HandleStoreWallet(rpc RPC) {
	
}


