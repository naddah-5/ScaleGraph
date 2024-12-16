package kademlia

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
	ID       [5]uint32
	CMD      cmd
	Response bool
	Sender   Contact
	Receiver [4]byte
}

// Generate a fresh send RPC, for a response RPC use GenerateResponse instead.
func GenerateRPC(sender Contact) RPC {
	rpc := RPC{
		ID: RandomID(),
		Response: false,
		Sender: sender,
	}
	return rpc
}

// Generates a fresh response RPC.
func GenerateResponse(id [5]uint32, sender Contact) RPC {
	rpc := RPC{
		ID: id,
		Response: true,
		Sender: sender,
	}
	return rpc
}

// Set a RPC as a ping.
func (rpc *RPC) Ping(receiver [4]byte) {
	rpc.CMD = PING
	rpc.Receiver = receiver
}

// Ping response.
func (rpc *RPC) Pong(receiver [4]byte) {
	rpc.CMD = PONG
	rpc.Receiver = receiver
}
