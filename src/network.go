package scalegraph

import (
	"errors"
	"log"
	"time"
)

type handler struct {
	set map[[5]uint32]chan RPC
}

func NewHandler(buffer int) handler {
	ch := make(map[[5]uint32]chan RPC, buffer)
	return handler{ch}
}

func (handler *handler) Add(id [5]uint32) chan RPC {
	respChan := make(chan RPC)
	// potential (low probability) collisions, would not break but deletes old channel
	_, exists := handler.set[id]
	if exists {
		log.Printf("rpc handler overwrite occured, id - %+v", id)
	}
	handler.set[id] = respChan

	return respChan
}

func (handler *handler) Retrieve(id [5]uint32) (chan RPC, error) {
	respChan, ok := handler.set[id]
	if !ok {
		return nil, errors.New("invalid RPC id")
	}
	return respChan, nil
}

func (handler *handler) Drop(id [5]uint32) {
	delete(handler.set, id)
}

type network struct {
	listener chan RPC
	sender   chan RPC
	serverIP [4]byte
	handler
}

func NewNetwork(ln chan RPC, sn chan RPC, servIP [4]byte) network {
	newNetwork := network{
		listener: ln,
		sender:   sn,
		serverIP: servIP,
		handler:  NewHandler(100),
	}
	return newNetwork
}

func (net *network) Send(rpc RPC) (RPC, error) {
	if DEBUG {
		log.Printf("%+v: sending rpc: %+v, to IP: %+v\n", rpc.Sender.id, rpc, rpc.Receiver)
	}
	if rpc.response {
		net.sender <- rpc
	} else {
		respChan := net.handler.Add(rpc.id)
		net.sender <- rpc
		select {
		case res := <-respChan:
			return res, nil
		case <-time.After(TIMEOUT):
			return rpc, errors.New("RPC timeout")
		}
	}
	return rpc, nil
}

func (net *network) Listen(node *Node) error {
	for {
		rpc, ok := <-net.listener
		if !ok {
			return errors.New("server not responding")
		}
		if DEBUG {
			log.Printf("%+v: received rpc: %+v", node.Me.id, rpc)
		}
		if rpc.response {
			respChan, err := net.handler.Retrieve(rpc.id)
			if err != nil {
				log.Printf(err.Error())
				continue
			}
			respChan <- rpc
		} else {
			node.Handler(rpc)
		}
	}
}
