package scalegraph

import (
	"fmt"
	"log"
	"time"
)

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 5   // K, number of contacts per bucket
	REPLICATION   = 10  // alpha
	PORT          = 8080
	DEBUG         = true
	TIMEOUT       = 100 * time.Second
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

func NewNode(id [5]uint32, ip [4]byte, listener chan RPC, sender chan RPC, serverIP [4]byte, master [4]byte) Node {
	net := NewNetwork(listener, sender, serverIP, master)
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

	rpc := GenerateRPC(PING, n.contact, n.serverIP)
	resp, err := n.network.Send(rpc)
	if err != nil {
		log.Printf("\t[node] - %+v", err.Error())
	}
	n.Controller(resp)

	time.Sleep(1 * time.Second)
	rpc = GenerateRPC(FIND_NODE, n.contact, n.master)
	rpc.FindNode(n.ID())
	resp, err = n.network.Send(rpc)
	if err != nil {
		log.Printf("\t[node] - %+v", err.Error())
	}
	n.Controller(resp)

	time.Sleep(60 * time.Second)
	if DEBUG {
		dump := ""
		dump += fmt.Sprintf("[node] - %+v - current routing table:\n", n.ID())
		for i := range n.router {
			for c := n.router[i].content.Front(); c != nil; c = c.Next() {
				dump += fmt.Sprintf("\tcontact: %+v\n", c.Value.(contact).id)
			}
		}
		log.Println(dump)
	}
}
