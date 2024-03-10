package scalegraph

import (
	"log"
	"time"
)

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 5   // K, number of contacts per bucket
	REPLICATION   = 10  // alpha
	PORT          = 8080
	DEBUG         = true
	TIMEOUT       = 10000 * time.Second
)

type Node struct {
	Replication int
	BucketSize  int
	KeySpace    int
	contact
	network
	routingTable
	vault
}

func NewNode(id [5]uint32, ip [4]byte, listener chan RPC, sender chan RPC, serverIP [4]byte) Node {
	net := NewNetwork(listener, sender, serverIP)
	me := BuildContact(ip, PORT, id)
	return Node{
		Replication:  REPLICATION,
		BucketSize:   KBUCKETVOLUME,
		KeySpace:     KEYSPACE,
		contact:      me,
		network:      net,
		routingTable: NewRoutingTable(id),
	}
}

func (n *Node) Start() {
	if DEBUG {
		log.Printf("started node: %+v", n.id)
	}
	go n.network.Listen(n)
	time.Sleep(2 * time.Second)
	rpc := GenerateRPC(PING, n.contact, n.serverIP)
	resp, err := n.network.Send(rpc)
	if err != nil {
		log.Println(err.Error())
	}
	n.Handler(resp)
	time.Sleep(10 * time.Second)
	if DEBUG {
		log.Printf("[node] %+v - current routing table:", n.ID()) 
		for i := range n.router {
			for c := n.router[i].content.Front(); c != nil; c = c.Next() {
				log.Printf("\tcontact: %+v", c.Value.(contact).id)
			}
		}
	}
}
