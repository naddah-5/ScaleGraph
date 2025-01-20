package kademlia

import "fmt"

type cmd int

const (
	NO_CMD cmd = iota
	PING
	PONG
	ENTER
	FIND_NODE
	FOUND_NODES
	STORE_ACCOUNT
	STORED_ACCOUNT
	FIND_ACCOUNT
	FOUND_ACCOUNT
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
	case ENTER:
		return "ENTER"
	case FIND_NODE:
		return "FIND_NODE"
	case FOUND_NODES:
		return "FOUND_NODES"
	case STORE_ACCOUNT:
		return "STORE_ACCOUNT"
	case STORED_ACCOUNT:
		return "STORED_ACCOUNT"
	case FIND_ACCOUNT:
		return "FIND_ACCOUNT"
	case FOUND_ACCOUNT:
		return "FOUND_ACCOUNT"
	}
	return "unknown"
}

type RPC struct {
	id              [5]uint32
	cmd             cmd
	response        bool
	sender          Contact
	receiver        [4]byte
	findNodeTarget  [5]uint32
	foundNodes      []Contact
	accountID       [5]uint32
	storeAccSucc    bool
	findAccountSucc bool
}

// Generate a fresh send RPC, for a response RPC use GenerateResponse instead.
func GenerateRPC(receiver [4]byte, sender Contact) RPC {
	rpc := RPC{
		id:       RandomID(),
		receiver: receiver,
		response: false,
		sender:   sender,
	}
	return rpc
}

// Generates a fresh response RPC.
func GenerateResponse(id [5]uint32, receiver [4]byte, sender Contact) RPC {
	rpc := RPC{
		id:       id,
		receiver: receiver,
		response: true,
		sender:   sender,
	}
	return rpc
}

// Set a RPC as a ping.
func (rpc *RPC) Ping() {
	rpc.cmd = PING
}

func (rpc *RPC) Pong() {
	rpc.cmd = PONG
}

// Used to get a random existing node in the network from the simulated network.
func (rpc *RPC) Enter() {
	rpc.cmd = ENTER
}

func (rpc *RPC) FindNode(targetNode [5]uint32) {
	rpc.cmd = FIND_NODE
	rpc.findNodeTarget = targetNode
}

func (rpc *RPC) FoundNodes(target [5]uint32, nodes []Contact) {
	rpc.cmd = FOUND_NODES
	rpc.findNodeTarget = target
	rpc.foundNodes = nodes
}

func (rpc *RPC) StoreAccount(accID [5]uint32) {
	rpc.cmd = STORE_ACCOUNT
	rpc.accountID = accID
}

func (rpc *RPC) StoredAccount(accID [5]uint32, success bool) {
	rpc.cmd = STORED_ACCOUNT
	rpc.accountID = accID
	rpc.storeAccSucc = success
}

func (rpc *RPC) FindAccount(accID [5]uint32) {
	rpc.cmd = FIND_ACCOUNT
	rpc.accountID = accID
}

func (rpc *RPC) FoundAccount(accID [5]uint32, success bool) {
	rpc.cmd = FOUND_ACCOUNT
	rpc.accountID = accID
	rpc.findAccountSucc = success
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
	if rpc.cmd == FOUND_NODES && rpc.response {
		rpcString += "Found Nodes:"
		for _, val := range rpc.foundNodes {
			rpcString += fmt.Sprintf("\n%s", val.Display())
		}
		rpcString += "\n"
	}
	rpcString += "\n\n"

	return rpcString
}
