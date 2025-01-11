package kademlia

import (
	"fmt"
	"time"
)

const (
	KEYSPACE      = 160 // the number of buckets
	KBUCKETVOLUME = 20   // K, number of contacts per bucket
	REPLICATION   = 4   // alpha
	CONCURRENCY   = 10
	PORT          = 8080
	DEBUG         = true
	POINT_DEBUG   = true
	TIMEOUT       = 1 * time.Second
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
	net := NewNetwork(id, listener, sender, controller, serverIP, masterNode, false)
	me := NewContact(ip, id)
	router := NewRoutingTable(id, KEYSPACE, KBUCKETVOLUME)
	return &Node{
		Contact:      me,
		Network:      *net,
		RoutingTable: *router,
		debug:        debug,
	}
}

func (node *Node) Start(done chan [5]uint32) {
	go node.Network.Listen(node)
	if node.Contact.IP() == node.masterNode.IP() {
		return
	} else {
		node.Ping(node.masterNode.IP())
		node.FindNode(node.Contact.ID())
		done <- node.ID()
	}
}

// Wrapper for sending a rpc and also adding the responding contact.
func (node *Node) Send(rpc RPC) (RPC, error) {
	res, err := node.Network.Send(rpc)
	if err != nil {
		return res, err
	} else {
		go node.AddContact(res.sender)
		return res, nil
	}
}

func (node *Node) Debug(mode bool) {
	node.debug = mode
	node.Network.Debug(mode)
}

func (node *Node) Display() string {
	res := fmt.Sprintf("node ID routing table state: node IP %v\tnode ID: %v", node.IP(), node.ID())
	res += node.RoutingTable.Display()
	return res
}
