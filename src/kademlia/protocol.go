package kademlia

import "log"

// Protocol handles the logic for sending RPC's

// Logic for sending a ping RPC.
func (node *Node) Ping(target [4]byte) {
	rpc := GenerateRPC(node.Contact)
	rpc.Ping(target)
	go node.Send(rpc)
}

func (node *Node) FindNode(target [5]uint32) []Contact {
	initNodes, _ := node.FindXClosest(REPLICATION, target)
	return node.FindNodeRec(initNodes, target)
}

func (node *Node) FindNodeRec(initNodes []Contact, target [5]uint32) []Contact {
	retNodes := make([]Contact, 0, REPLICATION*REPLICATION)
	respChan := make(chan any, REPLICATION)

	// Launch parallel queries to initial nodes.
	for _, tar := range initNodes {
		rpc := GenerateRPC(node.Contact)
		rpc.FindNode(tar.ID())
		go node.nodeQuery(rpc, respChan)
	}

	// Extract results from parallel query.
	for i := 0; i < REPLICATION; i++ {
		resp, ok := <-respChan
		if !ok {
			continue
		}

		// Assert correct type for the response from respChan, else log error and continue.
		if respCon, ok := resp.([]Contact); ok {
			retNodes = append(retNodes, respCon...)
		} else {
			log.Println("[ERROR] - received incorrect response from find node query")
			continue
		}
	}

	SortContactsByDistance(&retNodes, target)
	if len(retNodes) == 0 || retNodes[0].ID() == initNodes[0].ID() {
		return initNodes
	} else {
		return node.FindNodeRec(retNodes, target)
	}
}

// Sends the given RPC and returns the reponse to the provided channel.
// If the RPC times out or returns an error, returns an empty contact.
// NOTE that you must assert the type of the result from respChan.
func (node *Node) nodeQuery(rpc RPC, respChan chan any) {
	resp, err := node.Send(rpc)
	if (err != nil) || (resp.FoundNodes == nil) {
		respChan <- make([]Contact, 0)
		return
	}
	respChan <- resp.FoundNodes
	return

}
