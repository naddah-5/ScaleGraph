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
	FIND_WALLET
	FIND_WALLET_RESPONSE
	PROPOSE_REQUEST
	PROPOSE_ACCEPT
	PROPOSE
	PROPOSE_VALIDATE
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
	case FIND_WALLET:
		return "FIND_WALLET"
	case FIND_WALLET_RESPONSE:
		return "FIND_WALLET_RESPONSE"
	}
	return "unknown"
}

// Contains fields for all RPC data
// note that fields may be nil.
type RPC struct {
	CMD
	response bool
	ID       [5]uint32
	Sender   contact
	receiver [4]byte
	wallet
	WalletID    [5]uint32
	WalletKey   []byte
	Transaction []byte
	FindTarget  [5]uint32
	KNodes      []contact
	Acknowledge bool
}

func GenerateRPC(cmd CMD, sender contact, receiver [4]byte) RPC {
	newRPC := RPC{
		CMD:      cmd,
		response: false,
		ID:       GenerateID(),
		Sender:   sender,
		receiver: receiver,
	}
	return newRPC
}

func GenerateResponse(cmd CMD, id [5]uint32, ip [4]byte, sender contact) RPC {

	newRPC := RPC{
		CMD:      cmd,
		response: true,
		ID:       id,
		Sender:   sender,
		receiver: ip,
	}
	return newRPC
}

func (rpc *RPC) Redirect(ip [4]byte) {
	rpc.receiver = ip
}

func (rpc *RPC) Pong() {
	if rpc.CMD != PONG {
		log.Println("WARNING: applying pong to non-PONG RPC")
	}
	rpc.Acknowledge = true
}

func (rpc *RPC) Store(wallet wallet) {
	if rpc.CMD != STORE {
		log.Println("WARNING: applying store to non-STORE RPC")
	}
	rpc.wallet = wallet
}

func (rpc *RPC) StoreResponse(ack bool) {
	if rpc.CMD != STORE_RESPONSE {
		log.Println("WARNING: applying store response to non-STORE_RESPONSE RPC")
	}
	rpc.Acknowledge = true
}

func (rpc *RPC) FindNode(target [5]uint32) {
	if rpc.CMD != FIND_NODE {
		log.Println("WARNING: applying find node to non-FIND_NODE RPC")
	}
	rpc.FindTarget = target
}

func (rpc *RPC) FindNodeResponse(found *list.List) {
	if rpc.CMD != FIND_NODE_RESPONSE {
		log.Println("WARNING: applying find node reponse to non-FIND_NODE_RESPONSE RPC")
	}
	for n := found.Front(); n != nil; n = n.Next() {
		rpc.KNodes = append(rpc.KNodes, n.Value.(contact))
	}
}

func (rpc *RPC) FindWallet(walletID [5]uint32) {
	if rpc.CMD != FIND_WALLET {
		log.Println("WARNING: applying find wallet to non-FIND_WALLET RPC")
	}
	rpc.WalletID = walletID
}

func (rpc *RPC) FindWalletResponse(found wallet) {
	if rpc.CMD != FIND_WALLET_RESPONSE {
		log.Println("WARNING: applying find wallet response to non-FIND_WALLET_RESPONSE")
	}
	rpc.wallet = found
}
