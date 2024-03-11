package scalegraph

import "log"

// Parses the RPC and dispatch subroutines
func (n *Node) Handler(rpc RPC) {
	switch rpc.CMD {
	case PING:
		n.handlePing(rpc)
	case PONG:
		n.handlePong(rpc)
	case FIND_NODE:
		n.handleFindNode(rpc)
	}
}

// Handles the internal logic for a received ping RPC.
func (n *Node) handlePing(rpc RPC) {
	if DEBUG {
		log.Printf("[node] - received %+v, from %+v", rpc.CMD, rpc.Sender.id)
	}
	n.AddContact(rpc.Sender)
	resp := GenerateResponse(PONG, rpc.ID, rpc.Sender.ip, n.contact)
	resp.Pong()
	n.network.Send(resp)
}

// Handles the internal logic for a received pong RPC.
func (n *Node) handlePong(rpc RPC) {
	n.AddContact(rpc.Sender)
	if DEBUG {
		log.Printf("[node] - received %+v, from %+v", rpc.CMD, rpc.Sender.id)
	}
	if rpc.Sender.IP() != n.serverIP {
		if DEBUG {
			log.Printf("[node] - adding contact to routing table, id: %+v", rpc.Sender.id)
		}
	} else {
		log.Println("[node] - received pong from server")
	}
	if DEBUG {
		log.Printf("[node] - received pong from %+v", rpc.Sender)
	}
}

func (n *Node) handleStore(rpc RPC) {}

func (n *Node) handleStoreResponse(rpc RPC) {}

func (n *Node) handleStoresResponse(rpc RPC) {}

func (n *Node) handleFindNode(rpc RPC) {
	n.AddContact(rpc.Sender)
	res, err := n.FindXClosest(REPLICATION, rpc.FindTarget)
	if err != nil {
		log.Printf("%+v - find node error: %+v", n.ID(), err)
	}
	resp := GenerateResponse(FIND_NODE_RESPONSE, rpc.ID, rpc.Sender.ip, n.contact)
	resp.FindNodeResponse(res)
	n.network.Send(resp)
}

func (n *Node) handleFindNodeResponse(rpc RPC) {}

func (n *Node) handleFindWallet(rpc RPC) {}

func (n *Node) handleFindWalletResponse(rpc RPC) {}

func (n *Node) handleProposeRequest(rpc RPC) {}

func (n *Node) handlePoposeAccept(rpc RPC) {}

func (n *Node) handlePropose(rpc RPC) {}

func (n *Node) handleProposeValidate(rpc RPC) {}
