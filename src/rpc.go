package scalegraph

import (
	"container/list"
	"log"
)

type CMD int

const (
	PING CMD = iota
	PONG
	STORE
	STORE_RESPONSE
	FIND_NODE
	FIND_NODE_RESPONSE
	FIND_VALUE
	FIND_VALUE_RESPONSE
)

func (c CMD) String() string {
	switch c {
	case PING:
		return "PING"
	case PONG:
		return "PONG"
	case STORE:
		return "STORE"
	case STORE_RESPONSE:
		return "STORE_RESPONSE"
	case FIND_NODE:
		return "FIND_NODE"
	case FIND_NODE_RESPONSE:
		return "FIND_NODE_RESPONSE"
	case FIND_VALUE:
		return "FIND_VALUE"
	case FIND_VALUE_RESPONSE:
		return "FIND_VALUE_RESPONSE"
	}
	return "unknown"
}

// Contains fields for all RPC data
// note that fields may be nil.
type RPC struct {
	CMD
	ID          [5]uint32
	Sender      contact
	wallet
	WalletKey   []byte
	Transaction []byte
	KNodes      []contact
	Acknowledge bool
}

func GenerateRPC(cmd CMD, sender contact) RPC {
	newRPC := RPC{
		CMD:    cmd,
		ID:     GenerateID(),
		Sender: sender,
	}
	return newRPC
}

func GenerateResponse(cmd CMD, id [5]uint32, sender contact) RPC {

	newRPC := RPC{
		CMD: cmd,
		ID: id,
		Sender: sender,
	}
	return newRPC
}

func (rpc *RPC) Pong() {
	if rpc.CMD != PONG {
		log.Println("WARNING: setting acknowledgement in non-PONG RPC")
	}
	rpc.Acknowledge = true
}

func (rpc *RPC) Store(wallet wallet) {
	if rpc.CMD != STORE {
		log.Println("WARNING: appending wallet to non-STORE RPC")
	}
	rpc.wallet = wallet
}

func (rpc *RPC) FoundNodes(found *list.List) {
		// insert nodes into the rpc
	if rpc.CMD != FIND_NODE_RESPONSE {
		log.Println("WARNING: appending found nodes to RPC without FOUND_NODE_RESPONSE command")
	}
	for n := found.Front(); n != nil; n = n.Next() {
		rpc.KNodes = append(rpc.KNodes, n.Value.(contact))
	}
}



