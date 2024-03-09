package scalegraph

import "log"

// Parses the RPC and dispatch subroutines
func (n *Node) Handler(rpc RPC) {
	switch rpc.CMD {
	case PING:
		n.handlePing(rpc)
	case PONG:
		n.handlePong(rpc)
	}
}

func (n *Node) handlePing(rpc RPC) {
	resp := GenerateResponse(PONG, rpc.ID, rpc.Sender.ip, n.contact)
	resp.Pong()
	n.network.Send(resp)
}

func (n *Node) handlePong(rpc RPC) {
	if DEBUG {
		log.Printf("received %+v, from %+v", rpc.CMD, rpc.Sender.id)
	}
	if rpc.Sender.IP() != n.network.serverIP {
		n.routingTable.AddContact(rpc.Sender)
		if DEBUG {
			log.Printf("adding contact to routing table, id: %+v", rpc.Sender.id)
		}
	}
	log.Printf("received pong from %+v", rpc.Sender)
}

func (n *Node) handleStore(rpc RPC) {}

func (n *Node) handleStoreResponse(rpc RPC) {}

func (n *Node) handleStoresResponse(rpc RPC) {}

func (n *Node) handleFindNode(rpc RPC) {}

func (n *Node) handleFindNodeResponse(rpc RPC) {}

func (n *Node) handleFindWallet(rpc RPC) {}

func (n *Node) handleFindWalletResponse(rpc RPC) {}

func (n *Node) handleProposeRequest(rpc RPC) {}

func (n *Node) handlePoposeAccept(rpc RPC) {}

func (n *Node) handlePropose(rpc RPC) {}

func (n *Node) handleProposeValidate(rpc RPC) {}
