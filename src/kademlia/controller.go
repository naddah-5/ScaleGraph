package kademlia

import "log"

// Controller handles the logic for receiving RPC's

func (node *Node) Handler(rpc RPC) {
	go node.AddContact(rpc.Sender)
	log.Printf("hanling: %s", rpc.CMD)
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
	res, _ := node.Send(resp)
	log.Printf(res.Display() + "\n")
}

// Logic for sending a store wallet RPC.
func (node *Node) StoreWallet() {

}

// Response logic for an incoming store RPC.
func (node *Node) HandleStoreWallet(rpc RPC) {
	
}


