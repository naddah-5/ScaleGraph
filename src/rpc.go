package scalegraph

type CMD int

const (
	PING CMD = iota
	PONG
	STORE
	FIND_NODE
	FIND_VALUE
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
	case FIND_VALUE:
		return "FIND_VALUE"
	}
	return "unknown"
}

type RPC struct {
	CMD
	Sender  [4]byte
	Content string
}
