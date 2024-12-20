package kademlia

// Protocol handles the logic for sending RPC's

// Logic for sending a ping RPC.
func (node *Node) Ping(ip [4]byte) {
	rpc := GenerateRPC(node.Contact)
	rpc.Ping(ip)
	go node.Send(rpc)
}

func (node *Node) FindNode(target [5]uint32) []Contact {
	initNodes, _ := node.FindXClosest(REPLICATION, target)
	return node.FindNodeRec(initNodes, target)
}

func (node *Node) FindNodeRec(initNodes []Contact, target [5]uint32) []Contact {
	retNodes := make([]Contact, 0, REPLICATION*REPLICATION)
	respChan := make(chan []Contact, REPLICATION)
	for _, tar := range initNodes {
		rpc := GenerateRPC(node.Contact)
		rpc.FindNode(target)
		go func(target Contact, rpc RPC, respChan chan []Contact) {
			resp, err := node.Send(rpc)
			if (err != nil) || (resp.FoundNodes == nil) {
				respChan <- make([]Contact, 0)
				return
			}
			respChan <- resp.FoundNodes
			return
		}(tar, rpc, respChan)
	}
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
		return node.FindNodeRec(retNodes, target)
	}
}
