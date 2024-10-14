package kademlia

import "sync"

type handler struct {
	lock  sync.RWMutex
	table map[[5]uint32]chan RPC
}

type Network struct {
	listener   chan RPC
	sender     chan RPC
	serverIP   [4]byte
	masterNode Contact
	handler
}
