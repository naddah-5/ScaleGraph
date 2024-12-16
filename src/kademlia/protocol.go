package kademlia

// Protocol handles the logic for sending RPC's

// Logic for sending a ping RPC.
func (node *Node) Ping(ip [4]byte) {
	rpc := GenerateRPC(node.Contact)
	rpc.Ping(ip)
	go node.Send(rpc)
}

func (node *Node) FindNode(id [5]uint32) {
	initNodes, _ := node.FindXClosest(REPLICATION, id)
}
