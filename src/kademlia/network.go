package kademlia

import (
	"errors"
	"sync"
)

type handler struct {
	lock  sync.RWMutex
	table map[[5]uint32]chan RPC
}

type Network struct {
	listener   chan RPC
	sender     chan RPC
	serverIP   [4]byte
	masterNode Contact
	*table
}

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

func (table *table) RetrieveChan(id [5]uint32) (chan RPC, error) {

	return make(chan RPC), nil
}
