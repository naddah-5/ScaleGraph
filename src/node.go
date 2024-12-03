package src

import (
	"main/src/kademlia"
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
	kademlia.Contact
	kademlia.Network
	kademlia.RoutingTable
	controller chan kademlia.RPC // the channel for internal network, new rpc's are to be sent here for handling
}

func NewNode(id [5]uint32, ip [4]byte, listener chan kademlia.RPC, sender chan kademlia.RPC, serverIP [4]byte, masterNode kademlia.Contact) *Node {
	controller := make(chan kademlia.RPC)
	net := kademlia.NewNetwork(listener, sender, controller, serverIP, masterNode)
	me := kademlia.NewContact(ip, id)
	router := kademlia.NewRoutingTable(id, KEYSPACE, KBUCKETVOLUME)
	return &Node{
		Contact: me,
		Network: *net,
		RoutingTable: *router,
	}
}

func (node *Node) Start() {
	go node.Network.Listen()
	time.Sleep(time.Millisecond * 10)
}
