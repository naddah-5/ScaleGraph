package scalegraph

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 5   // K, number of contacts per bucket
	REPLICATION   = 10  // alpha
	PORT          = 8080
)

type Node struct {
	Replication int
	BucketSize  int
	KeySpace    int
	ID          [5]uint32
	IP          [4]byte
	network
	routingTable
}

func NewNode(id [5]uint32, ip [4]byte, listener chan RPC, sender chan RPC) Node {
	net := NewNetwork(listener, sender)
	return Node{
		Replication:  REPLICATION,
		BucketSize:   KBUCKETVOLUME,
		KeySpace:     KEYSPACE,
		ID:           id,
		IP:           ip,
		network:      net,
		routingTable: NewRoutingTable(id),
	}
}

func (n *Node) Start(node [4]byte) {}

