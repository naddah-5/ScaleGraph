package kademlia

import (
	"fmt"
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
	Contact
	Network
	RoutingTable
	controller chan RPC // the channel for internal network, new rpc's are to be sent here for handling
	debug      bool
}

func NewNode(id [5]uint32, ip [4]byte, listener chan RPC, sender chan RPC, serverIP [4]byte, masterNode Contact, debug bool) *Node {
	controller := make(chan RPC)
	net := NewNetwork(listener, sender, controller, serverIP, masterNode)
	me := NewContact(ip, id)
	router := NewRoutingTable(id, KEYSPACE, KBUCKETVOLUME)
	return &Node{
		Contact:      me,
		Network:      *net,
		RoutingTable: *router,
		debug:        debug,
	}
}

func (node *Node) Start() {
	go node.Network.Listen(node)
	if node.Contact.IP() == node.masterNode.IP() {
	} else {
		node.Ping(node.masterNode.IP())
	}
	time.Sleep(time.Millisecond * 10)
	node.FindNode(node.Contact.ID())
}

func (node *Node) Display() string {
	res := fmt.Sprintf("node ID routing table state: %v\tnode IP: %v", node.IP(), node.ID())
	res += node.RoutingTable.Display()
	return res
}
