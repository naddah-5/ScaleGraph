package scalegraph

import (
	"log"
)

// Parses the RPC and dispatch subroutines
func (n *Node) Controller(rpc RPC) {
	switch rpc.CMD {
	case PING:
		n.controlPing(rpc)
	case PONG:
		n.controlPong(rpc)
	case FIND_NODE:
		n.controlFindNode(rpc)
	case SEND:
		n.controlSend(rpc)
	}
}

// Redirection command from the simnet, the redirection target will be stored in the sender.
func (node *Node) controlSend(rpc RPC) {
	rpc.Redirect(rpc.Sender.IP())
	rpc.Sender = node.contact
	rpc.CMD = rpc.order
}

// Handles the internal logic for a received ping RPC.
func (n *Node) controlPing(rpc RPC) {
	if DEBUG {
		log.Printf("[node] - received %+v, from %+v", rpc.CMD, rpc.Sender.id)
	}
	n.AddContact(rpc.Sender)
	resp := GenerateResponse(PONG, rpc.ID, rpc.Sender.ip, n.contact)
	resp.Pong()
	n.network.Send(resp)
}

// Handles the internal logic for a received pong RPC.
func (n *Node) controlPong(rpc RPC) {
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

func (n *Node) controlStore(rpc RPC) {
	n.AddContact(rpc.Sender)
}


// Handles the internal logic for find node rpc.
func (n *Node) controlFindNode(rpc RPC) {
	n.AddContact(rpc.Sender)
	res, err := n.FindXClosest(REPLICATION, rpc.FindTarget)
	if err != nil {
		log.Printf("%+v - find node error: %+v", n.ID(), err)
	}
	resp := GenerateResponse(FIND_NODE_RESPONSE, rpc.ID, rpc.Sender.ip, n.contact)
	resp.FindNodeResponse(res, rpc.FindTarget)
	n.network.Send(resp)
}

//func (n *Node) controlFindNodeResponse(rpc RPC) {
//	go func() {
//		for _, node := range rpc.KNodes {
//			go func(node contact) {
//				ping := GenerateRPC(PING, n.contact, node.IP())
//				resp, err := n.network.Send(ping)
//				if err != nil {
//					return
//				}
//				go n.Controller(resp)
//			}(node)
//		}
//	}()
//}


func (n *Node) controlFindWallet(rpc RPC) {}

func (n *Node) controlFindWalletResponse(rpc RPC) {}

func (n *Node) controlProposeRequest(rpc RPC) {}

func (n *Node) controlPoposeAccept(rpc RPC) {}

func (n *Node) controlPropose(rpc RPC) {}

func (n *Node) controlProposeValidate(rpc RPC) {}
