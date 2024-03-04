package scalegraph

import (
	"errors"
	"log"
)

type network struct {
	listener chan RPC
	sender   chan RPC
	serverIP [4]byte
}

func NewNetwork(ln chan RPC, sn chan RPC, servIP [4]byte) network {
	newNetwork := network{
		listener: ln,
		sender:   sn,
		serverIP: servIP,
	}
	return newNetwork
}

func (n *network) Send(rpc RPC) {
	if DEBUG {
		log.Printf("%+v: sending rpc: %+v, to IP: %+v", rpc.Sender.nodeID, rpc, rpc.Receiver)
	}
	n.sender <- rpc
}

func (n *network) Listen(node *Node) error {
	for {
		rpc, ok := <-n.listener
		if !ok {
			return errors.New("server not responding")
		}
		if DEBUG {
			log.Printf("%+v: received rpc: %+v", node.Me.nodeID, rpc)
		}
		node.Handler(rpc)
	}
}
