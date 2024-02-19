package scalegraph

import (
	"errors"
)

type network struct {
	listener chan RPC
	sender   chan RPC
}

func NewNetwork(ln chan RPC, sn chan RPC) network {
	newNetwork := network{
		listener: ln,
		sender: sn,
	}
	return newNetwork
}

func (n *network) Send(msg RPC) {
	n.sender<-msg
}

func (n *network) Listen() error {
	for {
		rpc, ok := <-n.listener
		if !ok {
			return errors.New("server not responding")
		}
		go Handler(rpc)
	}
}
