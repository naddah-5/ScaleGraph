package kademlia

import ()

type cmd int

const (
	PING cmd = iota
	PONG
	STORE
	FIND_NODE
	FIND_WALLET
	PROPOSE_TRANSACTION
	ACCEPT_TRANSACTION
	SEND
	SUBMIT_WALLET
	SHOW_WALLET
)

func (c cmd) String() string {
	switch c {
	case PING:
		return "PING"
	case PONG:
		return "PONG"
	case STORE:
		return "STORE"
	case FIND_NODE:
		return "FIND_NODE"
	case FIND_WALLET:
		return "FIND_WALLET"
	}
	return "unknown"
}

type RPC struct {
	ID       [5]uint32
	Response bool
	Receiver [4]byte
}
