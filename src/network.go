package scalegraph

import (
	"errors"
	"log"
	"time"
)

type network struct {
	listener chan RPC
	sender   chan RPC
	serverIP [4]byte
	handler  map[[5]uint32]chan RPC
}

func NewNetwork(ln chan RPC, sn chan RPC, servIP [4]byte) network {
	newNetwork := network{
		listener: ln,
		sender:   sn,
		serverIP: servIP,
		handler:  make(map[[5]uint32]chan RPC, 100),
	}
	return newNetwork
}

func (n *network) Send(rpc RPC) (RPC, error) {
	if DEBUG {
		log.Printf("%+v: sending rpc: %+v, to IP: %+v", rpc.Sender.id, rpc, rpc.Receiver)
	}
	if rpc.response {
		n.sender <- rpc
	} else {
		respChan := make(chan RPC)
		// potential (low probability) collisions, would not break but deletes old channel
		n.handler[rpc.id] = respChan
		n.sender <- rpc
		select {
		case res := <-respChan:
			return res, nil
		case <-time.After(TIMEOUT):
			return rpc, errors.New("RPC timeout")
		}
	}
	return rpc, nil
}

func (n *network) Listen(node *Node) error {
	for {
		rpc, ok := <-n.listener
		if !ok {
			return errors.New("server not responding")
		}
		if DEBUG {
			log.Printf("%+v: received rpc: %+v", node.Me.id, rpc)
		}
		if rpc.response {
			n.handler[rpc.id] <- rpc
			close(n.handler[rpc.id])
		} else {
			node.Handler(rpc)
		}
	}
}
