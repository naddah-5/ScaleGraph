package kademlia

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type table struct {
	content map[[5]uint32]chan RPC
	sync.RWMutex
}

func NewTable() *table {
	ch := make(map[[5]uint32]chan RPC, 1024)
	return &table{
		content: ch,
	}
}

// Creates a RPC channel corresponding to the given id.
// Channel is entered into network table and returned.
// Returns an error if id is already in use.
func (table *table) Add(id [5]uint32) (chan RPC, error) {
	table.Lock()
	defer table.Unlock()

	_, exists := table.content[id]
	if exists {
		return make(chan RPC), errors.New("RPC id in use")
	}

	respChan := make(chan RPC)
	table.content[id] = respChan
	return respChan, nil
}

// Returns the matching RPC channel, or an error if there is no match.
func (table *table) RetrieveChan(id [5]uint32) (chan RPC, error) {
	table.Lock()
	defer table.Unlock()

	ch, ok := table.content[id]
	if !ok {
		return nil, errors.New("no matching RPC id")
	}
	delete(table.content, id)
	return ch, nil
}

// Removes entry with id from table.
func (table *table) DropChan(id [5]uint32) {
	table.Lock()
	defer table.Unlock()
	delete(table.content, id)
}

type Network struct {
	listener   chan RPC
	sender     chan RPC
	controller    chan RPC
	serverIP   [4]byte
	masterNode Contact
	patience   int // Waiting time before giving up on reponse
	*table
}

// Returns a network pointer.
func NewNetwork(listener chan RPC, sender chan RPC, controller chan RPC, serverIP [4]byte, master Contact) *Network {
	newNetwork := Network{
		listener:   listener,
		sender:     sender,
		controller:    controller,
		serverIP:   serverIP,
		masterNode: master,
		table:      NewTable(),
	}
	return &newNetwork
}

// Sends a RPC and creates a corresponding RPC id handle.
// Returns an error if the response exceedes the timeout.
func (net *Network) Send(rpc RPC) (RPC, error) {
	if rpc.response {
		net.sender <- rpc
	} else {
		rpc.id = RandomID()
		respChan, err := net.Add(rpc.id)
		// loops through randomly generated rpc id's until a free one is found.
		for err != nil {
			rpc.id = RandomID()
			respChan, err = net.Add(rpc.id)
		}
		net.sender <- rpc
		select {
		case res := <-respChan:
			return res, nil
		case <-time.After(time.Duration(net.patience) * time.Second):
			net.DropChan(rpc.id)
			break
		}
	}
	return rpc, errors.New("timeout")
}

// Start a listener on the network channel.
// Returns an error if the channel closes.
func (net *Network) Listen() error {
	for {
		rpc, ok := <-net.listener
		if !ok {
			return errors.New("server down")
		}
		go net.route(rpc)
	}
}

// Routes the rpc to the appropriate components.
// If the rpc is a response it tries to route it to that channel, otherwise routes it to the controller.
func (net *Network) route(rpc RPC) {
	if rpc.response {
		respChan, err := net.RetrieveChan(rpc.id)
		if err != nil {
			errMSg := fmt.Sprintf("[ERROR] - possible time out\n error: %s", err.Error())
			log.Println(errMSg)
			return
		}
		net.DropChan(rpc.id)
		respChan <- rpc
	} else {
		net.controller <- rpc
	}
}

