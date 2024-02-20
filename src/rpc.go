package scalegraph

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
	ContentKey  []byte
	Content     []byte
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
