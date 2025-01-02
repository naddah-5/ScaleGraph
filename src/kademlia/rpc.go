package kademlia

import "fmt"

type cmd int

const (
	PING cmd = iota
	STORE_WALLET
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
	case FIND_NODE:
		return "FIND NODE"
	case STORE_WALLET:
		return "STORE WALLET"
	case FIND_WALLET:
		return "FIND WALLET"
	}
	return "unknown"
}

type RPC struct {
	ID             [5]uint32
	CMD            cmd
	Response       bool
	Sender         Contact
	Receiver       [4]byte
	FindNodeTarget [5]uint32
	FoundNodes     []Contact
}

// Generate a fresh send RPC, for a response RPC use GenerateResponse instead.
func GenerateRPC(sender Contact) RPC {
	rpc := RPC{
		ID:       RandomID(),
		Response: false,
		Sender:   sender,
	}
	return rpc
}

// Generates a fresh response RPC.
func GenerateResponse(id [5]uint32, sender Contact) RPC {
	rpc := RPC{
		ID:       id,
		Response: true,
		Sender:   sender,
	}
	return rpc
}

// Set a RPC as a ping.
func (rpc *RPC) Ping(receiver [4]byte) {
	rpc.CMD = PING
	rpc.Receiver = receiver
}

func (rpc *RPC) FindNode(targetNode [5]uint32) {
	rpc.CMD = FIND_NODE
	rpc.FindNodeTarget = targetNode
}

func (rpc *RPC) Display() string {
	rpcString := fmt.Sprintf("id: %v\n", rpc.ID)
	rpcString += fmt.Sprintf("CMD: %s\n", rpc.CMD)
	rpcString += fmt.Sprintf("Response: %t\n", rpc.Response)
	rpcString += fmt.Sprintf("Sender: %s\n", rpc.Sender.Display())
	rpcString += fmt.Sprintf("Receiver: %v\n", rpc.Receiver)

	if rpc.CMD == FIND_NODE {
		rpcString += fmt.Sprintf("Find Node Target: %v", rpc.FindNodeTarget)
	}
	if rpc.CMD == FIND_NODE && rpc.Response {
		rpcString += "Found Nodes:"
		for _, val := range rpc.FoundNodes {
			rpcString += fmt.Sprintf("\n%s", val.Display())
		}
		rpcString += "\n"
	}

	return rpcString
}
