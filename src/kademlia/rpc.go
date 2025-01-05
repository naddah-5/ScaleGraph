package kademlia

import "fmt"

type cmd int

const (
	NO_CMD cmd = iota
	PING
	PONG
	STORE_WALLET
	FIND_NODE
	FOUND_NODES
	FIND_WALLET
	PROPOSE_TRANSACTION
	ACCEPT_TRANSACTION
	SEND
	SUBMIT_WALLET
	SHOW_WALLET
)

func (c cmd) String() string {
	switch c {
	case NO_CMD:
		return "NO_CMD"
	case PING:
		return "PING"
	case PONG:
		return "PONG"
	case FIND_NODE:
		return "FIND_NODE"
	case FOUND_NODES:
		return "FOUND_NODES"
	case STORE_WALLET:
		return "STORE_WALLET"
	case FIND_WALLET:
		return "FIND_WALLET"
	}
	return "unknown"
}

type RPC struct {
	id             [5]uint32
	cmd            cmd
	response       bool
	sender         Contact
	receiver       [4]byte
	findNodeTarget [5]uint32
	foundNodes     []Contact
}

// Generate a fresh send RPC, for a response RPC use GenerateResponse instead.
func GenerateRPC(sender Contact) RPC {
	rpc := RPC{
		id:       RandomID(),
		response: false,
		sender:   sender,
	}
	return rpc
}

// Generates a fresh response RPC.
func GenerateResponse(id [5]uint32, sender Contact) RPC {
	rpc := RPC{
		id:       id,
		response: true,
		sender:   sender,
	}
	return rpc
}

// Set a RPC as a ping.
func (rpc *RPC) Ping(receiver [4]byte) {
	rpc.cmd = PING
	rpc.receiver = receiver
}

func (rpc *RPC) Pong(receiver [4]byte) {
	rpc.cmd = PONG
	rpc.receiver = receiver
}

func (rpc *RPC) FindNode(receiver [4]byte, targetNode [5]uint32) {
	rpc.cmd = FIND_NODE
	rpc.receiver = receiver
	rpc.findNodeTarget = targetNode
}

func (rpc *RPC) FoundNodes(receiver [4]byte, target [5]uint32, nodes []Contact) {
	rpc.cmd = FOUND_NODES
	rpc.receiver = receiver
	rpc.findNodeTarget = target
	rpc.foundNodes = nodes
}

func (rpc *RPC) Display() string {
	rpcString := fmt.Sprintf("id: %v\n", rpc.id)
	rpcString += fmt.Sprintf("CMD: %s\n", rpc.cmd)
	rpcString += fmt.Sprintf("Response: %t\n", rpc.response)
	rpcString += fmt.Sprintf("Sender: %s\n", rpc.sender.Display())
	rpcString += fmt.Sprintf("Receiver: %v\n", rpc.receiver)

	if rpc.cmd == FIND_NODE {
		rpcString += fmt.Sprintf("Find Node Target: %v", rpc.findNodeTarget)
	}
	if rpc.cmd == FIND_NODE && rpc.response {
		rpcString += "Found Nodes:"
		for _, val := range rpc.foundNodes {
			rpcString += fmt.Sprintf("\n%s", val.Display())
		}
		rpcString += "\n"
	}
	rpcString += "\n"

	return rpcString
}
