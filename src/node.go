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
	*network
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

func (node *Node) Start(delay chan struct{}, done chan struct{}, prt chan struct{}) {
	if DEBUG {
		log.Printf("started node: %+v", node.id)
	}
	go node.network.Listen(node)

	_, start := <-delay
	if !start {
		log.Printf("\tstarting node: %+v\n", node.contact.ID())
	}

	go node.Ping(node.serverIP)

	time.Sleep(10 * time.Second)
	go node.FindNode(node.ID())
	time.Sleep(5 * time.Second)
	go node.FindNode(node.ID())
	time.Sleep(5 * time.Second)
	


	done <- struct{}{}
	_, allDone := <-prt
	if !allDone {
		if DEBUG {
			dump := ""
			dump += fmt.Sprintf("[node] - %+v - current routing table:\n", node.ID())
			for i := range node.router {
				for c := node.router[i].content.Front(); c != nil; c = c.Next() {
					dump += fmt.Sprintf("\tcontact: %+v\n", c.Value.(contact).id)
				}
			}
			log.Println(dump)
		}
	}
}
