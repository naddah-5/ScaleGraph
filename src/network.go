package scalegraph

import (
	"errors"
	"log"
	"time"
)

type handler struct {
	set map[[5]uint32]chan RPC
}

// Initializes and returns a new handler.
func NewHandler(buffer int) handler {
	ch := make(map[[5]uint32]chan RPC, buffer)
	return handler{ch}
}

// Creates a RPC id handle corresponding to the given RPC id.
// Returns the response channel.
// Note that this can over write existing handles.
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

// Returns the response channel for a given RPC id.
// Returns an error if there is no matching RPC id.
func (handler *handler) Retrieve(id [5]uint32) (chan RPC, error) {
	respChan, ok := handler.set[id]
	if !ok {
		return nil, errors.New("invalid RPC id")
	}
	return respChan, nil
}

// Removes the given id from the active RPC map.
// If there is no match, does nothing.
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

// Sends a RPC and creates a corresponding RPC id handle.
// Returns an error if the response eceedes the timeout.
func (net *network) Send(rpc RPC) (RPC, error) {
	if DEBUG {
		log.Printf("sending rpc: %+v, to IP: %+v\n", rpc, rpc.receiver)
	}
	if rpc.response {
		net.sender <- rpc
	} else {
		respChan := net.Add(rpc.walletID)
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

// Start a listener on the network channel.
// Returns an error if the channel closes.
func (net *network) Listen(node *Node) error {
	for {
		rpc, ok := <-net.listener
		if !ok {
			return errors.New("server not responding")
		}
		if DEBUG {
			log.Printf("%+v: received rpc: %+v, is repsonse: %+v", node.id, rpc, rpc.response)
		}
		if rpc.response {
			respChan, err := net.Retrieve(rpc.walletID)
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
