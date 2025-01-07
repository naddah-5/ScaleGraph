package kademlia

import "log"

// Protocol handles the logic for sending RPC's

// Logic for sending a ping RPC.
func (node *Node) Ping(address [4]byte) {
	rpc := GenerateRPC(node.Contact)
	rpc.Ping(address)
	res, err := node.Send(rpc)
	if err != nil {
		log.Printf("[ERROR] - %s", err.Error())
	}
	node.AddContact(res.sender)
}

func (node *Node) FindNode(target [5]uint32) []Contact {
	initNodes, _ := node.FindXClosest(REPLICATION, target)
	return node.findNodeRec(initNodes, target)
}

func (node *Node) findNodeRec(initNodes []Contact, target [5]uint32) []Contact {
	retNodes := make([]Contact, 0, REPLICATION*REPLICATION)
	respChan := make(chan []Contact, REPLICATION)

	// Launch parallel queries to initial nodes.
	for _, tar := range initNodes {
		rpc := GenerateRPC(node.Contact)
		rpc.FindNode(tar.IP(), tar.ID())
		go node.nodeQuery(rpc, respChan)
	}

	// Extract results from parallel query.
	for i := 0; i < REPLICATION; i++ {
		resp, ok := <-respChan
		if !ok {
			continue
		}

		retNodes = append(retNodes, resp...)
	}

	SortContactsByDistance(&retNodes, target)
	if len(retNodes) == 0 || retNodes[0].ID() == initNodes[0].ID() {
		return initNodes
	} else {
		return node.findNodeRec(retNodes, target)
	}
}

// Sends the given RPC and returns the reponse to the provided channel.
// If the RPC times out or returns an error, returns an empty contact.
// NOTE that you must assert the type of the result from respChan.
func (node *Node) nodeQuery(rpc RPC, respChan chan []Contact) {
	resp, err := node.Send(rpc)
	if err != nil {
		log.Printf("[ERROR] - %s\nin node %v with rpc:\n%s\n", err.Error(), node.ID(), rpc.Display())
		respChan <- resp.foundNodes
		return
	}
	respChan <- resp.foundNodes
	return

}
