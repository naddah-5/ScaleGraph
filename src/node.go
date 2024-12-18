package scalegraph

import (
	"log"
	"time"
)

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 5   // K, number of contacts per bucket
	REPLICATION   = 3   // alpha
	CONCURRENCY   = 3
	PORT          = 8080
	DEBUG         = true
	POINT_DEBUG   = true
	TIMEOUT       = 10 * time.Second
)

type Node struct {
	Replication int
	BucketSize  int
	KeySpace    int
	contact
	*network
	routingTable
	vault
}

func NewNode(id [5]uint32, ip [4]byte, listener chan RPC, sender chan RPC, serverIP [4]byte, master [4]byte) Node {
	net := NewNetwork(listener, sender, serverIP, master)
	me := BuildContact(ip, id)
	return Node{
		Replication:  REPLICATION,
		BucketSize:   KBUCKETVOLUME,
		KeySpace:     KEYSPACE,
		contact:      me,
		network:      net,
		routingTable: NewRoutingTable(id),
		vault: *NewVault(),
	}
}

func (node *Node) Start() {
	if DEBUG {
		log.Printf("started node: %+v", node.id)
	}
	go node.network.Listen(node)
	time.Sleep(time.Millisecond * 10)
	go node.Ping(node.serverIP)
}
