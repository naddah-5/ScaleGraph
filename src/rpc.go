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
	FIND_NODE
	FIND_WALLET
	PROPOSE_TRANSACTION
	ACCEPT_TRANSACTION
	SEND
	SUBMIT_WALLET
	SHOW_WALLET
)

func (c CMD) String() string {
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
	case PROPOSE_TRANSACTION:
		return "PROPOSE_TRANSACTION"
	case ACCEPT_TRANSACTION:
		return "ACCEPT_TRANSACTION"
	case SEND:
		return "SEND"
	case SHOW_WALLET:
		return "SHOW_WALLET"
	}
	return "unknown"
}

// Contains fields for all RPC data
// note that fields may be nil.
type RPC struct {
	CMD
	order         CMD
	response      bool
	timeout       bool
	id            [5]uint32
	sender        contact
	receiver      [4]byte
	walletID      [5]uint32
	walletBalance int
	walletKey     []byte
	transaction   transaction
	findTarget    [5]uint32
	kNodes        []contact
	acknowledge   bool
	success       bool
}

// Creates and returns a new base case RPC.
func GenerateRPC(cmd CMD, sender contact, receiver [4]byte) RPC {
	newRPC := RPC{
		CMD:      cmd,
		response: false,
		timeout:  false,
		id:       GenerateID(),
		sender:   sender,
		receiver: receiver,
	}
	return newRPC
}

// Creates and returns a base case response RPC.
func GenerateResponse(cmd CMD, id [5]uint32, ip [4]byte, sender contact) RPC {
	newRPC := RPC{
		CMD:      cmd,
		response: true,
		timeout:  false,
		id:       id,
		sender:   sender,
		receiver: ip,
	}
	return newRPC
}

// Wrapper for redirecting the target IP of a RPC.
func (rpc *RPC) Redirect(ip [4]byte) {
	rpc.receiver = ip
}

func (rpc *RPC) Fail() {
	rpc.success = false
}

// Sets the rpc acknowledgement to true
func (rpc *RPC) Pong() {
	if rpc.CMD != PONG {
		log.Println("WARNING: applying pong to non-PONG RPC")
	}
	rpc.acknowledge = true
}

// Sets the rpc walletID
func (rpc *RPC) Store(walletID [5]uint32, balance int) {
	if rpc.CMD != STORE {
		log.Println("WARNING: applying store to non-STORE RPC")
	}
	rpc.walletID = walletID
	rpc.walletBalance = balance
}

// Sets the rpc acknowledge to true
func (rpc *RPC) StoreResponse() {
	if rpc.CMD != STORE {
		log.Println("WARNING: applying store response to non-STORE_RESPONSE RPC")
	}
	rpc.acknowledge = true
}

// Sets the target id for the rpc.
func (rpc *RPC) FindNode(target [5]uint32) {
	if rpc.CMD != FIND_NODE {
		log.Println("WARNING: applying find node to non-FIND_NODE RPC")
	}
	rpc.findTarget = target
}

// Attatches the given list of contacts to the rpc as a slice.
func (rpc *RPC) FindNodeResponse(found *list.List, target [5]uint32) {
	if rpc.CMD != FIND_NODE {
		log.Println("WARNING: applying find node reponse to non-FIND_NODE_RESPONSE RPC")
	}
	rpc.kNodes = make([]contact, 0)
	for n := found.Front(); n != nil; n = n.Next() {
		rpc.kNodes = append(rpc.kNodes, n.Value.(contact))
	}
	rpc.findTarget = target
}

func (rpc *RPC) FindWallet(success bool) {
	if rpc.CMD != FIND_WALLET {
		log.Println("WARNING: applying find wallet to non-FIND_WALLET RPC")
	}
	rpc.acknowledge = true
	rpc.success = success
}

func (rpc *RPC) ShowWallet(walletID [5]uint32) {
	rpc.walletID = walletID
}

func (rpc *RPC) ShowWalletResponse(walletID [5]uint32, balance int) {
	rpc.walletID = walletID
	rpc.walletBalance = balance
	rpc.success = true
}

//func (rpc *RPC) FindWalletResponse(found *wallet) {
//	if rpc.CMD != FIND_WALLET_RESPONSE {
//		log.Println("WARNING: applying find wallet response to non-FIND_WALLET_RESPONSE")
//	}
//	rpc.wallet = found
//}
