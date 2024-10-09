package scalegraph

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type handler struct {
	lock sync.RWMutex
	set  map[[5]uint32]chan RPC
}

// Initializes and returns a new handler.
func NewHandler(buffer int) handler {
	ch := make(map[[5]uint32]chan RPC, buffer)
	return handler{
		lock: sync.RWMutex{},
		set:  ch,
	}
}

// Creates a RPC id handle corresponding to the given RPC id.
// Returns the response channel.
// Note that this can over write existing handles.
func (handler *handler) Add(id [5]uint32) chan RPC {
	respChan := make(chan RPC)
	// potential (low probability) collisions, would not break but over writes old channel
	handler.lock.Lock()
	defer handler.lock.Unlock()
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
	handler.lock.RLock()
	respChan, ok := handler.set[id]
	handler.lock.RUnlock()
	if !ok {
		return nil, errors.New("invalid RPC id")
	}
	return respChan, nil
}

// Removes the given id from the active RPC map.
// If there is no match, does nothing.
func (handler *handler) Drop(id [5]uint32) {
	handler.lock.Lock()
	delete(handler.set, id)
	handler.lock.Unlock()
}

type network struct {
	listener chan RPC
	sender   chan RPC
	serverIP [4]byte
	master   [4]byte
	handler
}

func NewNetwork(ln chan RPC, sn chan RPC, servIP [4]byte, master [4]byte) *network {
	newNetwork := network{
		listener: ln,
		sender:   sn,
		serverIP: servIP,
		master:   master,
		handler:  NewHandler(100),
	}
	return &newNetwork
}

// Sends a RPC and creates a corresponding RPC id handle.
// Returns an error if the response exceedes the timeout.
func (net *network) Send(rpc RPC) (RPC, error) {
	if rpc.response {
		net.sender <- rpc
	} else {
		respChan := net.Add(rpc.id)
		net.sender <- rpc
		select {
		case res := <-respChan:
			return res, nil
		case <-time.After(TIMEOUT):
			net.Drop(rpc.id)
			break
		}
	}
	return rpc, errors.New("timeout")
}

// Start a listener on the network channel.
// Returns an error if the channel closes.
func (net *network) Listen(node *Node) error {
	for {
		rpc, ok := <-net.listener
		if !ok {
			return errors.New("server not responding")
		}
		go net.understand(node, rpc)

	}
}

func (net *network) understand(node *Node, rpc RPC) {
	if rpc.response {
		respChan, err := net.Retrieve(rpc.id)
		if err != nil {
			err := fmt.Sprintf(err.Error())
			data := fmt.Sprintf("ERROR DATA: %+v, id: %+v, sender id: %+v", rpc.CMD, rpc.id, rpc.sender.id)
			log.Printf("%s\n%s", err, data)
			return
		}
		net.Drop(rpc.id)
		respChan <- rpc
	} else {
		node.Controller(rpc)
	}
}
